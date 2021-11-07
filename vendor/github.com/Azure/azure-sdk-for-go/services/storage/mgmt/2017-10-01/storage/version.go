package storage

import "github.com/Azure/azure-sdk-for-go/version"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " storage/2017-10-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return version.Number
}
