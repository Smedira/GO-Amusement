package main
 
import (
  "fmt"
  "github.com/zond/god/client"
  "github.com/zond/god/common"
  "github.com/zond/god/murmur"
)
 
// a make believe user data structure
type User struct {
  Email    string
  Password string
  Name     string
}
 
func main() {
  // connect to the default local server
  conn := client.MustConn("localhost:9191")
  // create a user
  
  // try to fetch the user again
  data, _ := conn.Get(murmur.HashString("hi"))
  var found []int
  // to unserialize it
  common.MustJSONDecode(data, &found)
  fmt.Printf("stored and found ", found)
}
