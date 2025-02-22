package jwtgenerator

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func GenerateAndWriteSecretKey() error {
	// Generate random bytes
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return fmt.Errorf("error generating random bytes: %v", err)
	}

	// Encode random bytes to base64
	secretKey := base64.StdEncoding.EncodeToString(randomBytes)

	// Write secret key to .env file
	err = writeEnvFile(".env", "JWT_SECRET_KEY", secretKey)
	if err != nil {
		return fmt.Errorf("error writing to .env file: %v", err)
	}

	fmt.Println("JWT secret key has been generated and written to .env file.")
	return nil
}

// Function to write key-value pair to .env file
func writeEnvFile(filename, key, value string) error {
	// Check if file exists, if not, create it
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	// Read existing content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Append or update key-value pair
	updatedContent := appendOrReplaceEnvVar(content, key, value)

	// Write updated content to file
	err = ioutil.WriteFile(filename, updatedContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Function to append or replace key-value pair in .env content
func appendOrReplaceEnvVar(content []byte, key, value string) []byte {
	lines := string(content)
	newLine := "\n" + key + "=\"" + value + "\"\n"
	if idx := findEnvVarIndex(lines, key); idx != -1 {
		// If key exists, replace its value
		start := idx
		end := start + len(key)
		for end < len(lines) && lines[end] != '\n' && lines[end] != '\r' {
			end++
		}
		lines = lines[:start] + newLine + lines[end:]
	} else {
		// If key does not exist, append new line
		lines += newLine
	}
	return []byte(lines)
}

// Function to find index of key in .env content
func findEnvVarIndex(content, key string) int {
	index := -1
	lines := len(content)
	for i := 0; i < lines; {
		j := i
		for ; j < lines && content[j] != '='; j++ {
		}
		if j < lines && content[i:j] == key {
			index = i
			break
		}
		i = j + 1
		for ; i < lines && (content[i] == '\n' || content[i] == '\r'); i++ {
		}
	}
	return index
}
