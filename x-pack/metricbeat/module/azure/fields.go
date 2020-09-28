// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded gzipped contents of module/azure.
func AssetAzure() string {
	return "eJzsWL1u4zgQ7v0UA5cBkgdwcUBw11xx3fXCmBwr3EgkQY6y6336hX4oUyLln7WySIC4CBCR/H7Mb4aUH+GVjjvAn42jDQArrmgH2+f2/+0GQJIXTllWRu/grw0A9HOhNrKp2iWOKkJPOyhxA3BQVEm/6yY+gsaaTuDth4+2nepMY4cnGYYpTAzFqqbSodLjSIB8peN342T0PAvcf/5/IXjubRA7JTK4gdGRN40TlBDGHq6gCzjgLQl1UBRLndudWD5amgwsO74gI0hpl4M5AEeystRziytwj19Dij0axtI/PWRpzf4bCZ4N9Q+Lc8KiKUWN1ipdDvO3D9vbTPSxGW10YpPQtH+9xUxqbo7pCAWeKhIc5Saw+WY/QhRK3s8ZA8K//ySEUtWkvTJ6uk8Le3Rhf67dmzOaJ6UciUuE91P808PNug+VwYXB31X9Xy8GHHHjNMlULlpbKO1V+cL+Yv8Z+/CeGK9TgNZWSmC3zRmec10pWjqNHNzfI2JZUfpies/ouJDI+daYGbiCtwNN1wZO0nJtRtJymS8fVrihGc5DC6u2wtbAoDGbXM9TZ+8S2znJV2YnnF+ZzThYCu1eVZXS5XtEdoAG1BIajyWBJEZVXdtuReMcaXFcN7RZ1EBpHTH+KITx8+1Y3qgLnH/PscYLBVl0XJPmon2wrssTOCTgJ7NGNkkw7yQeQNMbfGDtklB0Jbxigfb5utAYemrS+TZ4D/HZ/jDUQWHJKSNX78KhzHp4mMGP54AQpnmXqI1l3jMs5w0FN1itWlzdKdTBQgIbaA/GkUDPqxMH4GXqPnArn0VDK43Xjh3T1LZhKt7q2WE0/b1iquKPvwOcVJ4xUHiBFXniz+FklDuOZ7xpRqXJta83jFrQx7Y2qIWg9owhR6Xy7I6fw1BQe8aQJ/emPssGDWITOxIZ9+ipGDrzR3YTtIZTJDFTG63YuOvvqeWVt9QUeH4nXS5pz8Zh+aFjMkgcv9dfAQAA//89MfiH"
}
