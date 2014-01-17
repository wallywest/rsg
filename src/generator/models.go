package generator

import (
	"encoding/json"
	"fmt"
  "io/ioutil"
)

type Label struct {
	Id          string `json:"id"`
	Vlabel      string `json:"vlabel"`
	Description string `json:"description"`
}

type Segment struct {
	Id    string `json:"id"`
	Days  string `json:"days"`
	Start int    `json:"start_time"`
	End   int    `json:"end_time"`
}

type GenericRouteSet struct {
  TimeZone    string        `json:"time_zone"`
	Labels      []Label       `json:"labels"`
	Segments    []Segment     `json:"segments"`
	Allocations []*Allocation `json:"allocations"`
}

type SuperProfile struct {
	*GenericRouteSet
  Tree map[string]interface{} `json:"tree"`
}

type Week struct {
	*GenericRouteSet
  Tree map[string][]string  `json:"tree"`
}

type RouteSetInterface interface {
	addLabels()
	addSegments()
	addAllocations()
}

func (s *SuperProfile) addLabels(l []Label) {
	s.Labels = l
}

func (s *SuperProfile) addSegments(segs []Segment) {
  for _,seg := range segs {
    s.Segments = append(s.Segments,seg)
  }
}

func (s *SuperProfile) addAllocations(allocs []*Allocation) {
  for _,alloc := range allocs {
	  s.Allocations = append(s.Allocations,alloc)
  }
}

func (s *SuperProfile) toJSON(name string) {
	b, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println(err)
	}
  ioutil.WriteFile("dump/"+name+".json",b,0644)
}

func (s *SuperProfile) Print() {
	fmt.Println(s.Labels)
	fmt.Println("------------")
	fmt.Println(s.Segments)
	fmt.Println("------------")
	fmt.Println(s.Allocations)
}
