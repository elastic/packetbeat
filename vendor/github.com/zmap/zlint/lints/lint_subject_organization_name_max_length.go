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

/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-organization-name INTEGER ::= 64
************************************************/

import (
	"unicode/utf8"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectOrganizationNameMaxLength struct{}

func (l *subjectOrganizationNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectOrganizationNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectOrganizationNameMaxLength) Execute(c *x509.Certificate) *LintResult {
	for _, j := range c.Subject.Organization {
		if utf8.RuneCountInString(j) > 64 {
			return &LintResult{Status: Error}
		}
	}

	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_organization_name_max_length",
		Description:   "The 'Organization Name' field of the subject MUST be less than 64 characters",
		Citation:      "RFC 5280: A.1",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &subjectOrganizationNameMaxLength{},
	})
}
