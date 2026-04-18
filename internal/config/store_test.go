package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSavePersistsConfigWhenFileIsMissing(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	cfg := Config{
		SrcGlob: "/src/*.txt",
		Dest:    "/dest",
	}

	if err := Save(cfg); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	got := readPersistedConfigList(t)
	want := ConfigList{cfg}
	assertConfigListEqual(t, got, want)
}

func TestSaveAppendsToExistingConfigList(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	existing := ConfigList{
		{SrcGlob: "/src/*.jpg", Dest: "/images"},
	}
	writePersistedConfigList(t, existing)

	newCfg := Config{
		SrcGlob: "/src/*.txt",
		Dest:    "/docs",
	}

	if err := Save(newCfg); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	got := readPersistedConfigList(t)
	want := ConfigList{
		existing[0],
		newCfg,
	}
	assertConfigListEqual(t, got, want)
}

func TestSaveRejectsDuplicateConfig(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	existing := ConfigList{
		{SrcGlob: "/src/*.txt", Dest: "/docs"},
	}
	writePersistedConfigList(t, existing)

	err := Save(existing[0])
	if err == nil {
		t.Fatal("Save() error = nil, want duplicate error")
	}
	if !strings.Contains(err.Error(), "config mapping already exists") {
		t.Fatalf("Save() error = %q, want duplicate error", err)
	}

	got := readPersistedConfigList(t)
	assertConfigListEqual(t, got, existing)
}

func TestSaveTreatsWhitespaceFileAsEmptyConfig(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	path := configPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	if err := os.WriteFile(path, []byte(" \n\t "), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	cfg := Config{
		SrcGlob: "/src/*.md",
		Dest:    "/notes",
	}

	if err := Save(cfg); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	got := readPersistedConfigList(t)
	want := ConfigList{cfg}
	assertConfigListEqual(t, got, want)
}

func readPersistedConfigList(t *testing.T) ConfigList {
	t.Helper()

	data, err := os.ReadFile(configPath())
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	var got ConfigList
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	return got
}

func writePersistedConfigList(t *testing.T, cfgList ConfigList) {
	t.Helper()

	path := configPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}

	data, err := json.Marshal(cfgList)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
}

func assertConfigListEqual(t *testing.T, got, want ConfigList) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d; got = %#v", len(got), len(want), got)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %#v, want %#v", i, got[i], want[i])
		}
	}
}
