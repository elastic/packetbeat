// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package oracle

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "oracle", asset.ModuleFieldsPri, AssetOracle); err != nil {
		panic(err)
	}
}

// AssetOracle returns asset data.
// This is the base64 encoded gzipped contents of module/oracle.
func AssetOracle() string {
	return "eJy0VU1v2kAQvfMrnnJJKxF+gA+VaAMSEglSEkXqKRrwGFZZ77r7kYT8+mo3NjFmTalKfcDW7sy8N28+uMIzbzNoQyvJA8AJJznDxSIeXAyAnO3KiMoJrTJ8HCMnR0uyjFLnPvrZjTbuaaVVIdYZCpI2nBqWTJYzLNnRACgEy9xmAwC4gqKSW9DhcduKM6yN9lV9koLfoYanHbMd19FSsq1oxburVPwDjKRfN5Hm6YK3CYTfvYsG/pm3r9rknbs9Eg87EvthkkChGk+FkGm0brIHWNdNMUMIC6EKbUoKdx2vVLJtHqKb0icJqdU6cdmXs1fil2eInJUThWDTi5kQGX8Q+gB4KiSHONAF3Oaju5HQ8xPVivd+1JTgB5j34p3bUoOW2rsIn0TuV7/Nq6S30XLr2CatGoLHLPZI3tCbKH0ZGcWcIdQR74bFGRlEmeqyRBYnESgM8zl1+NZjBDxsamHaHOmFhAzNjEIbeMsmttQoWtPKeZKHTqVQ3sJthMULydD8NrjmcBrWaVObGZbkOEfJjmLQ/h515Hwqvb+ajXTql9PYEREhw/hxPJuPv88n0Aaz28fxfHaNL81HyaRCYvTZ3VC+XLIJKSrtQk295WFUi9+orCQPQU29kwSotaY/Yr+SRW50VXH+9bJXFK2kUPx0Hm3mZB2elX5VddxakYM9MsJCxWrf/7xfTKfD8H6Y3AyxmE7ns9vJEIvb8A763U1+LB4nd6PjK7/7/4TT131rzdbL1tKa/2HtH52245PWR8xpRzIGrjnuRmq42wD9nR/G5syE4iTGz1MIRPr/RZKwNtoMfgcAAP//zNCXbQ=="
}
