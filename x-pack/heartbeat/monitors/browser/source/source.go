// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package source

import (
	"fmt"
)

type Source struct {
	Local      *LocalSource  `config:"local"`
	Inline     *InlineSource `config:"inline"`
	ActiveMemo ISource       // cache for selected source
}

func (s *Source) Active() ISource {
	if s.ActiveMemo != nil {
		return s.ActiveMemo
	}

	if s.Local != nil {
		s.ActiveMemo = s.Local
	} else if s.Inline != nil {
		s.ActiveMemo = s.Inline
	}

	return s.ActiveMemo
}

func (s *Source) Validate() error {
	if s.Active() == nil {
		return fmt.Errorf("no valid source specified! Choose one of local, github, zip_url")
	}
	return nil
}

type ISource interface {
	Fetch() error
	Workdir() string
	Close() error
}

type BaseSource struct {
	Type string `config:"type"`
}

type PollingSource struct {
	CheckEvery int `config:"check_every"`
	BaseSource
}
