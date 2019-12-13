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

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertLocalityNameMustAppear struct{}

func (l *subCertLocalityNameMustAppear) Initialize() error {
	return nil
}

func (l *subCertLocalityNameMustAppear) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertLocalityNameMustAppear) Execute(c *x509.Certificate) *LintResult {
	if len(c.Subject.Organization) > 0 || len(c.Subject.GivenName) > 0 || len(c.Subject.Surname) > 0 {
		if len(c.Subject.Province) == 0 {
			if len(c.Subject.Locality) == 0 {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_locality_name_must_appear",
		Description:   "Subscriber Certificate: subject:localityName MUST appear if subject:organizationName, subject:givenName, or subject:surname fields are present but the subject:stateOrProvinceName field is absent.",
		Citation:      "BRs: 7.1.4.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABGivenNameDate,
		Lint:          &subCertLocalityNameMustAppear{},
	})
}
