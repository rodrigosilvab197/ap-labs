// Clock Server is a concurrent TCP server that periodically writes the time.
package main
import (
	"io"
	"log"
	"net"
	"time"
	"os"
	"fmt"
	
)
func handleConn(c net.Conn, tZone string) {
	defer c.Close()
	for {
		_, timeErr := time.LoadLocation(tZone)
		if timeErr != nil{
			log.Print(timeErr)
			break
		}
		_, err := io.WriteString(c, tZone+" "+time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Insuficient Parameters !")
		os.Exit(1)
	}
	lHost := "localhost:" + args[1]
	listener, err := net.Listen("tcp", lHost)
	if err != nil {
		log.Fatal(err)
	}
	tZone := os.Getenv("TZ")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, tZone) // handle connections concurrently
	}
}