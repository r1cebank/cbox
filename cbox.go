package main 

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/r1cebank/cbox/os"
)

func main() {
	// Get the current user os, if supported, continue
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	os.OSSupported()

	fmt.Printf("You choose %q\n", result)
}
