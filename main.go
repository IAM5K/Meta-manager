package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iam5k/file-auditor/internal/filehandling"
	"github.com/iam5k/file-auditor/internal/verification"
)

func main() {
	fmt.Println("Welcome to File Handler!")

	// Ask user for directory path
	fmt.Print("Enter the directory path: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dirPath := scanner.Text()

	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Println("Directory does not exist.")
		os.Exit(1)
	}

	fmt.Println("Choose an option:")
	fmt.Println("1. Extract Metadata to JSON")
	fmt.Println("2. Verify for Corrupt Files")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	switch choice {
	case 1:
		err = filehandling.ExtractMetadataToJson(dirPath)
		if err != nil {
			fmt.Println("Error extracting metadata:", err)
			os.Exit(1)
		}
		fmt.Println("Metadata extraction completed and saved as JSON.")

	case 2:
		fmt.Println("Verifying for corrupt files...")
		corruptFiles, err := verification.VerifyFilesForCorruption(dirPath)
		if err != nil {
			fmt.Println("Error during verification:", err)
			os.Exit(1)
		}

		if len(corruptFiles) == 0 {
			fmt.Println("No corrupt files found.")
		} else {
			fmt.Println("Corrupt files found:")
			for _, filename := range corruptFiles {
				fmt.Println(filename)
			}
		}
	default:
		fmt.Println("Invalid choice.")
	}

	fmt.Print("Press Enter to exit...")
	fmt.Scanln() // Wait for Enter key press
	// Trigger the metadata extraction function
	// err = filehandling.ExtractMetadataToJson(dirPath)
	// if err != nil {
	// 	fmt.Println("Error extracting metadata:", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Metadata extraction completed and saved as JSON.")

}
