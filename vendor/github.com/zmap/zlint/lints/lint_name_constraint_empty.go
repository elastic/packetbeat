package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/***********************************************************************
 Restrictions are defined in terms of permitted or excluded name
   subtrees.  Any name matching a restriction in the excludedSubtrees
   field is invalid regardless of information appearing in the
   permittedSubtrees.  Conforming CAs MUST mark this extension as
   critical and SHOULD NOT impose name constraints on the x400Address,
   ediPartyName, or registeredID name forms.  Conforming CAs MUST NOT
   issue certificates where name constraints is an empty sequence.  That
   is, either the permittedSubtrees field or the excludedSubtrees MUST
   be present.
************************************************************************/

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type nameConstraintEmpty struct{}

func (l *nameConstraintEmpty) Initialize() error {
	return nil
}

func (l *nameConstraintEmpty) CheckApplies(c *x509.Certificate) bool {
	if !(util.IsExtInCert(c, util.NameConstOID)) {
		return false
	}
	nc := util.GetExtFromCert(c, util.NameConstOID)
	var seq asn1.RawValue
	rest, err := asn1.Unmarshal(nc.Value, &seq) //only one sequence, so rest should be empty
	if err != nil || len(rest) != 0 || seq.Tag != 16 || seq.Class != 0 || !seq.IsCompound {
		return false
	}
	return true
}

func (l *nameConstraintEmpty) Execute(c *x509.Certificate) *LintResult {
	nc := util.GetExtFromCert(c, util.NameConstOID)
	var seq asn1.RawValue
	_, err := asn1.Unmarshal(nc.Value, &seq) //only one sequence, so rest should be empty
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	if len(seq.Bytes) == 0 {
		return &LintResult{Status: Error}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_name_constraint_empty",
		Description:   "Conforming CAs MUST NOT issue certificates where name constraints is an empty sequence. That is, either the permittedSubtree or excludedSubtree fields must be present",
		Citation:      "RFC 5280: 4.2.1.10",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &nameConstraintEmpty{},
	})
}
