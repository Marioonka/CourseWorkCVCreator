package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const MaxSearchDepth = 5 // Максимальная глубина поиска по директориям

// GetPathToFile ищет файл по заданному имени начиная с нескольких базовых путей и рекурсивно проходит папки.
func GetPathToFile(filename string) (string, error) {
	// Список базовых путей для поиска
	basePaths := []string{}

	// Текущая рабочая директория
	if cwd, err := os.Getwd(); err == nil {
		basePaths = append(basePaths, cwd)
	}

	// Директория исполняемого файла
	if execPath, err := os.Executable(); err == nil {
		basePaths = append(basePaths, filepath.Dir(execPath))
	}

	// Директория исходного кода (для среды разработки)
	_, callerName, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(callerName), "..", "..")
	basePaths = append(basePaths, projectRoot)

	// Рекурсивный поиск
	for _, basePath := range basePaths {
		if path, err := recursiveSearch(basePath, filename, MaxSearchDepth); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("файл %s не найден в доступных путях", filename)
}

// recursiveSearch выполняет рекурсивный поиск файла в указанной директории с заданной глубиной.
func recursiveSearch(basePath, filename string, depth int) (string, error) {
	if depth < 0 {
		return "", fmt.Errorf("достигнута максимальная глубина поиска")
	}

	// Проверяем, существует ли файл в текущей директории
	targetPath := filepath.Join(basePath, filename)
	if _, err := os.Stat(targetPath); err == nil {
		return targetPath, nil
	}

	// Получаем список всех поддиректорий и ищем в них
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(basePath, entry.Name())
			if path, err := recursiveSearch(subDir, filename, depth-1); err == nil {
				return path, nil
			}
		}
	}

	return "", fmt.Errorf("файл %s не найден в директории %s", filename, basePath)
}
