package verification

import (
	"encoding/hex"
	"io"
	"os"
)

func getFileHexSignature(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 216) // Read 216 bytes for the hex signature
	_, err = io.ReadFull(file, buffer)
	if err != nil {
		return "", err
	}

	hexSignature := hex.EncodeToString(buffer)
	return hexSignature, nil
}
