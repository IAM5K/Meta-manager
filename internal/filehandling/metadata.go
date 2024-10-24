package filehandling

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type FileMetadata struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Extension string    `json:"extension"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
}

type Folder struct {
	Name     string         `json:"name"`
	Metadata []FileMetadata `json:"metadata"`
	Subdirs  []Folder       `json:"subdirs"`
}

func getFileMetadata(filePath string) (FileMetadata, error) {
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return FileMetadata{}, err
	}

	return FileMetadata{
		Name:      fileStat.Name(),
		Size:      fileStat.Size(),
		Extension: filepath.Ext(fileStat.Name()),
		Created:   fileStat.ModTime(),
		Modified:  fileStat.ModTime(),
	}, nil
}

func createFolderStructure(baseDir string) (Folder, error) {
	folderName := filepath.Base(baseDir)
	folder := Folder{
		Name:     folderName,
		Metadata: []FileMetadata{},
		Subdirs:  []Folder{},
	}

	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return Folder{}, err
	}

	for _, file := range files {
		filePath := filepath.Join(baseDir, file.Name())
		if file.IsDir() {
			subfolder, err := createFolderStructure(filePath)
			if err != nil {
				return Folder{}, err
			}
			folder.Subdirs = append(folder.Subdirs, subfolder)
		} else {
			metadata, err := getFileMetadata(filePath)
			if err != nil {
				return Folder{}, err
			}
			folder.Metadata = append(folder.Metadata, metadata)
		}
	}

	return folder, nil
}

func ExtractMetadataToJson(baseDir string) error {
	folderStructure, err := createFolderStructure(baseDir)
	if err != nil {
		return err
	}

	parentFolderName := filepath.Base(baseDir)
	timestamp := time.Now().Format("20060102150405") // Format: YYYYMMDDHHmmss
	jsonFilePath := filepath.Join(baseDir, parentFolderName+timestamp+".json")
	jsonData, err := json.MarshalIndent(folderStructure, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
