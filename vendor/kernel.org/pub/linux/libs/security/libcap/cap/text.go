package cap

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

// String converts a capability Value into its canonical text
// representation.
func (v Value) String() string {
	name, ok := names[v]
	if ok {
		return name
	}
	// Un-named capabilities are referred to numerically (in decimal).
	return strconv.Itoa(int(v))
}

// FromName converts a named capability Value to its binary
// representation.
func FromName(name string) (Value, error) {
	startUp.Do(multisc.cInit)
	v, ok := bits[name]
	if ok {
		if v >= Value(words*32) {
			return 0, ErrBadValue
		}
		return v, nil
	}
	i, err := strconv.Atoi(name)
	if err != nil {
		return 0, err
	}
	if i >= 0 && i < int(words*32) {
		return Value(i), nil
	}
	return 0, ErrBadValue
}

const (
	eBin uint = (1 << Effective)
	pBin      = (1 << Permitted)
	iBin      = (1 << Inheritable)
)

var combos = []string{"", "e", "p", "ep", "i", "ei", "ip", "eip"}

// histo generates a histogram of flag state combinations.
func (c *Set) histo(bins []int, patterns []uint, from, limit Value) uint {
	for v := from; v < limit; v++ {
		b := uint(v & 31)
		u, bit, err := bitOf(0, v)
		if err != nil {
			break
		}
		x := uint((c.flat[u][Effective]&bit)>>b) * eBin
		x |= uint((c.flat[u][Permitted]&bit)>>b) * pBin
		x |= uint((c.flat[u][Inheritable]&bit)>>b) * iBin
		bins[x]++
		patterns[uint(v)] = x
	}
	// Note, in the loop, we use >= to pick the smallest value for
	// m with the highest bin value. That is ties break towards
	// m=0.
	m := uint(7)
	for t := m; t > 0; {
		t--
		if bins[t] >= bins[m] {
			m = t
		}
	}
	return m
}

// String converts a full capability Set into a single short readable
// string representation (which may contain spaces). See the
// cap.FromText() function for an explanation of its return values.
//
// Note (*cap.Set).String() may evolve to generate more compact
// strings representing the a given Set over time, but it should
// maintain compatibility with the libcap:cap_to_text() function for
// any given release. Further, it will always be an inverse of
// cap.FromText().
func (c *Set) String() string {
	if c == nil || len(c.flat) == 0 {
		return "<invalid>"
	}
	bins := make([]int, 8)
	patterns := make([]uint, maxValues)

	c.mu.RLock()
	defer c.mu.RUnlock()

	// Note, in order to have a *Set pointer, startUp.Do(cInit)
	// must have been called which sets maxValues.
	m := c.histo(bins, patterns, 0, Value(maxValues))

	// Background state is the most popular of the named bits.
	vs := []string{"=" + combos[m]}
	for i := uint(8); i > 0; {
		i--
		if i == m || bins[i] == 0 {
			continue
		}
		var list []string
		for j, p := range patterns {
			if p != i {
				continue
			}
			list = append(list, Value(j).String())
		}
		x := strings.Join(list, ",")
		var y, z string
		if cf := i & ^m; cf != 0 {
			op := "+"
			if len(vs) == 1 && vs[0] == "=" {
				// Special case "= foo+..." == "foo=...".
				// Prefer because it
				vs = nil
				op = "="
			}
			y = op + combos[cf]
		}
		if cf := m & ^i; cf != 0 {
			z = "-" + combos[cf]
		}
		vs = append(vs, x+y+z)
	}

	// The unnamed bits can only add to the above named ones since
	// unnamed ones are always defaulted to lowered.
	uBins := make([]int, 8)
	uPatterns := make([]uint, 32*words)
	c.histo(uBins, uPatterns, Value(maxValues), 32*Value(words))
	for i := uint(7); i > 0; i-- {
		if uBins[i] == 0 {
			continue
		}
		var list []string
		for j, p := range uPatterns {
			if p != i {
				continue
			}
			list = append(list, Value(j).String())
		}
		vs = append(vs, strings.Join(list, ",")+"+"+combos[i])
	}

	return strings.Join(vs, " ")
}

// ErrBadText is returned if the text for a capability set cannot be parsed.
var ErrBadText = errors.New("bad text")

// FromText converts the canonical text representation for a Set into
// a freshly allocated Set.
//
// The format follows the following pattern: a set of space separated
// sequences. Each sequence applies over the previous sequence to
// build up a Set. The format of a sequence is:
//
//   [comma list of cap_values][[ops][flags]]*
//
// Examples:
//
//   "all=ep"
//   "cap_chown,cap_setuid=ip cap_setuid+e"
//   "=p cap_setpcap-p+i"
//
// Here "all" refers to all named capabilities known to the hosting
// kernel, and "all" is assumed if no capabilities are listed before
// an "=".
//
// The ops values, "=", "+" and "-" imply "reset and raise", "raise"
// and "lower" respectively. The "e", "i" and "p" characters
// correspond to the capabilities of the corresponding Flag: "e"
// (Effective); "i" (Inheritable); "p" (Permitted).
//
// This syntax is overspecified and there are many ways of building
// the same final Set state. Any sequence that includes a '=' resets
// the accumulated state of all Flags ignoring earlier sequences. On
// each of the following lines we give three or more examples of ways
// to specify a common Set. The last entry on each line is the one
// generated by (*cap.Set).String() from that Set.
//
//    "=p all+ei"  "all=pie"   "=pi all+e"   "=eip"
//
//    "cap_setuid=p cap_chown=i"  "cap_chown=ip-p"   "cap_chown=i"
//
//    "cap_chown=-p"   "all="   "cap_setuid=pie-pie"   "="
//
// Note: FromText() is tested at release time to completely match the
// import ability of the libcap:cap_from_text() function.
func FromText(text string) (*Set, error) {
	c := NewSet()
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	chunks := 0
	for scanner.Scan() {
		chunks++

		// Parsing for xxx([-+=][eip]+)+
		t := scanner.Text()
		i := strings.IndexAny(t, "=+-")
		if i < 0 {
			return nil, ErrBadText
		}
		var vs []Value
		sep := t[i]
		if vals := t[:i]; vals == "all" {
			for v := Value(0); v < Value(maxValues); v++ {
				vs = append(vs, v)
			}
		} else if vals != "" {
			for _, name := range strings.Split(vals, ",") {
				v, err := FromName(name)
				if err != nil {
					return nil, ErrBadText
				}
				vs = append(vs, v)
			}
		} else if sep != '=' {
			if vals == "" {
				// Only "=" supports ""=="all".
				return nil, ErrBadText
			}
		} else if j := i + 1; j+1 < len(t) {
			switch t[j] {
			case '+':
				sep = 'P'
				i++
			case '-':
				sep = 'M'
				i++
			}
		}
		i++

		// There are 5 ways to set: =, =+, =-, +, -. We call
		// the 2nd and 3rd of these 'P' and 'M'.

		for {
			// read [eip]+ setting flags.
			var fE, fP, fI bool
			for ok := true; ok && i < len(t); i++ {
				switch t[i] {
				case 'e':
					fE = true
				case 'i':
					fI = true
				case 'p':
					fP = true
				default:
					ok = false
				}
				if !ok {
					break
				}
			}

			if !(fE || fI || fP) {
				if sep != '=' {
					return nil, ErrBadText
				}
			}

			switch sep {
			case '=', 'P', 'M', '+':
				if sep != '+' {
					c.Clear()
					if sep == 'M' {
						break
					}
				}
				if keep := len(vs) == 0; keep {
					if sep != '=' {
						return nil, ErrBadText
					}
					c.forceFlag(Effective, fE)
					c.forceFlag(Permitted, fP)
					c.forceFlag(Inheritable, fI)
					break
				}
				// =, + and P for specific values are left.
				if fE {
					c.SetFlag(Effective, true, vs...)
				}
				if fP {
					c.SetFlag(Permitted, true, vs...)
				}
				if fI {
					c.SetFlag(Inheritable, true, vs...)
				}
			case '-':
				if fE {
					c.SetFlag(Effective, false, vs...)
				}
				if fP {
					c.SetFlag(Permitted, false, vs...)
				}
				if fI {
					c.SetFlag(Inheritable, false, vs...)
				}
			}

			if i == len(t) {
				break
			}

			switch t[i] {
			case '+', '-':
				sep = t[i]
				i++
			default:
				return nil, ErrBadText
			}
		}
	}
	if chunks == 0 {
		return nil, ErrBadText
	}
	return c, nil
}
