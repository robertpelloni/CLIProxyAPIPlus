package translator

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Plugin represents an external translator plugin.
type Plugin interface {
	// Name returns the name of the plugin.
	Name() string
	// Register registers the plugin's translators with the provided registry.
	Register(registry *Registry) error
}

// LoadPlugins scans the specified directory for .so files and attempts to load them
// as translator plugins. It looks for an exported symbol named "Plugin" of type Plugin.
func (r *Registry) LoadPlugins(dir string) error {
	if dir == "" {
		return nil // No plugin directory configured
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Directory doesn't exist, which is fine
		}
		return fmt.Errorf("failed to read plugin directory %s: %w", dir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".so") {
			continue
		}

		pluginPath := filepath.Join(dir, entry.Name())
		err := r.loadPlugin(pluginPath)
		if err != nil {
			log.Errorf("translator: failed to load plugin %s: %v", entry.Name(), err)
		} else {
			log.Infof("translator: successfully loaded plugin %s", entry.Name())
		}
	}

	return nil
}

func (r *Registry) loadPlugin(path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("plugin.Open: %w", err)
	}

	// Look for the "TranslatorPlugin" exported symbol
	symPlugin, err := p.Lookup("TranslatorPlugin")
	if err != nil {
		return fmt.Errorf("symbol 'TranslatorPlugin' not found: %w", err)
	}

	// Assert the type
	trPlugin, ok := symPlugin.(Plugin)
	if !ok {
		return fmt.Errorf("symbol 'TranslatorPlugin' is not of type translator.Plugin")
	}

	// Register it
	err = trPlugin.Register(r)
	if err != nil {
		return fmt.Errorf("plugin %s failed to register: %w", trPlugin.Name(), err)
	}

	return nil
}
