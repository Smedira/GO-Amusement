package Path/**

func min(a int, b int) int {
  if (a < b) {
    return a
  }
  return b
}

func (l1 *location) equals (l2 *location) bool {
  //just compare keys

func in



func genPath (l1 *location, l2 *location, l []location, d int) []location, int {// returns list of locations and int for distance
  dis := 90000000
  if l1.equals(l2) {return l, d}
  if in(l1, l){ return nil, 9000000 }
  for (i := 0; i < len(l1.path); i++){
    if !l1.path[i].l1.equals(l1) {
      ll, dd := genpath(LGet(l1.path[i].l1),l2,append(l,l1),d + l1.path[i].length)
    } else {
      ll, dd := genpath(LGet(l1.path[i].l2),l2,append(l,l1),d + l1.path[i].length)
    }
    if (dd < dis) {
      dis = dd
      locs := ll
    }
  }
  return l, d + dis
}
  **/
