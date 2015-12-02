// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package language

import (
	"fmt"
	"sort"
)

// The Coverage interface is used to define the level of coverage of an
// internationalization service. Note that not all types are supported by all
// services. As lists may be generated on the fly, it is recommended that users
// of a Coverage cache the results.
type Coverage interface {
	// Tags returns the list of supported tags.
	Tags() []Tag

	// BaseLanguages returns the list of supported base languages.
	BaseLanguages() []Base

	// Scripts returns the list of supported scripts.
	Scripts() []Script

	// Regions returns the list of supported regions.
	Regions() []Region

	// Currencies returns the list of supported currencies.
	Currencies() []Currency
}

var (
	// Supported defines a Coverage that lists all supported subtags. Tags
	// always returns nil.
	Supported Coverage = allSubtags{}
)

// TODO:
// - Support Variants, numbering systems.
// - CLDR coverage levels.
// - Set of common tags defined in this package.

type allSubtags struct{}

// Regions returns the list of supported regions. As all regions are in a
// consecutive range, it simply returns a slice of numbers in increasing order.
// The "undefined" region is not returned.
func (s allSubtags) Regions() []Region {
	reg := make([]Region, numRegions)
	for i := range reg {
		reg[i] = Region{regionID(i + 1)}
	}
	return reg
}

// Scripts returns the list of supported scripts. As all scripts are in a
// consecutive range, it simply returns a slice of numbers in increasing order.
// The "undefined" script is not returned.
func (s allSubtags) Scripts() []Script {
	scr := make([]Script, numScripts)
	for i := range scr {
		scr[i] = Script{scriptID(i + 1)}
	}
	return scr
}

// BaseLanguages returns the list of all supported base languages. It generates
// the list by traversing the internal structures.
func (s allSubtags) BaseLanguages() []Base {
	base := make([]Base, 0, numLanguages)
	for i := 0; i < langNoIndexOffset; i++ {
		// We included "und" already for the value 0.
		if i != nonCanonicalUnd {
			base = append(base, Base{langID(i)})
		}
	}
	i := langNoIndexOffset
	for _, v := range langNoIndex {
		for k := 0; k < 8; k++ {
			if v&1 == 1 {
				base = append(base, Base{langID(i)})
			}
			v >>= 1
			i++
		}
	}
	return base
}

// Currencies returns the list of supported currencies. As all currencies are in
// a consecutive range, it simply returns a slice of numbers in increasing
// order. The "undefined" currency is not returned.
func (s allSubtags) Currencies() []Currency {
	cur := make([]Currency, numCurrencies)
	for i := range cur {
		cur[i] = Currency{currencyID(i + 1)}
	}
	return cur
}

// Tags always returns nil.
func (s allSubtags) Tags() []Tag {
	return nil
}

// coverage is used used by NewCoverage which is used as a convenient way for
// creating Coverage implementations for partially defined data. Very often a
// package will only need to define a subset of slices. coverage provides a
// convenient way to do this. Moreover, packages using NewCoverage, instead of
// their own implementation, will not break if later new slice types are added.
type coverage struct {
	tags       func() []Tag
	bases      func() []Base
	scripts    func() []Script
	regions    func() []Region
	currencies func() []Currency
}

func (s *coverage) Tags() []Tag {
	if s.tags == nil {
		return nil
	}
	return s.tags()
}

// bases implements sort.Interface and is used to sort base languages.
type bases []Base

func (b bases) Len() int {
	return len(b)
}

func (b bases) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bases) Less(i, j int) bool {
	return b[i].langID < b[j].langID
}

// BaseLanguages returns the result from calling s.bases if it is specified or
// otherwise derives the set of supported base languages from tags.
func (s *coverage) BaseLanguages() []Base {
	if s.bases == nil {
		tags := s.Tags()
		if len(tags) == 0 {
			return nil
		}
		a := make([]Base, len(tags))
		for i, t := range tags {
			a[i] = Base{langID(t.lang)}
		}
		sort.Sort(bases(a))
		k := 0
		for i := 1; i < len(a); i++ {
			if a[k] != a[i] {
				k++
				a[k] = a[i]
			}
		}
		return a[:k+1]
	}
	return s.bases()
}

func (s *coverage) Scripts() []Script {
	if s.scripts == nil {
		return nil
	}
	return s.scripts()
}

func (s *coverage) Regions() []Region {
	if s.regions == nil {
		return nil
	}
	return s.regions()
}

func (s *coverage) Currencies() []Currency {
	if s.currencies == nil {
		return nil
	}
	return s.currencies()
}

// NewCoverage returns a Coverage for the given lists. It is typically used by
// packages providing internationalization services to define their level of
// coverage. A list may be of type []T or func() []T, where T is either Tag,
// Base, Script or Region. The returned Coverage derives the value for Bases
// from Tags if no func or slice for []Base is specified. For other unspecified
// types the returned Coverage will return nil for the respective methods.
func NewCoverage(list ...interface{}) Coverage {
	s := &coverage{}
	for _, x := range list {
		switch v := x.(type) {
		case func() []Base:
			s.bases = v
		case func() []Script:
			s.scripts = v
		case func() []Region:
			s.regions = v
		case func() []Tag:
			s.tags = v
		case func() []Currency:
			s.currencies = v
		case []Base:
			s.bases = func() []Base { return v }
		case []Script:
			s.scripts = func() []Script { return v }
		case []Region:
			s.regions = func() []Region { return v }
		case []Tag:
			s.tags = func() []Tag { return v }
		case []Currency:
			s.currencies = func() []Currency { return v }
		default:
			panic(fmt.Sprintf("language: unsupported set type %T", v))
		}
	}
	return s
}
