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
BRs: 7.1.4.2.1
Subject Alternative Name Extension
Certificate Field: extensions:subjectAltName
Required/Optional: Required
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANMissing struct{}

func (l *SANMissing) Initialize() error {
	return nil
}

func (l *SANMissing) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c)
}

func (l *SANMissing) Execute(c *x509.Certificate) *LintResult {
	if util.IsExtInCert(c, util.SubjectAlternateNameOID) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_missing",
		Description:   "Subscriber certificates MUST contain the Subject Alternate Name extension",
		Citation:      "BRs: 7.1.4.2.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &SANMissing{},
	})
}
