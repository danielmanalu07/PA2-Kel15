package utils

import (
	"fmt"
	"path/filepath"
)

func GenerateImageFile(filename string, typ string) string {
	ext := filepath.Ext(typ)
	filenames := fmt.Sprintf("Image_%s%s", filename, ext)
	return filenames
}
