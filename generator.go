package main

import (
	"flag"
	"fmt"
	"math/rand"
  "github.com/jinzhu/gorm"
  "github.com/BurntSushi/toml"
  _ "github.com/go-sql-driver/mysql"
	"time"
  "generator"

)


var DB gorm.DB
var config generator.TomlConfig

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()

  if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
    fmt.Println(err)
    fmt.Println("Missing config.toml file")
    return
  }

  generators := generator.NewGeneratorConfig(config)
  generator.Build(config.DB,generators)
}
