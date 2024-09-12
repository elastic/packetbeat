// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package vsphere

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "vsphere", asset.ModuleFieldsPri, AssetVsphere); err != nil {
		panic(err)
	}
}

// AssetVsphere returns asset data.
// This is the base64 encoded zlib format compressed contents of module/vsphere.
func AssetVsphere() string {
	return "eJzkXM1uGzkSvucpCnPZZOHoAXxYIJtgJgOsM4PY46tBsUturtlkL1ktQXn6Bcn+k/pHUovdsTE6BIEssb76L1ZV6yO84P4WtjZP0eA7ABIk8RZ+2d77d355B5Cg5UbkJLS6hX+9AwAo/wqZTgrpvmZQIrN4C8/sHcBGoEzsrf/oR1AswzYJ96J97j5sdJGX7/RQOTyofRiXhSU09ft9B7pXDWuNxFrv9xILr8/haBBqo03G3EdWrQ8cI2qjShgxS7pmchzb0GntE92/tvPX6sQX3O+0SXr+PsJf9fqPsAR6A0xKoBThSwU+EAVmreaCESawE5T6z5RiXw3i5bpQNIhXavU8Dey3IlujcXBrmBcgbPRjn7hWG3GMYrqCWJIJa52NcK3IaLlCxdYS+5QSqKy1lsjUNDn8rhLBGaGFXYqUogFLRnBqcECJA4SFEsqwuhbGWkWNr59gg4wKg4MoK4Sptsf2NF1Z81vnV21pouss6OoO5aVe3obZq5B+hCfQVcHWnTpAD2mnzcsbsoFvAfHrN4MS6CWWMJbp+nUygqvvkMM6onqNJV3OcsYF7Vcbg7ha76kjvkHVnpDZrwYR/IFOak4eNeJj/YVS4Ra65DswSROTUXE+uBPjAy0sJlFx/mUxmQlmzvuzhOVMYvK0kZodf+AE2D/RcFTEntGBrYHWdMHTHYCdhy/3FyHCvqxq9DHl+9mXNoXBBKz4gRGEfIi1sOw5roc9pAgscxHZgXUgnbxrEfPCGFQk97BGoZ5daWOLzBsRaANaef5QkUM3jbnc6K1wdVNkW//U4coiAbMiQQfNWY/jgrUsSxtgsBWGCiYhYzwVaoLKNtZBjpehfxUS7d4SZv6U4Spt1ZdZp4qvyabuaDtCtS9fXsHucY5M61JpiXLo8A7WT9Egi2uq3xn5WOEOdl7mLBI2RmfXRg9LjIqIurn3550Ia3W7IZvJIB/vBkxhm81tjI93Y6a4M4LihufKMtzJtWWQnmIXwxXkjK2cxp0+d4iMlZRxnfrhoLly+sYzS20QasQ6s3ofErZOTrlObkCooLQJSccgPtmc8cjlgQftDgd/eMj48WA3Njyz457TW+uYyFAPq8I8T3hLmv6aUKeADXRpLr4J+oZEq+Ea0k/VLkK1FUarDBWtrrsq5kW4LmTpj5h3ms9//uWEdff1x4BH50V57YtHOPjGOZT9vTgeYX8rPkV3Mc9KDj3rp/tMeSfxpenf9EqV4FZwlIxQ8f2KbdEhWGXRblbhwBaTJDIEQUDsBa2rT7jOcomEwBTcf77/3b2RMZWEsJKneys4kxCAOjvOhJTCItcqGTIhx1nFU+nLsTj6Kp5TtATl8bBlskBg3GhrvaE74tY3HJya6mvJ+cCj3xkqJajaB0JfxxFq3R2EfQFkPIUAcaJBxS9sh+A7SoSqrnHj4I/f8bsvMh93HD4vcmfcXkxg/MSlBfqK+ijDTJu5Gqt3/vDr0c3VT42Fb5426tXo4t5x6onS6BDnKXIxcDDAGKU8dxmizkcyTwmiqqHTyQKkwrFmKtmJhNIVGaZs5iLfPEnChSRgBLtU8DQ0EXbMQossJIVx1YhDLhSh2TK5ggd31TOYG7SoyPq/1qirNlDFzuUO0JGDQY5iu6wQKpqvQwKzZarKP9o6dzmrZj+krSqc9UpjMn854y9I9sDO5/HCklKbzXEnrKDVpjcvrprMWaDQGG2WFJvvhQSql4uwRLuQJNtQp0nVu1pMkKFsUWNQz0OYFZIEZ5YWUH1Na7rvNHBn1n0X62Wab8l1duV3sJ4HMTE6zzFZQPHtTFCJs6J+EdSZlV5nqKtALufto/BiT+QeUgTtChwpIUUmKS1J+Nq47lK4NH6imdtgLHISA1eTad0wL/vy2ENcZePER8hLYf7sad9s882jDYSBG8Xwel606V15u/tofC8xmbKQzThHa8Va9ttT/3rrCUl111oZgWOOQKuqM2eBN8s4pMGvmSQYple9BXuzn+i+t4rtqV3UTqflYkfKLCRIyH1ZXkMPoxhhbTGy9tEL8ZWtun/zQ7fyuuQXShyPKvDre31DWoFFN0cvgjZPI8eJavRiOVcq+Q2VyyNVGikhfBuDEHfR6WGfH3MO73H1vLqpULy/J6YSZpIPN/BFWDJiXRAmjyFg5trQh6FY/ea8pJMFXqXDTEHZ7FFZXRiOudZyjvT1vTwfHIGLh9snxtjR7CkMILvTYYigp38zK7gfFedoPPuKow8cwpLg9qZ3hNxGl2DGVN+4/nXA82NSidnxqm8bX59SYEQx7fP7+T7J+5n8Q2da7GRRuUXF/Q1Qyij40uNdeF7Hc+1c7AaYBc4kL0KJtt7Dl+/3XYHBUeIQfECnsXkLxA4Ya2vNr+KeWLo9HOfEdbxZrOa5QEu9TdxDGtdL+DdHqRQMFCSk+BEi3JET9Y6oapYG+7jHjPkbyBJ8+WnWrGwtFUGWkNZhDAlia+QTLX5cYilljFmK+zLKlBZzfqCJZUW5EVs2sCJxktcz+HQ8uvLW+cCwioWFZ8NUWXx1mA1lj9Lqo02ZuxanjZf1sX8e6+Gst8B5i936W6Uk1kg7RNVdLb9IEjvmG39vQBTVNfwf1qMOy73TOV8zKfXgEzyxeG8/UlVShMSIrf+pgEt9+zzG9BZNimz4+ey3mBqa+Pi+ktCHYDs7ISWssV5061oPW+stgiAbipzmxP8VwvhAG6HMqFYiF0sfjYQqhfd41+DiJmkwhTPAx7vrsgjXWW7QDq0rneT6TD9quG0IVopsmGzvpT7eTXWm+R7oUK0+mWnf9N/u4CVil+ondH7KKcWILmDhNtqFEGs1hK+V3zrZmRpBNHjS5U9ORO30i+GEdpXkw1MkIz934rjoCQURybfDwsgKXtRwVPaf4a5MksNrmDpiBDqm+keOhvmnBO/DUGnkAbeZH8NJF38M5ztaNNszIUR+HufTlgnJ1hJPEm8vJA83iWKsJetN2RuKsz092PiJhLV8/CzKInp8wR6so8eTrDfD+Gjbu/0RwPLCks6eQkrphajX/8XOr3+EN5+u+WUiT7jMZa9tq/1N/dDcwVyxeV6v9QxsNas7LFh+5mAx6f2VuXNwzvW7VUspSPX8etXr089UlHPdx/44vIsxlRzcx84aMEVffGtW3sp1smblbUAqiuU21RErfbXRq38OWkJv6IbzwndHBl+QmJC1HVe82Pq3DpazXt8gqG31fCTNf/4fAAD//4bsVMY="
}
