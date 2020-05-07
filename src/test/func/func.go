package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func add(x , y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var x, y = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(add(42, 13))

	a,b :=swap("hello","world")
	fmt.Println(a,b)
	fmt.Println(swap(a,b))
	fmt.Println(split(19))

	var i int
	fmt.Println(i, c, python, java)
	var e, f, g = true, false, "no!"
	fmt.Println(x, y, e, f, g)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var o int
	var p float64
	var q bool
	var r string
	fmt.Printf("%v %v %v %q\n", o, p, q, r)

	m, n := 3, 4
	k := math.Sqrt(float64(m*m + n*n))
	l := int(k)
	fmt.Println(m, n, l)

	v := 123.54+4i// 修改这里！
	fmt.Printf("v is of type %T\n", v)

	fmt.Println((1+3i)*(2+4i))

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
