package main

import (
	"fmt"
)

/* Here is a list of basic types: 

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

PS, a complex number can work like this: 
z complex128 = cmplx.Sqrt(-5 + 12i)
*/

//Using var to declare 3 boolean variables, note that the type is at the end
var a, c, d bool

var l, j int = 1, 2 

func main(){
	var i bool
	fmt.Println("Varables that were created: ", a, c, d, i)
	fmt.Println("Variables with type omitted: ", l, j)
}