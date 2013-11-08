package generator

import(
  //"fmt"
  "github.com/nu7hatch/gouuid"
  "math/rand"
)

type Allocation struct {
  Id string `json:"id"`
  Percentage int `json:"percentage"`
  Destinations []Destination `json:"destinations"`
}

type Destination struct {
  Destination_id int `json:"destination_id"`
  Exit_type string `json:"type"`
}

func NewAllocationCollection(count int) []*Allocation{
  percentages := []int{30,30,20,20}
  allocations := make([]*Allocation,count)
  for i := 0; i < count; i++ {
    allocations[i] = NewAllocation(percentages[i])
  }
  return allocations
}

func NewAllocation(percentage int) (a *Allocation){
  u,_ := uuid.NewV4()
  destinations := []Destination{}
  destinations = append(destinations,NewDestination())
  a = &Allocation{Id: u.String(), Percentage: percentage, Destinations: destinations}
  return
}

func NewDestination() (d Destination){
  id := rand.Intn(100)
  d = Destination{Destination_id: id, Exit_type: "Destination"}
  return
}

func(a *Allocation) createAdrs(adrs int) {
  destinations := make([]Destination,adrs)
  for i:=0; i<adrs; i++ {
    destinations[i] = NewDestination()
  }
  total := append(a.Destinations,destinations...)
  a.Destinations = total
}
