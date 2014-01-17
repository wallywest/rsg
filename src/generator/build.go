package generator

import (
  "os"
  "log"
	"github.com/nu7hatch/gouuid"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
	"strconv"
  "fmt"
)

type Generator struct {
	config   GeneratorConfig
	routeset *SuperProfile
}

type DBConnection struct{
  connection gorm.DB
}

var Connection DBConnection

func Build(dbconfig Database, configs []GeneratorConfig) {
  NewDBConnection(dbconfig)

  for _,c := range configs {
    g := NewGenerator(c)
    g.parse()
    g.routeset.toJSON(c.Name)
  }
}

func NewDBConnection(config Database){
  db, err := gorm.Open(config.Driver, config.Dsn)
  db.SetLogger(log.New(os.Stdout, "\r\n", 0))
  db.LogMode(true)
  if err != nil {
    fmt.Println("Cannot connect to DB")
  }
  Connection.connection = db
  FillCache()
}

func NewGenerator(c GeneratorConfig) (g *Generator) {
  gr := &GenericRouteSet{TimeZone: c.TimeZone}
	r := &SuperProfile{GenericRouteSet: gr}
  g = &Generator{config: c, routeset: r}
	return
}

func (g *Generator) parse() {
	g.parseLabelOptions()
	g.parseSegments()
	g.parseAllocations()
}

func (g *Generator) parseLabelOptions() {
	for key, value := range g.config.Labels {
		if key == "count" {
			g.makeLabels(value)
		}
	}
}

func (g *Generator) parseSegments() {
	for _, value := range g.config.Segments {
		g.makeSegments(value)
	}
}

func (g *Generator) parseAllocations() {
	count := g.config.Allocations["count"]
	adrs := g.config.Allocations["adrs"]
	g.makeAllocations(count, adrs)
}

func (g *Generator) makeLabels(value interface{}) {
	v := int(value.(float64))
  fmt.Printf("Making %v Vlabels\n",v)
	labels := make([]Label, v)
	for i := 1; i <= int(value.(float64)); i++ {
		u, _ := uuid.NewV4()
		l := &Label{Id: u.String(), Vlabel: "test_name" + u.String(), Description: "test_description"}
		labels[i-1] = *l
	}
	g.routeset.addLabels(labels)
}

func (g *Generator) makeSegments(value interface{}) {
	s := value.(map[string]interface{})
	split := s["split"].(string)
	days := s["days"].(string)
  fmt.Printf("Making %v Segment with split %s\n",days,split)
	s_int, _ := strconv.Atoi(split)
	segments := SplitRanges(s_int, days)
	g.routeset.addSegments(segments)
}

func (g *Generator) makeAllocations(c interface{}, a interface{}) {
	count := int(c.(float64))
	adrs := int(a.(float64))
	allocations := NewAllocationCollection(count)
	for _, allocation := range allocations {
		allocation.createAdrs(adrs)
	}
	g.routeset.addAllocations(allocations)
}
