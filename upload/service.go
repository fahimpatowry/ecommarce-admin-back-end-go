package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type UploadServices struct{}

func NewUploadService() *UploadServices {
	return &UploadServices{}
}

func (s *UploadServices) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(header.Filename) // It extracts the file extension from a file name. png, jpg, etc
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	uploadDir := "./public/uploads"
	err := os.MkdirAll(uploadDir, os.ModePerm) // This line is about making sure a directory exists

	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return "/public/uploads/" + fileName, nil
}
