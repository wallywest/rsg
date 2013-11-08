package main

import(
  "fmt"
  "flag"
  "io/ioutil"
  "generator"
  "time"
  "math/rand"
)

var option_file = flag.String("options","","options json for generator")

func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  flag.Parse()
  file,e := ioutil.ReadFile(*option_file)

  if e != nil {
    fmt.Println("Cannot find options file")
  }

  config := generator.NewGeneratorConfig(file)

  generator.Build(config)
}
