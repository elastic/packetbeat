//+build ignore

// msg_generate.go is meant to run with go generate. It will use
// go/{importer,types} to track down all the RR struct types. Then for each type
// it will generate pack/unpack methods based on the struct tags. The generated source is
// written to zmsg.go, and is meant to be checked into git.
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/importer"
	"go/types"
	"log"
	"os"
	"strings"
)

var packageHdr = `
// *** DO NOT MODIFY ***
// AUTOGENERATED BY go generate from msg_generate.go

package dns

`

// getTypeStruct will take a type and the package scope, and return the
// (innermost) struct if the type is considered a RR type (currently defined as
// those structs beginning with a RR_Header, could be redefined as implementing
// the RR interface). The bool return value indicates if embedded structs were
// resolved.
func getTypeStruct(t types.Type, scope *types.Scope) (*types.Struct, bool) {
	st, ok := t.Underlying().(*types.Struct)
	if !ok {
		return nil, false
	}
	if st.Field(0).Type() == scope.Lookup("RR_Header").Type() {
		return st, false
	}
	if st.Field(0).Anonymous() {
		st, _ := getTypeStruct(st.Field(0).Type(), scope)
		return st, true
	}
	return nil, false
}

func main() {
	// Import and type-check the package
	pkg, err := importer.Default().Import("github.com/miekg/dns")
	fatalIfErr(err)
	scope := pkg.Scope()

	// Collect actual types (*X)
	var namedTypes []string
	for _, name := range scope.Names() {
		o := scope.Lookup(name)
		if o == nil || !o.Exported() {
			continue
		}
		if st, _ := getTypeStruct(o.Type(), scope); st == nil {
			continue
		}
		if name == "PrivateRR" {
			continue
		}

		// Check if corresponding TypeX exists
		if scope.Lookup("Type"+o.Name()) == nil && o.Name() != "RFC3597" {
			log.Fatalf("Constant Type%s does not exist.", o.Name())
		}

		namedTypes = append(namedTypes, o.Name())
	}

	b := &bytes.Buffer{}
	b.WriteString(packageHdr)

	fmt.Fprint(b, "// pack*() functions\n\n")
	for _, name := range namedTypes {
		o := scope.Lookup(name)
		st, _ := getTypeStruct(o.Type(), scope)

		fmt.Fprintf(b, "func (rr *%s) pack(msg []byte, off int, compression map[string]int, compress bool) (int, error) {\n", name)
		fmt.Fprint(b, `off, err := rr.Hdr.pack(msg, off, compression, compress)
if err != nil {
	return off, err
}
headerEnd := off
`)
		for i := 1; i < st.NumFields(); i++ {
			o := func(s string) {
				fmt.Fprintf(b, s, st.Field(i).Name())
				fmt.Fprint(b, `if err != nil {
return off, err
}
`)
			}

			if _, ok := st.Field(i).Type().(*types.Slice); ok {
				switch st.Tag(i) {
				case `dns:"-"`: // ignored
				case `dns:"txt"`:
					o("off, err = packStringTxt(rr.%s, msg, off)\n")
				case `dns:"opt"`:
					o("off, err = packDataOpt(rr.%s, msg, off)\n")
				case `dns:"nsec"`:
					o("off, err = packDataNsec(rr.%s, msg, off)\n")
				case `dns:"domain-name"`:
					o("off, err = packDataDomainNames(rr.%s, msg, off, compression, compress)\n")
				default:
					log.Fatalln(name, st.Field(i).Name(), st.Tag(i))
				}
				continue
			}

			switch {
			case st.Tag(i) == `dns:"-"`: // ignored
			case st.Tag(i) == `dns:"cdomain-name"`:
				fallthrough
			case st.Tag(i) == `dns:"domain-name"`:
				o("off, err = PackDomainName(rr.%s, msg, off, compression, compress)\n")
			case st.Tag(i) == `dns:"a"`:
				o("off, err = packDataA(rr.%s, msg, off)\n")
			case st.Tag(i) == `dns:"aaaa"`:
				o("off, err = packDataAAAA(rr.%s, msg, off)\n")
			case st.Tag(i) == `dns:"uint48"`:
				o("off, err = packUint48(rr.%s, msg, off)\n")
			case st.Tag(i) == `dns:"txt"`:
				o("off, err = packString(rr.%s, msg, off)\n")

			case strings.HasPrefix(st.Tag(i), `dns:"size-base32`): // size-base32 can be packed just like base32
				fallthrough
			case st.Tag(i) == `dns:"base32"`:
				o("off, err = packStringBase32(rr.%s, msg, off)\n")

			case strings.HasPrefix(st.Tag(i), `dns:"size-base64`): // size-base64 can be packed just like base64
				fallthrough
			case st.Tag(i) == `dns:"base64"`:
				o("off, err = packStringBase64(rr.%s, msg, off)\n")

			case strings.HasPrefix(st.Tag(i), `dns:"size-hex`): // size-hex can be packed just like hex
				fallthrough
			case st.Tag(i) == `dns:"hex"`:
				o("off, err = packStringHex(rr.%s, msg, off)\n")

			case st.Tag(i) == `dns:"octet"`:
				o("off, err = packStringOctet(rr.%s, msg, off)\n")
			case st.Tag(i) == "":
				switch st.Field(i).Type().(*types.Basic).Kind() {
				case types.Uint8:
					o("off, err = packUint8(rr.%s, msg, off)\n")
				case types.Uint16:
					o("off, err = packUint16(rr.%s, msg, off)\n")
				case types.Uint32:
					o("off, err = packUint32(rr.%s, msg, off)\n")
				case types.Uint64:
					o("off, err = packUint64(rr.%s, msg, off)\n")
				case types.String:
					o("off, err = packString(rr.%s, msg, off)\n")
				default:
					log.Fatalln(name, st.Field(i).Name())
				}
			default:
				log.Fatalln(name, st.Field(i).Name(), st.Tag(i))
			}
		}
		// We have packed everything, only now we know the rdlength of this RR
		fmt.Fprintln(b, "rr.Header().Rdlength = uint16(off- headerEnd)")
		fmt.Fprintln(b, "return off, nil }\n")
	}

	fmt.Fprint(b, "// unpack*() functions\n\n")
	for _, name := range namedTypes {
		o := scope.Lookup(name)
		st, _ := getTypeStruct(o.Type(), scope)

		fmt.Fprintf(b, "func unpack%s(h RR_Header, msg []byte, off int) (RR, int, error) {\n", name)
		fmt.Fprintf(b, "rr := new(%s)\n", name)
		fmt.Fprint(b, "rr.Hdr = h\n")
		fmt.Fprint(b, `if noRdata(h) {
return rr, off, nil
	}
var err error
rdStart := off
_ = rdStart

`)
		for i := 1; i < st.NumFields(); i++ {
			o := func(s string) {
				fmt.Fprintf(b, s, st.Field(i).Name())
				fmt.Fprint(b, `if err != nil {
return rr, off, err
}
`)
			}

			// size-* are special, because they reference a struct member we should use for the length.
			if strings.HasPrefix(st.Tag(i), `dns:"size-`) {
				structMember := structMember(st.Tag(i))
				structTag := structTag(st.Tag(i))
				switch structTag {
				case "hex":
					fmt.Fprintf(b, "rr.%s, off, err = unpackStringHex(msg, off, off + int(rr.%s))\n", st.Field(i).Name(), structMember)
				case "base32":
					fmt.Fprintf(b, "rr.%s, off, err = unpackStringBase32(msg, off, off + int(rr.%s))\n", st.Field(i).Name(), structMember)
				case "base64":
					fmt.Fprintf(b, "rr.%s, off, err = unpackStringBase64(msg, off, off + int(rr.%s))\n", st.Field(i).Name(), structMember)
				default:
					log.Fatalln(name, st.Field(i).Name(), st.Tag(i))
				}
				fmt.Fprint(b, `if err != nil {
return rr, off, err
}
`)
				continue
			}

			if _, ok := st.Field(i).Type().(*types.Slice); ok {
				switch st.Tag(i) {
				case `dns:"-"`: // ignored
				case `dns:"txt"`:
					o("rr.%s, off, err = unpackStringTxt(msg, off)\n")
				case `dns:"opt"`:
					o("rr.%s, off, err = unpackDataOpt(msg, off)\n")
				case `dns:"nsec"`:
					o("rr.%s, off, err = unpackDataNsec(msg, off)\n")
				case `dns:"domain-name"`:
					o("rr.%s, off, err = unpackDataDomainNames(msg, off, rdStart + int(rr.Hdr.Rdlength))\n")
				default:
					log.Fatalln(name, st.Field(i).Name(), st.Tag(i))
				}
				continue
			}

			switch st.Tag(i) {
			case `dns:"-"`: // ignored
			case `dns:"cdomain-name"`:
				fallthrough
			case `dns:"domain-name"`:
				o("rr.%s, off, err = UnpackDomainName(msg, off)\n")
			case `dns:"a"`:
				o("rr.%s, off, err = unpackDataA(msg, off)\n")
			case `dns:"aaaa"`:
				o("rr.%s, off, err = unpackDataAAAA(msg, off)\n")
			case `dns:"uint48"`:
				o("rr.%s, off, err = unpackUint48(msg, off)\n")
			case `dns:"txt"`:
				o("rr.%s, off, err = unpackString(msg, off)\n")
			case `dns:"base32"`:
				o("rr.%s, off, err = unpackStringBase32(msg, off, rdStart + int(rr.Hdr.Rdlength))\n")
			case `dns:"base64"`:
				o("rr.%s, off, err = unpackStringBase64(msg, off, rdStart + int(rr.Hdr.Rdlength))\n")
			case `dns:"hex"`:
				o("rr.%s, off, err = unpackStringHex(msg, off, rdStart + int(rr.Hdr.Rdlength))\n")
			case `dns:"octet"`:
				o("rr.%s, off, err = unpackStringOctet(msg, off)\n")
			case "":
				switch st.Field(i).Type().(*types.Basic).Kind() {
				case types.Uint8:
					o("rr.%s, off, err = unpackUint8(msg, off)\n")
				case types.Uint16:
					o("rr.%s, off, err = unpackUint16(msg, off)\n")
				case types.Uint32:
					o("rr.%s, off, err = unpackUint32(msg, off)\n")
				case types.Uint64:
					o("rr.%s, off, err = unpackUint64(msg, off)\n")
				case types.String:
					o("rr.%s, off, err = unpackString(msg, off)\n")
				default:
					log.Fatalln(name, st.Field(i).Name())
				}
			default:
				log.Fatalln(name, st.Field(i).Name(), st.Tag(i))
			}
			// If we've hit len(msg) we return without error.
			if i < st.NumFields()-1 {
				fmt.Fprintf(b, `if off == len(msg) {
return rr, off, nil
	}
`)
			}
		}
		fmt.Fprintf(b, "return rr, off, err }\n\n")
	}
	// Generate typeToUnpack map
	fmt.Fprintln(b, "var typeToUnpack = map[uint16]func(RR_Header, []byte, int) (RR, int, error){")
	for _, name := range namedTypes {
		if name == "RFC3597" {
			continue
		}
		fmt.Fprintf(b, "Type%s: unpack%s,\n", name, name)
	}
	fmt.Fprintln(b, "}\n")

	// gofmt
	res, err := format.Source(b.Bytes())
	if err != nil {
		b.WriteTo(os.Stderr)
		log.Fatal(err)
	}

	// write result
	f, err := os.Create("zmsg.go")
	fatalIfErr(err)
	defer f.Close()
	f.Write(res)
}

// structMember will take a tag like dns:"size-base32:SaltLength" and return the last part of this string.
func structMember(s string) string {
	fields := strings.Split(s, ":")
	if len(fields) == 0 {
		return ""
	}
	f := fields[len(fields)-1]
	// f should have a closing "
	if len(f) > 1 {
		return f[:len(f)-1]
	}
	return f
}

// structTag will take a tag like dns:"size-base32:SaltLength" and return base32.
func structTag(s string) string {
	fields := strings.Split(s, ":")
	if len(fields) < 2 {
		return ""
	}
	return fields[1][len("\"size-"):]
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
