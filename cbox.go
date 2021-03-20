package main 

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/r1cebank/cbox/os"
)

func main() {

    os.GetOperatingSystem()
	fmt.Printf("You choose %q\n", result)
}
