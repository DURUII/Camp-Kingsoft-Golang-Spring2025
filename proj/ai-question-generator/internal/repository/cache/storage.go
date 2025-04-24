package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

// 使用锁确保文件操作是线程安全的
var fileLock sync.Mutex

// 确保目录存在，如果不存在则创建
func ensureDirExists(filePath string) error {
	// 使用 filepath.Dir 获取目录部分
	dir := filepath.Dir(filePath)

	// 如果目录不存在，则创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	return nil
}

// 重命名旧文件并创建新文件
func renameInvalidFile(fileName string) error {
	// 使用 UUID 重命名文件
	backupFileName := fmt.Sprintf("data/%s_%s.json", time.Now().Format("2006_01_02"), uuid.New().String())
	// Check if the backup file already exists and handle accordingly.
	if _, err := os.Stat(backupFileName); !os.IsNotExist(err) {
		return fmt.Errorf("backup file already exists: %s", backupFileName)
	}
	if err := os.Rename(fileName, backupFileName); err != nil {
		return fmt.Errorf("failed to rename invalid file %s: %v", fileName, err)
	}
	log.Printf("Renamed invalid file to %s", backupFileName)
	return nil
}

// 解析 JSON 文件，如果格式不正确则返回错误
func parseExistingFile(file *os.File) ([]interface{}, error) {
	var existingData []interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&existingData); err != nil {
		return nil, fmt.Errorf("failed to parse existing data: %v", err)
	}
	return existingData, nil
}

func SaveToJSONFile(data interface{}) error {
	// Check if data is empty
	if data == nil {
		log.Println("No data to save.")
		return nil
	}

	// Get the file name based on current date
	fileName := fmt.Sprintf("data/%s.json", time.Now().Format("2006_01_02"))

	// Ensure directory exists
	if err := ensureDirExists(fileName); err != nil {
		return fmt.Errorf("failed to ensure directory exists: %v", err)
	}

	// Lock file operation to ensure thread safety
	fileLock.Lock()
	defer fileLock.Unlock()

	// First, read the file content if it exists
	var existingData []interface{}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Try parsing the existing file data
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&existingData); err != nil && err.Error() != "EOF" {
		// If parsing fails, rename the invalid file and create a new one
		if err := renameInvalidFile(fileName); err != nil {
			return err
		}

		// Recreate a new file after renaming the invalid file
		file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf("failed to create new file: %v", err)
		}
		defer file.Close()

		// Initialize an empty slice if the file was newly created
		existingData = []interface{}{}
	} else {
		// Log the existing data for debugging purposes
		log.Printf("Existing data: %v", existingData)
	}

	// Append the new data to the existing data
	existingData = append(existingData, data)

	// Seek to the beginning of the file to overwrite it
	file.Seek(0, 0)

	// Use json.MarshalIndent to format the JSON with indentation
	prettyJSON, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %v", err)
	}

	// Write the indented JSON to the file
	if _, err := file.Write(prettyJSON); err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	log.Printf("Data successfully saved to %s", fileName)
	return nil
}
