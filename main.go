package main
import (
	"fmt"
	"time"
	"sync"
)
func main() {
    fmt.Println("Hello, Go!")

	var name string = "John Doe"
	var age int = 25
	
	fmt.Printf("Name: %s\nAge: %d\n", name,age)

	batsman := "David Warner"
	runs := 34
	fmt.Printf("Batsmen: %s\nRuns: %d\n", batsman, runs)

	if runs >= 50 {
		fmt.Println("Half-century!")
	} else{
		fmt.Println("Not a half-century!")
	}

	var wg sync.WaitGroup
	for i := 0; i < 50 ; i++ {
		wg.Add(1)
			go printNumbers(i, &wg)
	}
	wg.Wait()
	
	fmt.Println(add(5,6))

	p := Person{"John Doe", 25}
	fmt.Printf("Name: %s\nAge: %d\n", p.Name, p.Age)
	
}

func add(a int, b int) int{
	return a + b
}

type Person struct {
	Name string
	Age int
}

func printNumbers (n int, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println(n)
}
