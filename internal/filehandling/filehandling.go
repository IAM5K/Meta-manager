package filehandling

import (
	"os"
	"path/filepath"
)

// type FileMetadata struct {
// 	Name      string
// 	Size      int64
// 	Extension string
// 	// Add more fields as needed
// }

func ScanDirectory(dirPath string) ([]FileMetadata, error) {
	var filesMetadata []FileMetadata

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			extension := filepath.Ext(info.Name())
			filesMetadata = append(filesMetadata, FileMetadata{
				Name:      info.Name(),
				Size:      info.Size(),
				Extension: extension,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return filesMetadata, nil
}
