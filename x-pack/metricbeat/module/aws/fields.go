// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
<<<<<<< HEAD
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
=======
	if err := asset.SetFields("metricbeat", "Aws", asset.ModuleFieldsPri, AssetAws); err != nil {
>>>>>>> update fields.go files
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
<<<<<<< HEAD
	return "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"
=======
	return "eJzElsFu4zYQhu9+ih977uqQow8Fukbvi7aLPTpjcmwNTJEEZ2jB+/QFZSexU2ebBlGqoyjOfN8MRfIz9nxcgkZdACYWeIlPNOqnBeBZXZFskuISvy4A4J5GvceQfA0Ml0JgZ4rfvv+JIUWxVCTuMLAVcYptScM0tgqp+pHM9d0CKByYlJfYsNEC2AoHr8sp/mdEGviBpz12zLzErqSaz29uYF0HuQzE7u7x3a1gLwY8Pffs7u7hUjSSqLCeH+2sJ8PIhaGuUGb/zPd788XYi+ufAtyoknI0bI7TxN9Xd91F/udSl2Iu186SUeiys6svHjTVUWC/3oZEzz/4iXF7/uoZmYvjaLRjpC0ohOTI2DdEuDTkaowaxc6FoMJwtRSOFo6QiKqMFKeKSVSj6Lh7UcQV9mLrqrTjGVxiHTZcmsfq6zeckik0nyt/yYhtKtNX1STID2ph/5V7Q6HNnZWcqUT2VwKnwscn9p4U5Fyp7KHS3ohhJEWgGl3PHqlAjYqxf1lKa8mh6voD5c4pr816OjA2zPGpUxRRY5BB2kp81B57jmjTVl+/raYIX07MOFCoDFH84JJea6xr11PZsZ9XeXK6Kd7+pZgMmcTDpzE29X/2/xdQ9OcNxvqqkOhqaTUi76VRUMBJ5bZ6ZBtT2XcSu0xuz6azGp9zoLBjObTFGNu+8oABicZlS471+U/5c/xU7UP5pw07VXsvfond5mg8L/yUYZbSfxT7e5Xdi+4ldYXJz8L+5VxpOl8IGuvjVqWWCuOQQh1YQQeSQJvAsPR68rGI8YzoLb5xbEzvzj5VPeX3Bl+lIQduh8JU95S5TCe3vr0F7Q5DbZd2shX27T4kybflaDK8pkFzWk4ZLjXf2Ku3SKqRVe1cz26/3pKEFw7KkOLuv/n9wTkV03aeW8/lmrTdbTKpsscmWX89eGLCxDSdim1Uj2o8XI/J6UYaSA2DxGqvl1yf4n2w6xwiD3n+B5XbHbsl83cAAAD//7kBWzA="
>>>>>>> update fields.go files
}
