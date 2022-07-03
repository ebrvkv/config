package config

import "github.com/ilyakaznacheev/cleanenv"

// Get fulfills the passed struct with configuration options
func Get[configStruct any](conf configStruct, configPath ...string) error {
	if len(configPath) > 0 {
		if err := cleanenv.ReadConfig(configPath[0], conf); err != nil {
			return err
		}
	} else {
		if err := cleanenv.ReadEnv(conf); err != nil {
			return err
		}
	}
	return nil
}
