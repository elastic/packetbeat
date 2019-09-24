/*
 * ZLint Copyright 2017 Regents of the University of Michigan
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

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type qcStatemQcmandatoryEtsiStatems struct{}

func (l *qcStatemQcmandatoryEtsiStatems) Initialize() error {
	return nil
}

func (l *qcStatemQcmandatoryEtsiStatems) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	if util.IsAnyEtsiQcStatementPresent(util.GetExtFromCert(c, util.QcStateOid).Value) {
		return true
	}
	return false
}

func (l *qcStatemQcmandatoryEtsiStatems) Execute(c *x509.Certificate) *LintResult {
	errString := ""
	ext := util.GetExtFromCert(c, util.QcStateOid)

	oidList := make([]*asn1.ObjectIdentifier, 1)
	oidList[0] = &util.IdEtsiQcsQcCompliance

	for _, oid := range oidList {
		r := util.ParseQcStatem(ext.Value, *oid)
		util.AppendToStringSemicolonDelim(&errString, r.GetErrorInfo())
		if !r.IsPresent() {
			util.AppendToStringSemicolonDelim(&errString, "missing mandatory ETSI QC statement")
		}
	}

	if len(errString) == 0 {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error, Details: errString}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_qcstatem_mandatory_etsi_statems",
		Description:   "Checks that a QC Statement that contains at least one of the ETSI ESI statements, also features the set of mandatory ETSI ESI QC statements.",
		Citation:      "ETSI EN 319 412 - 5 V2.2.1 (2017 - 11) / Section 5",
		Source:        EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemQcmandatoryEtsiStatems{},
	})
}
