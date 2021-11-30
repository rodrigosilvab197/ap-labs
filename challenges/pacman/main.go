package main

import (
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"flag"
)
type posicionPlayer struct {
	row      int
	col      int
}
type posicionEnemie struct {
	row      int
	col      int
	init_row int
	init_col int
}
type enemies struct {
	position posicionEnemie
	status   EnemieStatus
}
type EnemieStatus string
const (
	EnemieStatusKilling EnemieStatus = "killing"
	EnemieStatusRunning   EnemieStatus = "Running"
)
var laberinto []string
var player posicionPlayer
var enemmiesVar []*enemies
var points int
var dots int
var pellet sync.Mutex
var EnemieStatusSync sync.RWMutex
var pelletTime *time.Timer
var (
	laberintoFile   = flag.String("laberinto", "modified.txt", "Laberinto Estatico")
)
func eatPoints(row, col int){
	laberinto[row] = laberinto[row][0:col] + " " + laberinto[row][col+1:]
 }
func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}
func CargarLaberinto(file string) error {
	x, err := os.Open(file)
	if err != nil {
		return err
	}
	defer x.Close()
	scn := bufio.NewScanner(x)
	for scn.Scan() {
		line := scn.Text()
		laberinto = append(laberinto, line)
	}
	for row, line := range laberinto {
		for col, char := range line {
			switch char {
			case '.':
				dots++
			case 'P':
				player = posicionPlayer{row, col}
			case 'G':
				enemmiesVar = append(enemmiesVar, &enemies{posicionEnemie{row, col, row, col},EnemieStatusKilling})
			}
		}
	}
	return nil
}
func imprimirPantalla() {
	fmt.Print("\x1b[2J")
	moveCursor(0, 0)
	for _, line := range laberinto {
		for _, chr := range line {
			switch chr {
			case 'X':
				fmt.Printf("X")
			case '#':
				fallthrough
			case '.':
				fmt.Printf("%c", chr)
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	moveCursor(player.row, player.col)
	fmt.Print("P")

	EnemieStatusSync.RLock()
	for _, g := range enemmiesVar {
		moveCursor(g.position.row, g.position.col)
		if g.status == EnemieStatusKilling {
			fmt.Printf("G")
		} else if g.status == EnemieStatusRunning {
			fmt.Printf("C")
		}
	}
	EnemieStatusSync.RUnlock()

	moveCursor(len(laberinto)+1, 0)

	fmt.Println("Puntuacion:", points)
}
func movimientoEnemigo() {
	for _, e := range enemmiesVar {
		dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
		e.position.row, e.position.col = makeMove(e.position.row, e.position.col, move[dir])
	}
}
func controller() (string, error) {
	buff := make([]byte, 100)
	keyboard, err := os.Stdin.Read(buff)
	if err != nil {
		return "", err
	}
	if keyboard == 1 && buff[0] == 0x20 {
		return "SPACEBAR", nil
	} else if keyboard >= 3 {
		if buff[0] == 0x1b && buff[1] == '[' {
			switch buff[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}
	return "", nil
}
func makeMove(x, y int, move string) (row, col int) {
	row, col = x, y
	switch move {
	case "RIGHT":
		col = col + 1
		if col == len(laberinto[0]) {
			col = 0
		}
	case "LEFT":
		col = col - 1
		if col < 0 {
			col = len(laberinto[0]) - 1
		}
	case "UP":
		row = row - 1
		if row < 0 {
			row = len(laberinto) - 1
		}
	case "DOWN":
		row = row + 1
		if row == len(laberinto)-1 {
			row = 0
		}
	}
	if laberinto[row][col] == '#' {
		row = x
		col = y
	}
	return
}
func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)
	switch laberinto[player.row][player.col] {
	case '.':
		points+=10
		eatPoints(player.row, player.col)
		dots--
	case 'X':
		points+=50
		eatPoints(player.row, player.col)
		go pelletPower()
	}
}
func sEnemie(enemie []*enemies, enemieStaus EnemieStatus) {
	EnemieStatusSync.Lock()
	defer EnemieStatusSync.Unlock()
	for _, e := range enemie {
		e.status = enemieStaus
	}
}
func pelletPower() {
	pellet.Lock()
	sEnemie(enemmiesVar, EnemieStatusRunning)
	if pelletTime != nil {
		pelletTime.Stop()
	}
	pelletTime = time.NewTimer(time.Second * 10)
	pellet.Unlock()
	<-pelletTime.C
	pellet.Lock()
	pelletTime.Stop()
	sEnemie(enemmiesVar, EnemieStatusKilling)
	pellet.Unlock()
}
func incializarcMode() {
	cmd := exec.Command("stty", "-cbreak", "echo")
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatalln("unable to activate cooked mode:", err)
	}
}
func main() {
	numberOfEnemies := os.Args[1]
	enemiesCount, err1 := strconv.Atoi(numberOfEnemies)
	if err1 != nil || enemiesCount >12||enemiesCount <1 {
		log.Fatalln("numero de enemigos mayor a 0 y menor a 13", err1)
	}
	var e int = 12-enemiesCount
	if e>=0{
		input, err := ioutil.ReadFile("laberinto.txt")
		if err != nil {
			log.Fatalln("no se encontro laberinto", err)
		}

		output := bytes.Replace(input, []byte("G"), []byte(" "), e)

		if err = ioutil.WriteFile("modified.txt", output, 0666); err != nil {
			log.Fatalln("no se pudo crear laberinto", err)
		}
	}
	flag.Parse()
		cdm := exec.Command("stty", "cbreak", "-echo")
		cdm.Stdin = os.Stdin
	
		err2 := cdm.Run()
		if err2 != nil {
			log.Fatalln("unable to activate cbreak mode:", err2)
		}
	defer incializarcMode()
	err3 := CargarLaberinto(*laberintoFile)
	if err3 != nil {
		log.Fatalln("no se pudo crear laberinto", err3)
	}
	hilo := make(chan string)
	go func(ch chan<- string) {
		for {
			hilo, err := controller()
			if err != nil {
				log.Print("error reading input:", err)
				ch <- "SPACEBAR"
			}
			ch <- hilo
		}
	}(hilo)
	for {
		select {
		case jugador := <-hilo:
			if jugador == "SPACEBAR" {
			fmt.Println("Juego Terminado")
			fmt.Println(points)
			os.Exit(1)
			break
			}
			movePlayer(jugador)
		default:
		}
		movimientoEnemigo()
		imprimirPantalla()
		for _, g := range enemmiesVar {
			if player.row == g.position.row && player.col == g.position.col {
				EnemieStatusSync.RLock()
				if g.status == EnemieStatusKilling {
					fmt.Println("Perdiste")
					fmt.Println(points)
					os.Exit(1)
					break
				} else if g.status == EnemieStatusRunning {
					points+=200
					EnemieStatusSync.RUnlock()
					sEnemie([]*enemies{g}, EnemieStatusKilling)
					g.position.row, g.position.col = g.position.init_row, g.position.init_col
				}
			}
		}
		if dots == 0  {
			fmt.Println("Ganaste")
			fmt.Println(points)
			os.Exit(1)
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
}