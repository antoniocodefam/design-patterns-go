package main

import "fmt"

type configManager struct{
	configs map[string]MotorcycleConfig
}

func newConfigManager() *configManager {
	return &configManager{
		configs: make(map[string]MotorcycleConfig),
	}
}

func (m *configManager) addConfig(name string, config MotorcycleConfig) {
	m.configs[name] = config
}

func (m *configManager) listConfigs() {
	for config := range m.configs {
		fmt.Println("Config " + config)
	}
}

func (m *configManager) loadConfig(name string) *MotorcycleConfig {
	config, ok := m.configs[name]

	if !ok {
		fmt.Println("Config " + name + " does not exist")
		return nil
	}

	return &config
}