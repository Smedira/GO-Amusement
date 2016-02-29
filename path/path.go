package Path

import (

  "github.com/zond/god/client"
  "github.com/zond/god/common"

)

//TODO add save and get

//for get distance series of ints stored into cloud between all the points for rough approximation

type Path struct {
	Length int //meters
	L1 []byte //Location
	L2 []byte //Location
  K []byte

}

func LGet([]byte) {}  //TODO



func (p *Path) Save() {

  conn := client.MustConn("10.0.0.28:9191")

  bytes := common.MustJSONEncode(*p)
  conn.Put(p.K, bytes)

}


type Location struct { 
//if not intersection, len(path) = 1, length = dis to l2
//  for movement, simple increment or decrement length
//if intersection, length = 0
//if in space, only s defined and direct dis calls for movement
// when reach space, set location to whatever entrance used
// when reach attract in space, set loc to that value
	Paths []Path //keys for paths
  Length int
  S []byte //space
  K []byte //need equal function
}

func (l *Location) Save() {

  conn := client.MustConn("10.0.0.28:9191")

  bytes := common.MustJSONEncode(*l)
  conn.Put(l.K, bytes)

}

func (l *Location) Spare() int {
  if l.Length < 0 { return l.Length * -1
  } else if l.Length > l.Paths[0].Length { return l.Length - l.Paths[0].Length 
  } else { return 0 }
}

type Space struct {
//Each space will have a location to broadly represent those meandering in the area
  Central Location
	I []Location
}
