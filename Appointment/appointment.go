package appointment

import (
  "GO-Amusement/ride"
  "GO-Amusement/Path"
)

type appointment struct {

  ride []key
  loc Path.location
  time []int
  rowWidth int
  numRows int
  positions []key
  present []bool //same len as positions
  timeDistribution key
  numUntil int
  exit []
  entrance []
  open bool

}
  
  
