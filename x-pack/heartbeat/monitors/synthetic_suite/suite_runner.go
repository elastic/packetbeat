package synthetic_suite

import (
	"context"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
)


type JourneyLister func(ctx context.Context, suitePath string, params common.MapStr) (journeyNames []string, err error)

var journeyListSingleton JourneyLister

type SyntheticSuite struct {
	rawCfg   *common.Config
	suiteCfg *Config
}

func NewSuite(rawCfg *common.Config) (ss *SyntheticSuite, err error) {
	ss = &SyntheticSuite{
		rawCfg: rawCfg,
	}

	err = rawCfg.Unpack(ss.suiteCfg)
	if err != nil {
		logp.Err("could not parse suite config: %s", err)
	}

	return
}

func (s *SyntheticSuite) String() string {
	panic("implement me")
}

func (s *SyntheticSuite) Fetch() error {
	return s.suiteCfg.Source.active().Fetch()
}

func (s *SyntheticSuite) Workdir() string {
	return s.suiteCfg.Source.active().Workdir()
}

func (s *SyntheticSuite) Start() {
	panic("implement me")
}

func (s *SyntheticSuite) Stop() {
	panic("implement me")
}

func (s *SyntheticSuite) Params() map[string]interface{} {
	return s.suiteCfg.Params
}
