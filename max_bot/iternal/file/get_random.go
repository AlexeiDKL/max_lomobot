package file

import (
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

// GetRandomCatImage возвращает путь к случайному изображению из папки "./cats".
// Путь относительный от текущей рабочей директории.
func GetRandomCatImage() (string, error) {
	catsDir := "./cats"

	entries, err := os.ReadDir(catsDir)
	if err != nil {
		return "", err
	}

	var imageFiles []string
	extensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if extensions[ext] {
			imageFiles = append(imageFiles, filepath.Join(catsDir, entry.Name()))
		}
	}

	if len(imageFiles) == 0 {
		return "", errors.New("no image files found in cats directory")
	}

	randomIndex := rand.Intn(len(imageFiles))
	return imageFiles[randomIndex], nil
}
