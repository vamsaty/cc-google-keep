package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := LoadAppConfig()
	x, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Printf(string(x))
}
