package generator

import (
  "io/ioutil"
  "fmt"
	"encoding/json"
)

type GeneratorConfig struct {
	RouteSetType string        `json:"routeset"`
  TimeZone     string        `json:"time_zone"`
	Segments     []interface{} `json:"segments"`
	Labels       map[string]interface{}
	Allocations  map[string]interface{}
  Name         string        `json:"name"`
}

type TomlConfig struct {
  Definitions map[string]Definition
  DB Database `toml:"database"`
}

type Definition struct {
  Files []string
}

type Database struct{
  Driver string
  Dsn string
}

func NewGeneratorConfig(c TomlConfig) (generatorConfigs []GeneratorConfig) {
  for k,v := range c.Definitions {
    if k == "utc" {
      for _,j := range v.Files {
        filename := "config/"+j+".json"
        file, e := ioutil.ReadFile(filename)
        if e != nil {
          fmt.Println("Cannot open definition file")
        }
        var gc GeneratorConfig
        err := json.Unmarshal(file,&gc)
        if err != nil {
          fmt.Println(err)
          fmt.Println("error rendering json")
        }
        gc.Name = j
        generatorConfigs = append(generatorConfigs,gc)
      }
    }
  }
	return
}
