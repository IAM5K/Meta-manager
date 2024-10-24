package verification

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isCorrupted(filePath string) (bool, error) {
	hexSignature, err := getFileHexSignature(filePath)
	if err != nil {
		return false, err
	}

	// Get the file extension from the file name
	fileName := filepath.Base(filePath)
	lastDotIndex := strings.LastIndex(fileName, ".")
	if lastDotIndex == -1 {
		return false, nil // No extension found
	}
	fileExtension := fileName[lastDotIndex+1:]
	fileExtension = strings.ToLower(fileExtension)

	fmt.Printf("Checking File %s ", filePath)
	// Remove null characters from hex signature
	hexSignature = strings.ReplaceAll(hexSignature, "00", "")

	// Compare hex signature with extension
	if hexSignature != "" && hexSignature != fileExtension {
		// fmt.Printf("Warning: Extension mismatch for file %s (Expected: %s, Hex: %s)\n", filePath, fileExtension, hexSignature)
		fmt.Printf("Warning: Extension mismatch for file %s (Expected: %s)\n", filePath, fileExtension)

		return false, nil
	}

	// Check if the hex signature is all zeros (corrupted)
	if hexSignature == "" {
		fmt.Printf("Corrupted file: %s\n", filePath)
		return true, nil
	}

	return false, nil
}

// ...

func VerifyFilesForCorruption(baseDir string) ([]string, error) {
	var corruptFiles []string

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			corrupted, err := isCorrupted(path)
			if err != nil {
				return err
			}

			if corrupted {
				corruptFiles = append(corruptFiles, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return corruptFiles, nil
}
