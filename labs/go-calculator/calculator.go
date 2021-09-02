package main

// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)
import (
	"fmt"
	"os"
	"strconv"
)
func calc(operator int, values []int) int {
	value := values[0]
	if operator == 1{
		for i:=1; i<len(values);i++{
			value += values[i]
		}
		
		return value
	} else if operator == 2{
		for i:=1; i<len(values);i++{
			value -= values[i]
		}
		return value
	} else {
		for i:=1; i<len(values);i++{
			value *= values[i]
		}
		return value
	}
}

func main() {
	if len(os.Args) > 3 {
		slice := make([]int, 0, 100)
		for i := 2 ; i< len(os.Args); i++{
			j,err:= strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Println("El numero no es valido")
				return
			}
			slice = append(slice, j)	
		} 
		if os.Args[1]== "add"{
			calc(1,slice)
		}else if os.Args[1]== "sub"{
			calc(2,slice)
		} else if os.Args[1]== "mult"{
			calc(3,slice)
		}else{
			fmt.Println("argumento no valido de operacion")
		}
       
		
    } else {
        fmt.Println("not enough arguments we need more to make a calculations ! please try Again")
    }
	
}
