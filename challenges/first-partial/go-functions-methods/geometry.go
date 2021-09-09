// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
	"math/rand"
)
type Point struct{ x, y float64 }
// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}
// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}
//!-point
//!+path
// A Path is a journey connecting the points with straight lines.
type Path []Point
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			r := path[i-1].Distance(path[i])
			sum += r
			fmt.Printf("%.2f + ", math.Floor(r*100)/100)
		}
	}
	return sum
}
// Given three colinear points p, q, r, the function checks if
// point q lies on line segment 'pr'
func onSegment(p, q, r Point) bool {
	if p.x <= math.Max(p.x, r.x) && q.x >= math.Min(p.x, r.x) && q.y <= math.Max(p.y, r.y) && q.y >= math.Min(p.y, r.y) {
		return true
	}
	return false
}
// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are colinear
// 1 --> Clockwise
// 2 --> Counterclockwise
func orientation(p, q, r Point) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)
	if val == 0 {
		return 0
	} else if val > 0 {
		return 1
	} else {
		return 2
	}
}
// Function that returns true if line segment 'p1q1'
// and 'p2q2' intersect.
func doIntersect(p1, q1, p2, q2 Point) bool {
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)
	if o1 != o2 && o3 != o4 {
		return true
	}
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}
	return false
}
func puntoRandom()Point{
	var min, max float64
		min = -100
		max = 100
	return Point{min + rand.Float64()*(max-min), min + rand.Float64()*(max-min)}
}
func imprimirResultado(slices int, puntos Path){
		fmt.Printf(" - Generating a [%d] sides figure\n", slices)
		fmt.Println(" - Figure's vertices")
		for p := range puntos {
				fmt.Printf("   - (%3.1f,%3.1f)\n", puntos[p].x, puntos[p].y)
			}
			fmt.Println(" - Figure's Perimeter")
			fmt.Print("   - ")
			perimetro := puntos[len(puntos)-1].Distance(puntos[0])
			perimetroFinal := puntos.Distance() + perimetro
			fmt.Printf("%.2f = ", math.Floor(perimetro*100)/100)
			fmt.Printf("%.2f\n", math.Floor(perimetroFinal*100)/100)
}
func main() {
	argc := len(os.Args)
	if argc < 2 {
		fmt.Println("Ups! Te falto poner el numero de lienas para la figura! por favor pasa un argumento(numero Entero)")
	} else {
		var puntos Path
		slices, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("EL numero debe de ser entero! y mayor a 2 por favor pasa argumento de nuevo")
		} else {
			if slices <3{
				fmt.Println("EL numero debe de ser mayor a 2 para realizar la figura ")	
			}else{
				puntos = make([]Point, slices)
				puntos[0] = puntoRandom()
				puntos[1] = puntoRandom()
				puntos[2]= puntoRandom()
			for onSegment(puntos[0], puntos[2], puntos[1]) {
				puntos[2] = puntoRandom()
			}
			for i := 3; i < slices; i++ {
				puntos[i] = puntoRandom()
					j := 0
					r := false
					for j < i-2 {
					inter := doIntersect(puntos[i], puntos[i-1], puntos[j], puntos[j+1])
					if inter {
					j = i
					r = inter
					} else {
					j++
					}
					}
					for r {
						puntos[i] = puntoRandom()
					j = 0
					r = false
					for j < i-2 {
					s := doIntersect(puntos[i], puntos[i-1], puntos[j], puntos[j+1])
					if s {
					r = s
					j = i
					} else {
					j++
					}
					}
					}
				
			}
			imprimirResultado(slices,puntos)
			
			}
		}
	}	
}
//!-path