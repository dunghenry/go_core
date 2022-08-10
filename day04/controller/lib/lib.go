package lib
import (
	"math/rand"
)
func Random(max int, min int) int{
	var random = rand.Intn(max - min) + min
	return random
}
func Pow(a int, b int) int{
	var c int = 1
	var i int = 0
	for i < b{
		c *= a
		i += 1
	}
	return c
}