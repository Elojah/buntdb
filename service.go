package buntdb

import (
	"github.com/tidwall/buntdb"
)

// Service embed a connected buntdb client.
type Service struct {
	*buntdb.DB
	Config
}

// Dial connects client to external buntdb service.
func (s *Service) Dial(cfg Config) error {
	var err error
	s.DB, err = buntdb.Open(cfg.Path)
	if err != nil {
		return err
	}
	s.Config = cfg
	return s.DB.CreateSpatialIndex(
		cfg.IndexName,
		cfg.IndexPattern,
		buntdb.IndexRect,
	)
}
