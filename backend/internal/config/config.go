package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ListenAt    string `toml:"listen_address"`
	LibraryPath string `toml:"library_path"`
	BasePath    string `toml:"base_path"`
}

func New() *Config {
	cfg := &Config{}

	cleanenv.ReadConfig("config.toml", cfg)

	return cfg
}

func (c *Config) CheckConfig() error {
	testFilePath := c.LibraryPath + "test.tmp"
	_, err := os.Create(testFilePath) // Check write permissions by creating test file
	if err != nil {
		return err
	}
	err = os.Remove(testFilePath) // Check delete permissions by removing test file
	if err != nil {
		return err
	}
	return nil
}
