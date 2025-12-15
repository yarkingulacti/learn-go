package main // special name means entry point
import "fmt"

func main() {
	fmt.Println("=== Variables ===")
	variables()
	fmt.Println("=== Conditions ===")
	conditions()
	fmt.Println("=== Functions ===")
	functions()
	fmt.Println("=== Arrays ===")
	arrays()
	fmt.Println("=== Loops ===")
	forLoop()
	fmt.Println("=== Maps ===")
	maps()
	fmt.Println("=== Slices ===")
	slices()
	fmt.Println("=== Runes ===")
	runes()
	fmt.Println("=== Structs ===")
	structs()
	fmt.Println("=== Pointers ===")
	pointers()
}
