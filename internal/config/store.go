package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func loadConfigList() (ConfigList, error) {
	path := configPath()
	data, err := os.ReadFile(path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return ConfigList{}, nil
	case err != nil:
		return nil, fmt.Errorf("read config: %w", err)
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return ConfigList{}, nil
	}

	var cfg ConfigList
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}
	return cfg, nil
}

func Save(cfg Config) error {
	existing, err := loadConfigList()
	if err != nil {
		return err
	}
	existing = append(existing, cfg)

	path := configPath()
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := json.MarshalIndent(existing, "", "  ")
	if err != nil {
		return fmt.Errorf("encode config: %w", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}

func configPath() string {
	if dir, err := os.UserConfigDir(); err == nil {
		return filepath.Join(dir, "cloudmirror", "config.json")
	}
	return filepath.Join("/etc", "cloudmirror", "config.json")
}
