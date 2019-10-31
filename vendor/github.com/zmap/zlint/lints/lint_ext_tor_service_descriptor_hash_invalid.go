/*
 * ZLint Copyright 2019 Regents of the University of Michigan
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
	"fmt"
	"net/url"
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type torServiceDescHashInvalid struct{}

func (l *torServiceDescHashInvalid) Initialize() error {
	// There is nothing to initialize for a torServiceDescHashInvalid linter.
	return nil
}

// CheckApplies returns true if the certificate is a subscriber certificate that
// contains a subject name ending in `.onion`.
func (l *torServiceDescHashInvalid) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.CertificateSubjInTLD(c, onionTLD)
}

// failResult is a small utility function for creating a failed lint result.
func failResult(format string, args ...interface{}) *LintResult {
	return &LintResult{
		Status:  Error,
		Details: fmt.Sprintf(format, args...),
	}
}

// torServiceDescExtName is a common string prefix used in many lint result
// detail messages to identify the extension at fault.
var torServiceDescExtName = fmt.Sprintf(
	"TorServiceDescriptor extension (oid %s)",
	util.BRTorServiceDescriptor.String())

// lintOnionURL verifies that an Onion URI value from a TorServiceDescriptorHash
// is:
//
// 1) a valid parseable url.
// 2) a URL with a non-empty hostname
// 3) a URL with an https:// protocol scheme
//
// If all of the above hold then nil is returned. If any of the above conditions
// are not met an error lint result pointer is returned.
func lintOnionURL(onion string) *LintResult {
	if onionURL, err := url.Parse(onion); err != nil {
		return failResult(
			"%s contained "+
				"TorServiceDescriptorHash object with invalid Onion URI",
			torServiceDescExtName)
	} else if onionURL.Host == "" {
		return failResult(
			"%s contained "+
				"TorServiceDescriptorHash object with Onion URI missing a hostname",
			torServiceDescExtName)
	} else if onionURL.Scheme != "https" {
		return failResult(
			"%s contained "+
				"TorServiceDescriptorHash object with Onion URI using a non-HTTPS "+
				"protocol scheme",
			torServiceDescExtName)
	}
	return nil
}

// Execute will lint the provided certificate. An Error LintResult will be
// returned if:
//
//   1) There is no TorServiceDescriptor extension present.
//   2) There were no TorServiceDescriptors parsed by zcrypto
//   3) There are TorServiceDescriptorHash entries with an invalid Onion URL.
//   4) There are TorServiceDescriptorHash entries with an unknown hash
//      algorithm or incorrect hash bit length.
//   5) There is a TorServiceDescriptorHash entry that doesn't correspond to
//      an onion subject in the cert.
//   6) There is an onion subject in the cert that doesn't correspond to
//      a TorServiceDescriptorHash.
func (l *torServiceDescHashInvalid) Execute(c *x509.Certificate) *LintResult {
	// If the BRTorServiceDescriptor extension is missing return a lint error. We
	// know the cert contains one or more `.onion` subjects because of
	// `CheckApplies` and all such certs are expected to have this extension after
	// util.CABV201Date.
	if ext := util.GetExtFromCert(c, util.BRTorServiceDescriptor); ext == nil {
		return failResult(
			"certificate contained a %s domain but is missing a TorServiceDescriptor "+
				"extension (oid %s)",
			onionTLD, util.BRTorServiceDescriptor.String())
	}

	// The certificate should have at least one TorServiceDescriptorHash in the
	// TorServiceDescriptor extension.
	descriptors := c.TorServiceDescriptors
	if len(descriptors) == 0 {
		return failResult(
			"certificate contained a %s domain but TorServiceDescriptor "+
				"extension (oid %s) had no TorServiceDescriptorHash objects",
			onionTLD, util.BRTorServiceDescriptor.String())
	}

	// Build a map of all the eTLD+1 onion subjects in the cert to compare against
	// the service descriptors.
	onionETLDPlusOneMap := make(map[string]string)
	for _, subj := range append(c.DNSNames, c.Subject.CommonName) {
		if !strings.HasSuffix(subj, onionTLD) {
			continue
		}
		labels := strings.Split(subj, ".")
		if len(labels) < 2 {
			return failResult("certificate contained a %s domain with too few "+
				"labels: %q",
				onionTLD, subj)
		}
		eTLDPlusOne := strings.Join(labels[len(labels)-2:], ".")
		onionETLDPlusOneMap[eTLDPlusOne] = subj
	}

	expectedHashBits := map[string]int{
		"SHA256": 256,
		"SHA384": 384,
		"SHA512": 512,
	}

	// Build a map of onion hostname -> TorServiceDescriptorHash using the parsed
	// TorServiceDescriptors from zcrypto.
	descriptorMap := make(map[string]*x509.TorServiceDescriptorHash)
	for _, descriptor := range descriptors {
		// each descriptor's Onion URL must be valid
		if errResult := lintOnionURL(descriptor.Onion); errResult != nil {
			return errResult
		}
		// each descriptor should have a known hash algorithm and the correct
		// corresponding size of hash.
		if expectedBits, found := expectedHashBits[descriptor.AlgorithmName]; !found {
			return failResult(
				"%s contained a TorServiceDescriptorHash for Onion URI %q with an "+
					"unknown hash algorithm",
				torServiceDescExtName, descriptor.Onion)
		} else if expectedBits != descriptor.HashBits {
			return failResult(
				"%s contained a TorServiceDescriptorHash with hash algorithm %q but "+
					"only %d bits of hash not %d",
				torServiceDescExtName, descriptor.AlgorithmName,
				descriptor.HashBits, expectedBits)
		}
		// NOTE(@cpu): Throwing out the err result here because lintOnionURL already
		//             ensured the URL is valid.
		url, _ := url.Parse(descriptor.Onion)
		hostname := url.Hostname()
		// there should only be one TorServiceDescriptorHash for each Onion hostname.
		if _, exists := descriptorMap[hostname]; exists {
			return failResult(
				"%s contained more than one TorServiceDescriptorHash for base "+
					"Onion URI %q",
				torServiceDescExtName, descriptor.Onion)
		}
		// there shouldn't be a TorServiceDescriptorHash for a Onion hostname that
		// isn't an eTLD+1 in the certificate's subjects.
		if _, found := onionETLDPlusOneMap[hostname]; !found {
			return failResult(
				"%s contained a TorServiceDescriptorHash with a hostname (%q) not "+
					"present as a subject in the certificate",
				torServiceDescExtName, hostname)
		}
		descriptorMap[hostname] = descriptor
	}

	// Check if any of the onion subjects in the certificate don't have
	// a TorServiceDescriptorHash for the eTLD+1 in the descriptorMap.
	for eTLDPlusOne, subjDomain := range onionETLDPlusOneMap {
		if _, found := descriptorMap[eTLDPlusOne]; !found {
			return failResult(
				"%s subject domain name %q does not have a corresponding "+
					"TorServiceDescriptorHash for its eTLD+1",
				onionTLD, subjDomain)
		}
	}

	// Everything checks out!
	return &LintResult{
		Status: Pass,
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_tor_service_descriptor_hash_invalid",
		Description:   "certificates with .onion names need valid TorServiceDescriptors in extension",
		Citation:      "BRS: Ballot 201",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABV201Date,
		Lint:          &torServiceDescHashInvalid{},
	})
}
