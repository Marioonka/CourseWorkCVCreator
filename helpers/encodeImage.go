package helpers

import (
	"encoding/base64"
	"io"
	"os"
)

func EncodeImageToBase64(filePath string) (string, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	imgBytes, err := io.ReadAll(imgFile)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(imgBytes)
	return encoded, nil
}
