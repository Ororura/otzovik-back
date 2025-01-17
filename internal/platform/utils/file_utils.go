package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveFile(file *multipart.FileHeader, uploadDir string) (string, error) {
    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        return "", err
    }

    ext := filepath.Ext(file.Filename)
    fileName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), strings.TrimSuffix(file.Filename, ext), ext)
    filePath := filepath.Join(uploadDir, fileName)

    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    dst, err := os.Create(filePath)
    if err != nil {
        return "", err
    }
    defer dst.Close()

    if _, err = io.Copy(dst, src); err != nil {
        return "", err
    }

    return fileName, nil
} 