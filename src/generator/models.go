package generator

import (
	"encoding/json"
	"fmt"
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
	Labels      []Label       `json:"labels"`
	Segments    []Segment     `json:"segments"`
	Allocations []*Allocation `json:"allocations"`
}

type SuperProfile struct {
	*GenericRouteSet
	Tree map[string]interface{}
}

type Week struct {
	*GenericRouteSet
	Tree map[string][]string
}

type RouteSetInterface interface {
	addLabels()
	addSegments()
	addAllocations()
}

func (s *SuperProfile) addLabels(l []Label) {
	s.Labels = l
}

func (s *SuperProfile) addSegments(seg []Segment) {
	s.Segments = seg
}

func (s *SuperProfile) addAllocations(a []*Allocation) {
	s.Allocations = a
}

func (s *SuperProfile) toJSON() {
	b, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func (s *SuperProfile) Print() {
	fmt.Println(s.Labels)
	fmt.Println("------------")
	fmt.Println(s.Segments)
	fmt.Println("------------")
	fmt.Println(s.Allocations)
}
