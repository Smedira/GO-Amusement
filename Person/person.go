package Person

import (
  "GO-Amusement/Path"
)

type person struct {
  loc Path.location
  path []Path.path
  dir []bool
  speed int
  appointments []byte
  times []int //len(times) = len(appointments)
  

}

func (p *person) move() {

  p.loc.length += p.dir[0] * p.speed
  if p.loc.spare() != 0 {
    p.dir = p.dir[1:]
    p.path = p.path[1:]
    if p.dir[0] > 0 {
      p.loc = p.path[0].l1
    } else {
      p.loc = p.path[0].l2
      p.loc.length = p.path[0].length
    }

  }

}


