package main
import "fmt"
func main() {
	first := 1
	second := 2
	next := 3
	sum := 0
	for second <= 4000000{
		if second % 2 == 0{
			sum = sum + second
		}
		next = first + second // next = 1 + 2 = 3
		first = second // first = 2
		second = next // second = 3
	}
	fmt.Println(sum)
}