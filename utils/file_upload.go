package utils

import (
	"bitroom/constants"
	"bitroom/types"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func HanldeFileUpload(file *multipart.FileHeader) (string, *types.CustomError) {
	// get file src
	src, err := file.Open()
	if err != nil {
		return "", NewError("failed to get file", http.StatusBadRequest)
	}
	defer src.Close()

	// change name
	newName, customErr := FileName(file.Filename)
	if customErr != nil {
		return "", customErr
	}
	file.Filename = newName

	// create a destination file
	uploadPath := filepath.Join(constants.StaticFolderName, file.Filename)
	dst, err := os.Create(uploadPath)
	if err != nil {
		return "", NewError("Failed to create file", http.StatusInternalServerError)
	}
	defer dst.Close()

	// move to destination
	if _, err := io.Copy(dst, src); err != nil {
		return "", NewError("Failed to copy file", http.StatusInternalServerError)
	}

	// success
	return uploadPath, nil
}

func FileName(file_name string) (string, *types.CustomError) {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())
	milliSecond := strconv.FormatInt(now.UnixMilli(), 10)

	// get current time
	name := year + month + day + hour + minute + second + milliSecond
	// get ext
	ext := filepath.Ext(file_name)
	if ext == "" {
		return "", NewError("please upload image", http.StatusInternalServerError)
	}
	// full name
	full_name := name + ext
	// success
	return full_name, nil
}
