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
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type qcStatemQctypeWeb struct{}

func (this *qcStatemQctypeWeb) getStatementOid() *asn1.ObjectIdentifier {
	return &util.IdEtsiQcsQcType
}

func (l *qcStatemQctypeWeb) Initialize() error {
	return nil
}

func (l *qcStatemQctypeWeb) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	if util.ParseQcStatem(util.GetExtFromCert(c, util.QcStateOid).Value, *l.getStatementOid()).IsPresent() {
		return true
	}
	return false
}

func (l *qcStatemQctypeWeb) Execute(c *x509.Certificate) *LintResult {

	errString := ""
	wrnString := ""
	ext := util.GetExtFromCert(c, util.QcStateOid)
	s := util.ParseQcStatem(ext.Value, *l.getStatementOid())
	errString += s.GetErrorInfo()
	if len(errString) == 0 {
		qcType := s.(util.Etsi423QcType)
		if len(qcType.TypeOids) == 0 {
			errString += "no QcType present, sequence of OIDs is empty"
		}
		found := false
		for _, t := range qcType.TypeOids {

			if t.Equal(util.IdEtsiQcsQctWeb) {
				found = true
			}
		}
		if found != true {
			wrnString += fmt.Sprintf("etsi Type does not indicate certificate as a 'web' certificate")

		}
	}

	if len(errString) == 0 {
		if len(wrnString) == 0 {
			return &LintResult{Status: Pass}
		} else {
			return &LintResult{Status: Warn, Details: wrnString}
		}
	} else {
		return &LintResult{Status: Error, Details: errString}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_qcstatem_qctype_web",
		Description:   "Checks that a QC Statement of the type Id-etsi-qcs-QcType features features at least the type IdEtsiQcsQctWeb",
		Citation:      "ETSI EN 319 412 - 5 V2.2.1 (2017 - 11) / Section 4.2.3",
		Source:        EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemQctypeWeb{},
	})
}
