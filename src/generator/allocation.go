package generator

import (
	//"fmt"
	"github.com/nu7hatch/gouuid"
	"math/rand"
)

type Allocation struct {
	Id           string        `json:"id"`
	Percentage   int           `json:"percentage"`
	Destinations []Destination `json:"destinations"`
}

type Destination struct {
	Destination_id int64    `json:"destination_id"`
	Exit_type      string `json:"type"`
}

func randInt(min int, max int) int {
  return min + rand.Intn(max-min)
}

func MakePercentages(count int) ([]int) {
  total := 100
  min := 0
  i := make([]int,count)
  for index,_ := range i {
    if index == count -1 {
      i[index] = total
    } else {
      val := randInt(min,total)
      i[index] = val
      total = total - val
    }
  }
  return i
}

func NewAllocationCollection(count int) []*Allocation {
  percentages := MakePercentages(count)
	allocations := make([]*Allocation, count)
	for i := 0; i < count; i++ {
		allocations[i] = NewAllocation(percentages[i])
	}
	return allocations
}

func NewAllocation(percentage int) (a *Allocation) {
	u, _ := uuid.NewV4()
	destinations := []Destination{}
	destinations = append(destinations, NewDestination())
	a = &Allocation{Id: u.String(), Percentage: percentage, Destinations: destinations}
	return
}

func NewDestination() (d Destination) {
	d = Destination{Destination_id: RandomDestination(), Exit_type: "Destination"}
	return
}

func (a *Allocation) createAdrs(adrs int) {
	destinations := make([]Destination, adrs)
	for i := 0; i < adrs; i++ {
		destinations[i] = NewDestination()
	}
	total := append(a.Destinations, destinations...)
	a.Destinations = total
}
