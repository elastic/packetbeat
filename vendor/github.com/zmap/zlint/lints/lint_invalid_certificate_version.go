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
Certificates MUST be of type X.509 v3.
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type InvalidCertificateVersion struct{}

func (l *InvalidCertificateVersion) Initialize() error {
	return nil
}

func (l *InvalidCertificateVersion) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *InvalidCertificateVersion) Execute(cert *x509.Certificate) *LintResult {
	if cert.Version != 3 {
		return &LintResult{Status: Error}
	}
	//else
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_invalid_certificate_version",
		Description:   "Certificates MUST be of type X.590 v3",
		Citation:      "BRs: 7.1.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABV130Date,
		Lint:          &InvalidCertificateVersion{},
	})
}
