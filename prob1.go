package main
import "fmt"
func main(){
	sum := 0
	for counter := 3; counter < 1000; counter++ {
		if counter % 3 == 0 || counter % 5 == 0{
			sum = sum + counter
		}
	}
	fmt.Println(sum)
}