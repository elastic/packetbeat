// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package openmetrics

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "openmetrics", asset.ModuleFieldsPri, AssetOpenmetrics); err != nil {
		panic(err)
	}
}

// AssetOpenmetrics returns asset data.
// This is the base64 encoded gzipped contents of module/openmetrics.
func AssetOpenmetrics() string {
	return "eJycUsFuqzAQvPMVI94hUpTkAzi8X8i7P1WJwQu4sb3WelGUv68I0JK0PbRznGHYmZH3uNCtAieKgVRckwtAnXqqsDl+sJsCEPJkMlWoSU0BWMqNuKSOY4W/BQCsHAhsB08FkEnVxS5X+F/2qqncoczZly8F0DryNld38x7RBHoOM0JviSp0wkOamS9ujzivvGc0HNW4mLFEaoUDKNrELmqG9kZhhNCy93x1sXso0LIEo4f55+ukI/7gKJYELsOFxKImKnoS2sGbmnzG1XmPYLTp0TrJuoP2BKE8HbU81Pd9JiztJ/Nh+y4s9bl+pUZX9EScJvVCtyuLXcnfTDTin3Ag7WlYppmvfgozT/HjNE/dHtRTMCm52M2fltvyl6FXaR+f5lsAAAD//0qh20E="
}
