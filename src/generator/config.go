package generator

import (
	"encoding/json"
)

type GeneratorConfig struct {
	RouteSetType string        `json:"routeset"`
	Segments     []interface{} `json:"segments"`
	Labels       map[string]interface{}
	Allocations  map[string]interface{}
}

func NewGeneratorConfig(f []byte) (generatorConfig GeneratorConfig) {
	json.Unmarshal(f, &generatorConfig)
	return
}
