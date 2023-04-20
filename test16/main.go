package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func main() {
	// specify the directory containing the folders
	dir := "/path/to/directory"

	// get a list of directories in the specified directory
	folders, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// loop through each directory and run the Git commands
	for _, folder := range folders {
		// run git fetch
		cmdFetch := exec.Command("git", "fetch")
		cmdFetch.Dir = folder
		outputFetch, err := cmdFetch.Output()
		if err != nil {
			fmt.Println("Error running git fetch:", err)
			continue
		}

		// run git pull
		cmdPull := exec.Command("git", "pull")
		cmdPull.Dir = folder
		outputPull, err := cmdPull.Output()
		if err != nil {
			fmt.Println("Error running git log:", err)
			continue
		}

		// print the output
		fmt.Printf("Git fetch for %s:\n%s\n", folder, string(outputFetch))
		fmt.Printf("Git pull for %s:\n%s\n", folder, string(outputPull))
		fmt.Println("--------------------------------------------------")
	}
}
