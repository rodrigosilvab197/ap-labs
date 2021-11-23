// Netcat2 is a read-only TCP client with channels
package main
import (
	"os"
  	"fmt"
	"strings"
	"net"
	"log"
		
  
  
)
func connect(conn net.Conn,) {
  for {
    lTime := make([]byte,11)
    _, err := conn.Read(lTime)
    if err != nil{
      conn.Close()
      log.Print(err)
    } else {
      fmt.Printf("\t%s", lTime)
    }
  }
}
func main() {
  args := os.Args[1:]
  if(len (args) < 1){
    fmt.Println("Insuficient Parameters !")
    os.Exit(1)
  }
	for i := 0; i < len(args); i++ {
    Tzone := strings.Split(os.Args[i+1], "=")
    if len(Tzone) != 2 {
		fmt.Println("Insuficient Parameters !")
    }
  	conn, err := net.Dial("tcp", Tzone[1])
    if err != nil {
      log.Fatal(err)
    }
    go connect(conn)
  }
  for {

  }
}