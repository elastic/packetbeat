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

/*************************************************************************
RFC 5280: 4.2 & 4.2.1.6
Further, if the only subject identity included in the certificate is
an alternative name form (e.g., an electronic mail address), then the
subject distinguished name MUST be empty (an empty sequence), and the
subjectAltName extension MUST be present.  If the subject field
contains an empty sequence, then the issuing CA MUST include a
subjectAltName extension that is marked as critical.  When including
the subjectAltName extension in a certificate that has a non-empty
subject distinguished name, conforming CAs SHOULD mark the
subjectAltName extension as non-critical.
*************************************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type emptyWithoutSAN struct{}

func (l *emptyWithoutSAN) Initialize() error {
	return nil
}

func (l *emptyWithoutSAN) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *emptyWithoutSAN) Execute(cert *x509.Certificate) *LintResult {
	if subjectIsEmpty(cert) && !util.IsExtInCert(cert, util.SubjectAlternateNameOID) {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func subjectIsEmpty(cert *x509.Certificate) bool {
	if cert.Subject.Names == nil {
		return true
	}
	return false
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_empty_without_san",
		Description:   "CAs MUST support subject alternative name if the subject field is an empty sequence",
		Citation:      "RFC 5280: 4.2 & 4.2.1.6",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &emptyWithoutSAN{},
	})
}
