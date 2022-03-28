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

package kubernetes

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kubernetes.
func AssetKubernetes() string {
	return "eJzsXU9z47aSv8+nQPnkbDmqrT1Obb2qF8/LPm+Sidf2JIetLQUiWxJiCmAA0B69T78F8B9EAiAogrLHlg6pjG11/9DdALobjcb36BH2H9FjsQJOQYL4gJAkMoOP6OKn5ocXHxBKQSSc5JIw+hH97QNCCLV/gHYgOUnUtzlkgAV8RBv8ASEBUhK6ER/R/14IkV1coYutlPnF/6nfbRmXy4TRNdl8RGucCfiA0JpAloqPmsH3iOIddOCpj9znigNnRV79xAJPfW7omvEdVj9GmKZISCyJkCQRiK1RzlKBdpjiDaRotTf4LCoKJhoTEc6JAP4EvPmNDZQHWEd+f7+9QSVBQ5T151Ck9acLzYTH4a8ChFwkGQEqD/6kxvkI+2fG087vPGjV51rTQ/AVkkLptWYkvCg4CFbwBOLhuCspQ4qstLsARLGaE4OLfA9GwvL4AJAmiy6TrBAS+JVmKnKcwFUjne+8uJ6Ar+LB+ufDwy3qkexZJksjikLz7JHs86QSqFwqRvHVUGHQLFCPRRdLyvdLXtB4MH4HuQWO5BZqHqgQIFDK96jLqAvmkdAutwlIfiI0VatrRX1AJbuc0bhrVE0SbTFNM7VKGULxoumu3RORqEVdk0RrVmsmYJl4Ai4Ii2gaFcEGRX+YXQhacgeb20QI9SSxEe4y34Hcsoj2qCemhWhv0ExENMNmxF2qNducswSEsHK0GaJtvzfpJXmxEJD0fl/TTFmxyrrrXm8g17dfkICE0bSLrOW0gx3je7WtkxSoXKz2rWfW55sxurH8svTLPiLXlw9Q/aD+CBGKap4VhiGIT4TLAmenRFixHAK4TsWC5UAXCSt6q98gtAPWn4vdCrhacRVBtCYZNH/AuFuNQmIuIY1gNPelwSBBaAJ6iamMu+ZhnQAqEIhm/c2+WnDt7S8KsciBJ0AlyWDxb84RstWfkNgUUP5iOUYO9ZyvQaAdSTirphNq4bh1YhuGKHYT9ePHlRS7IsOSPAGysfJBm268NTRNSe9QNf1BIIL8C8qZHVPTY0ArBKPUakD2aTXGgnSAcaSKDZhzaFiR92AQOaMCXlS9JYQx+u2Dnl/BJspgDfeBxlBxBcVOqu/0x7epemDWnaZMgyx8/J28HVttnfhAWCBLlqUz5HhOXkRvwZq7MZllWAJN9sdYsk1boiZ4pUxUISj/TUrHydyTBiHFM6EGEx0vmFWRPII86ZZTsUZbIiTbcLxDJQg32FBXYgyKmmapyVDlzeM5tFio6QiXPwwD8wJ6bFGHazIpOFfr2HTZ3dB1RjZbGWDqjG54QSmhm6ihSrt+JnrTUt9GFSN/Vhlkki5KuUdZydukf6VNgbDUXKzscZESuYAnlyLGstf0kKZnH2/JkIOCBmlEnjXJLvN2r6ESEzrtjMOQbkMvyhGHjiyXkuzsqdwUy+4vBhI294og6hE00ivBu/hQhvL2CyoE3oBFEK5hm1D0d53z0AbIR/VgkIzbCA8TH2JgMrEsyl02zrC2/gxI2PxcN2an5H7NOFTCp5g6t6wDvJgyJRgX7ADIgXBLw4B0gGUDjKWwyK37UotLJDiDdLnOGHb9YR12VJFOjDEo+WKBcE1T/ZutdWpIMokzjR3hLGMJlniVgfqed7AZ2RH57Y02hTWhkJbwmwx8uxReqp84JYLIGhVUfxdS+yFexjbhOeSBUf3MNsoVX7ORCxJ+wiTDdvOfvii5omEUNveGgmoUrm0tn2awKME5TojcKwfYTr1ZV6u/fA/yKa05XDZqwXsPctELe7hYiFoP3GcW03Z5ux+Pom5mD9oO2tniHJBxKMLB737Ew6VYhUByWOcckLSBWCAdnmlFSyW9l0W7a4cDB3PzOdevTSSlIJwDfuV+5i8G+pGupsMC0Kv3NkPGPMHhrAzC7XOaEuK90gX0xmbJ3f29f47UkJ8ZfyR0I8CdHHsbEvm9HCgSIMNXj9c5lVxDeelpleMNrHGRWVKy48787UNvc4CKEXJwajwO/CfjJ0OkuTlxNesOY3IdsWrqfURmd4xJXRkk9kLCbnSQ9l6cRbuczCDmHM3aZVRFLy8X1Z4kUvtiidHM8xLOsgx4eaVk0rnJdUOsuqAS59TkRQp7T1nrf+ri4RMXDav/xmP3Ge8grDb9X4xG5HtD1xwLyYtEFhz6xM8l0uVwziXS5xLpc4l0wDDOJdJ2IOcS6WCM5xLpc4n0uUR6eom0xcscWzT9zPjjXwUUdo/zmK1PgQblcJaFjNO3859Lgk3FYrWZ+3yJgq4JJWIbxZ340hALYY3TNIYN/17rRREcMOQUcrmNylNTHJw+kpMo87Xla9aFa+r2wIylsEhUwJ5IZo+vjzFceCKJ9iRi+sA6Q11T9hnsFnAmtzGq7VvmDVVkTwXNcdPBz6nE4zihCGd3e3Bm4B5ksyYBToEviFjusJCOnMyKsQxw19EbagWwbXsBaF0TgTo8PnTR6BrgD132IxJWD1swG5qUNcV1zgrUPqTnRvMbucUSYQ5oAxQ4lmUHlroCu1pXDzgQqgJbJdyfuv1g0IhkmNvAHLr2Svu63F4VF8QhYTwVpdwb45NkB+XPcswlSYoM81IIaIsFYoku608tCPU3Jd7lFpT9xcSX9lsTLuSyYkUdXVDGl0w/1ADVODUP1PJQP+talXmFZnZAisUAnjYXInrnmCUGCV9luDX8UtKpLAHStuUCeQJqEUfC8v1SMhuCdk/DohPquVNvXnR3mlIouMYKu61MjuT+sM+b01Q/R0se0mX0fo76hLbuBsIhZ1yW7UCIsOjCN4Fm7VOy5myHnrck2WrhlGsDEe3KaM8NRc08f1b7hCKMGA3FYuTccYolnq6xXypKCAvBEqJ3hWcit9455NObfQkd75E1dsChpxDkW7ACTpYOFi3NgDDqnyktoFovy7gnA/9Vka1MYt0ag937jX8sEcRTt7mKy1iTRKSeBOUEeMZDs7E+PVlG797zW9W9xxSI/7CmIBEPwL5Q8lcBSB8pkDVRbiUzgFhSSs0yDtl6mRH6GBHM3c9qHecgFJqqs5NrGyH0iWVPkC4tGOdanWqeNrn41imck/iW8/fbm6b3U2U9HnXFbQKmeD9WjcAGGMddPMwFy8N0vvlaUx4h+rgT9svNpwHeZtJiSsxnXO7Uceb5Xuf5XqfjE/9ep/ZYv/Urnee7Hfa/Od/t6H3i3e04l/D3IJ9L+F3QzwXpAwXpFKSynmhrN//6xk3wDhIgTzrfr5uu0PqChZIxIlQCX+MEFuhm3f8pIsrXlPV9jCv0TLIMraBK76mNlknl73BIJHrCWQHoj3//wysa4Nx2ehosm9Bxf604vdCQm5zYWzewB46p2BEp35+NPbygjTWHTefbOfUnUGs/ni/mjBbR+U6O+emJ531cxzHKShzNE7qwTtP7osX1WrpetIhcnS8a/7OgzuzbMWs42Sl/faZOJu79YZjBEBMUONNReHorZMajcWmwm52OT8bvIChwF0HvW5AB+wwas+y9SyHad6MmxXBwt27KKUTO0m/yEOKcQ6g/R+UQXkPk1UT357hbvia9PFj18q5O/17NaVcP2GtsLDWmgem7alqqNtem4Y3odrypupUyCohxtGMczD+uCCsSmMNQT9PIp6HnY78e8Fc5884t3eJNx6MbUL2XhOLBhHEPunNAvHz7J8SlYJ5758Tf4OSK0+Rtpon29qsvSktq+pEoLeiL2ANVOjnewHLGA/4SVnC5wfI0eNzFBkY3nK/7KYkS43qepjX9xfHmLoql+9PRl5RcDaXaFH4a5UKSrZGUcaWl2/FpCpceOeclnqlSO6RnmE6vkdKYe4QHAbP1+rp3GgS2Tzp0/4eaJ3lu5vrXsGPaJnWg+ZsmRUbmbZfUHCL5myV5IE1olHSQLnX0NQg3jDEtkpr5022QdJxVj26NJH39VELaIkVpiuSDb2txEguRt5OKD9Q04wxuhNSFENonJ1yr4WAHGuR4oLo1OHVpCW985IEXW5eWlkfGZZ1xDY/iKjKs1ZEP7JyqDG5x5AM4VZme5kZdE4poNzbncKiP0TFNHgI7GDXb4Z4mQZuSl+ljsYLSTa+c9T1NrEcLA1tbkYEI3BmGxX+/p8mtgnOnyHbeAWXr5gdDL7q60U0zDye+gLdB3Zic74PGXGec0IceCO0cIudc//GO0E00tX8uSSOD9qg3YAMhTvRdvSBHGMAAypNYg38wbpPoZQ1EsoW0yKZ1ozYyBw29c9qgz+ONpQ16N7OPZDPUZ9rwTIosysDuKytFWErY5bJPuubZrAYR2arJaqN7Tsec0zFDkM7pmHM6ZiSiczrmnI45p2PO6ZhzOsaKwdtoteRva7PqhTCmxWovFus2Nj1uk4T/gNOHpf+gKZIMAU2Nwdi3pUDYU9ISI9B4JmAX0bQZYcfkm4k5Sxc5BxWmKAS6L/NuUJ/DSG5Zilq6qKI7DsQU7dj5exThwDBNHw4UQwqp4sYY3GtSXp6NAVb2ekp/994yd4Y31h7iaS6uDUTQ/tnDMTGD7Jq0H7qMm5qwD10ux11eah8TjXGF6eiWcj3xXDe4iL33pZBYFvEu/edbLNzVp/YBdAfhq35vhqMZocuq1/gVesZE6v+RwHeEYn/5HuDU3ZfA3rc9EGWLUDOxy/fAgVQBubs4jVAJm16D+SPAlHwG32Do9as2wUzS34POH+ou1ofP7xIVBXMOVGb6ymulSnTZwL/WjXWVdq85FtufGct/wMkjW6+v0D841xcXb4ssu7Iybn5dfec7xLhhJorPLs9AQnrVSuwaU8rkXUE1B8av0K+//vITyTJIv9NKhYVTiro/estgObdUdbN0q2ifsU57GYPV0MOGrMTUDnphXT7G3MEafOhF1yG7bh6VdF3lx6Mmw/XtF91VUZQsPbOhjnxOAqliByk6Wa/BUuTzPx84NPyq+Lqs3R7sfFHr5eVxtyqry8dd96QTzuifbBXL3SipRXE2egdF4e4Guq5w9Gh0TzCnMrDSMZy56qEI+8wI4dOSQDnLSIdSc8EmUQ7thBe22vRHSUrFK6J9kr5nJIZPKJaiEDnQtNekwOe2HHA3Mx+1CREV1dnotpar2+ZbTiQ8AcJhWJ2zZItE70yihqB2KVtz/oP9tLaAaDiU0PXOWcPgBbVPEPg6E3tFeZB9CjjNCHVzHrK5TxWBhjVeS+DNlNJIEqYf6OHK71pjkhmaCPkf/z/dYViKYcfo4b2sKQUOnzS9e31l6MQrY7s55RlJcHhINbDhWEdXMTnytvrwNcGQN7Pc/mhbrtM2YKnFgnLg7UCcEFMQhHs6f00DWFE/eB5uFDx/ABlLemUEORJaQU+nXoPXEMxWsXnG9ruJD5YZrlBLMMqcz7GlC1DwBuudvgbSkostOXCCZcTA0SiN0DUbuYoMTdFJmYtPLcbW3Opp26C+FDkkU65Dx8LYnwaT5mcsWLb56QaWp9bnkaKDKvn0AdVAQsIk+8QOcxLsgdNcIZLlobvt4dNM2vnzvQmkYoRY68Bt22WrCTu6IUeAB6GDlxn2F2t85ITR8VTngFGy8MMQRZIA9OOyuEg0FyHWRdZHUyMZ1fwxfMdQFtq8xjbW23Q9u9YETraH14KdEQULYWk8e+eLKA9w1SnQkyKzc+1gstS3x1JjywEJkJLQzVh9zuuaJ4yuyabgOgvaQG1Lm2o5Xt73tv7WheM4yyAjonvsGEuIBodXL0UTq7HheOTHnqnjRePpktO0tbfJd7ZywyGxWXZiFO8gxdyVm6fA9GPH7JmK8u3R/mbZorM8wxYRXf0kmwcZuoTNAl1cc0b/m60u3K4xEcuEUclZ1r3DEQ3yr8/14V7DCF1eSF7AxRW6WONMqP9hHF38J2UU/nZht8aRB9TjzLEkPsEe6+V8HhGa+dqDzcMhyII+UvZMPXof8Jiioq1cp1CozRR/nQ/bxayS8KePJilBv6xmlh6gSyX+K6SFr0RfSd5tIwWtUsPeGHVSlYRGecBnqFCiPF9b5hyEKKwP4sUSXtnV7bZidLQUUyIeTwH3ExGPk8GyQi7Zeqkwzwj110L+ulZ4j8aZk/QUMr29+XSUSOeoijB6lM1XiNC8x2h2RLNza8ZYNakfADUh8q0xNe3wT/cEoyHzlyoy+OxqUDdUHdEo5kWRV2j9TxkYCa5ZdKdb/c1lmqZuvOmhRiGzwtGdDrtPajQCBi6IkEDlE8uKXSz/qiWLSrptapGznf7L73XB2fcvXajyWwlPkXAUsvlmTdg5e8XD+q6Jrz527CDK0lecJIyn+h11ZujE4b4yjjewTDLca1wZzP2+JII0kSYX07MnFFII4LLLJMNkN5txJhl+xSZ6+9u1xz7LISynMPiB0BTSWhhuVlVR27Kymgkz4q6taKynV/xZoeSmCdhpY501X+66HSRGcPi7JoEUCTuPGefX7W/X3ULj/ix6dW+wbJmQS9LduqtzxONTNgqeIo1ubmOE6eMYV4HAcSU3M96L6MCsLkbc1RcjboGqXWKxWBx7HyImummZibpkbaYEVlfhNTcb3qs+2m7pxrSqNmNW1tVeIlJd24y1JSZUd5HLaypRO6hbuiv/8XKVacfjerGStABsbKW7Vc8ltA3Qql9Pwwmt9nr7bMEhz3lHW2i8gnkKCg6luC6ybF9zG5SmcVdMn/v8VTCJoy0tBs0oi8t85eJ3Fdb/0ViHisa7UhqDoORQngpBii63mKd6gxKQfudrhhUnEjgcqPNuhaJ3LAtzhOXMUV+9Qn+oof6hxvqHGqz9kWjrwI8YX3neqkRZmh/O84yAQJL1I0b/P90RploOSBIr4VFRe/F7N/cVDk8+IyuEBO5ywgN43FAJnOIM3dw2Jl+N384SvpZfmBSk1iOriaFPn+/dU6BhefwwewwdsUXGcLpc4QzTZJJYf2Y4RT9UdBqDcjCdMsXrgfVoNDUAdMNVaDzFRDQFF/qagQrZpthEzeafNjr+MjNnIZdVVJqGWgwPvmBGl7AusniOfU0xmmfvE8JQssZeqNqIpKn/Q5egNuhyH7yvRtD1/k4QahwIr/Ghjoo2ZvZPjfa1tXt64PO5hIheIOzo1diHAnyx+GPIBNvoYG4jNOKQ4yuKZrbFxgINsK/DBmvLCwDWybt2067TlmQzC/vift4BGre3l3P2RARhrsLNEYdLLaXW6zNRuM4M9NHN0nJ3elRgoKlUN7DLtpR7inckwSpgrna36gTDftRVnZOsiM56Tkr7/8LS8nJsCvpVr1Y2hG4QpimquMT3Rw7UPuCV6LfaY1l/+fC78WJeFK/E0mxplCYsDzY3LRicDuGp34F/F89RJ8xaeDVMfIiByaR3+aHPxvsMIgpvKIUqaV8zDpXIKaaONogdlK/lZe6ZiqDOLy8fgH77T37e3d+HiaJ6X/btP6freoPWKZkcb2DG10bby4DBL6CeDNHwG6hRS88O681ezEHvS8WoNLNyW8e7RmCtNBsi6iOMXuUs/JFkUDmm5dvN/qJSNOoY+Y2KqL33PygjS78L9NbFo70pp2TMQby2Z9FvDx4+rwosdLToVDChLPUUgk9RscN2UFT390upLMcgjJWVgz82iQHmRw4QAsbeZzU2mtJ/DYDzDZhxNYr/DwAA///fdD7y"
}
