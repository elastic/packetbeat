// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]bool

func init() {
	// Packed Files
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/metricbeat.yml
	unpacked := packer.MustUnpack("eJzsWN9zq7gVfu+fcZ87XRBxWjqzD4YsAuyQa5xIQm+W5AC2hJkYG0On/3tHGGPs3Lu73e72odOHjDMgHZ0f3/nOJ/7xZV+u+Q/rQpS7vKj+0ij55e9fmPIq+rpLYzzZcmiXrFjMEmBuX3JHMhXXDMiDmJYtJSFY4Ui+5FMQtemDq+SBW3HG4Ok9AZlMQNW+pLsqgOiB4uidKrmnr7tZ4E7TwHcyAdPHwA2vNt3g/A56DbWeH918mgbutArcEFASlrTY5m7hHLkVv6/wpBXQ27On3Wy+dOQaog0BtGTwrd/n1AmJdy/5NKfXOPL+XcaV0L5p++l8Oc0HexcfXCMNcqMSxCm48rZ0efucK9QyCzUJQG0wsjM6a7Tm7bBWsg7SsuYKvQs8KVkRGQk+7X+H/Jx9KtBhHAMv0J4SvcZIqZIPxIoMrlDGXnfp2jJm3XOSGVx5G+HZ+xWJDAKkQbE55IVDz1g97VIKPSMh0VGQcEPJc++DXms8BlAeuI8MXp/3CPi38TNdG7XCJ4OS8F0oby8wutoHaE9xZDAr+GRnBdDkJXeMNXFkXzOQ4JN5jqnHBUQH6jvHFZ4YbpdHWjI/llzad2udluKTTKz4yDe7dIUntSBx27/7oGT7qG0lONxTvOjzTzPmI8nb3VB34YcmXXQxHYVCjbDCUlzxpriyq5fcKVnhmMJ/PjxPRzlxgwv25BpGkvuLm/pRdZKU9Gd3WHKOgsS1IIuZW4gdxQ+PwU/oJLC3500q12bV5YYQYx+4YoQRXgZuooKfsowb3mG9TIvedjFbTlWYO0FCopcEm5JbTpaAt10IbFP4jilcfrYFvVZAueEAZVxFu7DZzr78+UwV77lcs/XqE1Vo6OJQJmSgCl12mSiUiWl5brfcYUFuekFep0ERSeGjeq7kni0nkikvZxBtv2Kd9kh2a+7XFrFkxNknJJZzNZTLpsrbc/CWz91pPn87/zLsHRIsJMPoINxJxUAsv5K04tDbrBpTMYikcIN94AZVvNS/YZXgSUYBqiieGGP7XdmXN2v3DIhihSfFXJ2kUGj/FccyKVARSOMXW2++dEqmSg1HTWNbStILhLqSvuSOtt+eIS0PK4V0i7bCD2WCzXfuh0dNOxzYzaWVGJi8J8A+UHUqE2txbieAGuHZGS1iyft1wpe1rheDdtG3p1z7cZNg3f6nsqNWovP8fEejzjYhccaBbXIVyaFN9drXTy1lsPb6jFtxQ7FX9edVCZne2+5bJrrEeLHTCHy6acEE2PUa2d2IEdB+Z1C24mk3tM3l7NGZ6fAO2gcCoiNTdL/CkfGSOwcG7PrahlNNiRtKHKPDatHXjTx3mFrhRfdL8SRLdEt1+AlrruwNJVHLrLDV9b/z1WDmmV4vtbqhT99pGIgkt6IjL55/axx6nEhexO+JQopZoXTTf6u238t/KSDSlHb//Ka+Z5yaGX+6G113lH7ti56+O98mLYOeMYw998avEf2f/xj2ag5Rc5NL/VzjuT3jerxe061QmgPsjo/OvcVn470BRKDreSsyKJSHLl4Yl+zuDEqoZMXiKGBUa98Sa/odO2jLn3apwHF954s++0ihvVkB1OieYiD6oCS4s3M60sZuKI5Lbtotg7al43rJnfOz+nPscyuacHA60g4zsu3ykP7R4/d0pPVF+tzmbL50TOpPr6MTRvV1X3xMQDVw0nx5G5e2tSbRqE5Gytvn4X+mkEHV6Siu63fCj+tVER1H5x+f28RMCNquhrFqpAkpTa7QGRMwM5Pr+g2DtkkhPQot4YrtsIcC+0jBSXIf5dxC+TWOKqOqyv7L0mPDLGdyqSEtwuP3+JeOZnI/i/Zi1DvX+ocjKTye27tPUuTij8YH1zGpt5kAmWSbXbqEXrs888suJBUbnX+4+EeIMZInk4xh1OqZTJdpwSxkaAyGTZ2GAO0TEhkrHLUUe00C0mLuTguudD2fi3nnm/hIMP1Ilp1s6WamgF67cnnppj/+eJEsal195PwbouUVI4MruelFyoZhPThNKfywTEAvZkjYEQtuhuHfUhKb3J2UDBqHYRg8mSrBp/ZOKFzWal1dd0T3S2JHmRlTXkGxqcnmwLC9pa/mw5xorbavei33c2Lnap/EjcB3wgjaBQXyQJvJviO0J3NLcWjSJhTfuludgW5qvXt7PxndDwJlHqnfkfuBurZuXj3cDmts7oOBRB1D55KShb7TaEFmMojauVocNWHpQTMvZMXcyXZFosuA/T3EVEVJ3Ggg/aGCqsfS/0XV/7So6u7qwg8zjbcuT57diaQhn/fixxrivb03+3fP/fg4ujveEDax+oFg2taKxDtihZIC9DAILljJ9evujMVFT+7q7THwHg6zxh7wH05/dvD/p2Jh4J3vCga/46MuNi3k7r8z3OH9E0513lkRdRdAYolSwOydK1RQktU3OLjUwI8nHL6N69/oXiV5+sPb06njh6/5w8ds+TlHZzv6jPQxcOOxEK0CP5T9DBnbvtz/P4nW8zeRniMXw1CtElxJAryGK2/yTRxfZpEVSXHjV4eVwWc65Kz3DduAIrsTF2M/+jqN9l0HOleoYhaVBHRYGsf1bZHyq/bEmR7ItBcgXe9e1kHaMPC970e3M1jv5d2F3JT3woeNZ7ieyX3/jPN0mTU3Qm1z6ddIMog2AtrNGGs3dn+lELp+p9Hi5vwtbeaiv2JwE0830ztsm6b927/J1NqHqyByp8UKeGoFfur+7y5Elp4taTFb7H788s8//SsAAP//8hto4Q==")
	SupportedMap = make(map[string]bool)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Name)] = true
	}
}
