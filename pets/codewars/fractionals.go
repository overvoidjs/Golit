package main

import (
"fmt"
"math/big"
)



func main() {
var x float64
x = 1.0/2.0 + 1.0/3.0 + 1.0/4.0
fmt.Println(x)
r := new(big.Rat)
z2 := r.SetFloat64(1.0833)
fmt.Println(z2)

}
