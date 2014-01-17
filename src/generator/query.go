package generator
import(
  "math/rand"
)

type DestinationTable struct {
  DestinationId int64
  AppId int
  Destination string
  DestinationTitle string
  DestinationPropertyName string
  DestinationAttr string
  CallType string
}

type Cache struct {
  Values []int64
}

var DestinationCache Cache
func FillCache() {
  var dests []DestinationTable
  var ids []int64
  Connection.connection.Table("racc_destination").Find(&dests)
  for _,d := range dests {
    ids = append(ids,d.DestinationId)
  }
  DestinationCache.Values = ids
}

func RandomDestination() (random int64) {
  total := len(DestinationCache.Values)-1
  id := rand.Intn(total)
  return DestinationCache.Values[id]
}
