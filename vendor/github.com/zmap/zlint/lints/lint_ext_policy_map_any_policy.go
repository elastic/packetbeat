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

/********************************************************************
RFC 5280: 4.2.1.5
Each issuerDomainPolicy named in the policy mappings extension SHOULD
   also be asserted in a certificate policies extension in the same
   certificate.  Policies MUST NOT be mapped either to or from the
   special value anyPolicy (Section 4.2.1.4).
********************************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type policyMapAnyPolicy struct{}

func (l *policyMapAnyPolicy) Initialize() error {
	return nil
}

func (l *policyMapAnyPolicy) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.PolicyMapOID)
}

func (l *policyMapAnyPolicy) Execute(c *x509.Certificate) *LintResult {
	extPolMap := util.GetExtFromCert(c, util.PolicyMapOID)
	polMap, err := util.GetMappedPolicies(extPolMap)
	if err != nil {
		return &LintResult{Status: Fatal}
	}

	for _, pair := range polMap {
		if util.AnyPolicyOID.Equal(pair[0]) || util.AnyPolicyOID.Equal(pair[1]) {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_policy_map_any_policy",
		Description:   "Policies must not be mapped to or from the anyPolicy value",
		Citation:      "RFC 5280: 4.2.1.5",
		Source:        RFC5280,
		EffectiveDate: util.RFC3280Date,
		Lint:          &policyMapAnyPolicy{},
	})
}
