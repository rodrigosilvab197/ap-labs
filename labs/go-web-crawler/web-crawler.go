package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch5/links"
)

type node struct {
	url   string
	depth int
}

var tokens = make(chan struct{}, 20)

func crawl(link node) []node {
	tokens <- struct{}{}
	list, err := links.Extract(link.url)
	<-tokens

	if err != nil {
	}
	nextUrl := make([]node, 0, len(list))
	for _, val := range list {
		nextUrl = append(nextUrl, node{val, link.depth + 1})
	}
	return nextUrl
}
func displayProgress() {
	states := []string{"|", "/", "-", "\\"}
	for i := 0; ; i = (i + 1) % len(states) {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\r\033[K\r%s", states[i])
	}
}
func main() {
	if len(os.Args) < 4 {
		log.Println("Porfavor Pasa los parametro bien")
		return
	}
	depth := flag.Int("depth", 0, "Web crawler depth")
	results := flag.String("results", "results.txt", "Output file")
	flag.Parse()
	fmt.Printf("Web Crawler\nLink: %s\nDepth: %d\n", os.Args[3], *depth)
	f, err := os.Create(*results)
	if err != nil {
		log.Println("No se pudo crear el archivo")
		return
	}
	w := bufio.NewWriter(f)
	defer w.Flush()
	worklist := make(chan []node)
	var n int
	n++
	go func() {
		worklist <- []node{node{os.Args[3], 0}}
	}()
	currentDepth := -1
	visited := make(map[string]bool)
	go displayProgress()
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !visited[link.url] {
				if link.depth > currentDepth {
					currentDepth = link.depth
				}
				visited[link.url] = true
				fmt.Fprintf(w, "%s\n", link.url)
				if link.depth+1 > *depth {
					continue
				}
				n++
				go func(link node) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
	fmt.Printf("\nFile Resulsts was created\n")
}
