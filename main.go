package main
import "fmt"
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

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
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
