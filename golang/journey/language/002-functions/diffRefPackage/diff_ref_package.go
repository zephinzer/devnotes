package diffRefPackage

import "fmt"

// HelloFromADifferentPackage says hello (linter says we need a comment here)
func HelloFromARenamedPackage() {
	fmt.Printf("Hello from a renamed package!\n")
}
