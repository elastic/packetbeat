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

type subCAEKUMissing struct{}

func (l *subCAEKUMissing) Initialize() error {
	return nil
}

func (l *subCAEKUMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c)
}

func (l *subCAEKUMissing) Execute(c *x509.Certificate) *LintResult {
	if util.IsExtInCert(c, util.EkuSynOid) {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Notice}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "n_sub_ca_eku_missing",
		Description:   "To be considered Technically Constrained, the Subordinate CA certificate MUST have extkeyUsage extension",
		Citation:      "BRs: 7.1.5",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCAEKUMissing{},
	})
}
