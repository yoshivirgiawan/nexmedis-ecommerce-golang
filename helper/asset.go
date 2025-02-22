package helper

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func GetAsset(path string) string {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	return host + ":" + port + "/public/" + path
}

func SaveBase64Image(base64Data string) (string, error) {
	// Decode the base64 data
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	// Generate unique filename for the image
	fileName := fmt.Sprintf("product-%d.jpg", time.Now().UnixNano())

	// Define the path to save the image
	filePath := fmt.Sprintf("storage/public/%s", fileName)

	// Create the directory if it doesn't exist
	err = os.MkdirAll("storage/public", os.ModePerm)
	if err != nil {
		return "", err
	}

	// Write the image data to a file
	err = os.WriteFile(filePath, imageData, 0644)
	if err != nil {
		return "", err
	}

	// Return the relative path (can be used for URL generation)
	return fileName, nil
}
