package plugins

import (
	"fmt"
	"path/filepath"
	"plugin"
)

func (m *Manager) LoadDynamicPlugin(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	p, err := plugin.Open(absPath)
	if err != nil {
		return err
	}

	sym, err := p.Lookup("Plugin")
	if err != nil {
		return err
	}

	plug, ok := sym.(Plugin)
	if !ok {
		return fmt.Errorf("invalid plugin type")
	}

	return m.Install(plug)
}
