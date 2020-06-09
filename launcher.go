package buntdb

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for buntdb service with config file namespaces.
type Namespaces struct {
	Buntdb services.Namespace
}

// Launcher represents a buntdb launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	s *Service
	m sync.Mutex
}

// NewLauncher returns a new buntdb Launcher.
func (s *Service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		s:       s,
		ns:      ns,
	}
}

// Up starts the buntdb service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Buntdb]); err != nil {
		// Add namespace key when returning error with logrus
		return err
	}
	return l.s.Dial(sconfig)
}

// Down stops the buntdb service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.s.Close()
}
