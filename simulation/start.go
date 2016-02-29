package simulation

import (
  "bufio"
  "fmt"
//  "log"
  "os"
  "strings"
  "strconv"
//  "github.com/zond/god/client"
//  "github.com/zond/god/common"
  "github.com/zond/god/murmur"
  "GO-Amusement/path"
)

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func in(a string, b []string) int{
  for i := 0; i < len(b); i++ {
    if a == b[i] { return i }
  }
  return -1
}

func setupLoc(){

  fmt.Println("KJG")

//* dereferences pointer

//make(map[string]int)

  //ml := make(map[int][]*Path.Location)
  //mp := make(map[int]Path.Path)
  var l []string
  var locs []*Path.Location
  var l1 *Path.Location
  var l2 *Path.Location

  count := 0

  r,_ := readLines("/home/carmine/Documents/go/src/GO-Amusement/path/out.txt")
  for i := 1; i < len(r); i++ {
    s := strings.Split(r[i],",")
    if s[2] != "0" {
      n := in(s[0], l)
      if n == -1 {
        l = append(l,s[0])
        l1 = &Path.Location{K: murmur.HashString("location" + s[0])}
        locs = append(locs,l1)
        n = len(locs)
      } else {
        l1 = locs[n]
      }
      

      n = in(s[1], l)
      if n == -1 {
        l = append(l,s[1])
        l2 = &Path.Location{K: murmur.HashString("location" + s[1])}
        locs = append(locs,l2)
        n = len(locs)
      } else {
        l2 = locs[n]
      }

      k := murmur.HashString("path" + s[0] + s[1])

      s,_ := strconv.Atoi(s[2])

      p := Path.Path{K: k,Length: s, L1:l1.K, L2:l2.K}
      p.Save()

      l1.Paths = append(l1.Paths,p)
      l2.Paths = append(l2.Paths,p)


      fmt.Println(count)

      count++
    
    }

  }

  //fmt.Println(len(mp))

  
  for i:= 0; i < len(locs); i++ {
    locs[i].Save()
  }

  fmt.Println(l)
}
      
    
    

  

func setup() {}
  
