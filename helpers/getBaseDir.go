package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetBaseDir() {
	// Получаем директорию текущего исполняемого файла
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	execDir := filepath.Dir(execPath)

	// Строим путь к файлу
	filePath := filepath.Join(execDir, "data", "config.json")
	fmt.Println(filePath)
}
