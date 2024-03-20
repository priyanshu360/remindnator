package util

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/priyanshu360/remindnator/config"
	"golang.org/x/oauth2"
)

func LoadTokenFromFile() error {
	// Check if token file exists
	if _, err := os.Stat("token.json"); err != nil {
		return err
	}
	// Token file exists, load token
	file, err := os.ReadFile("token.json")
	if err != nil {
		fmt.Println("Error reading token file:", err)
		return err
	}
	var token oauth2.Token
	err = json.Unmarshal(file, &token)
	if err != nil {
		fmt.Println("Error unmarshalling token:", err)
		return err
	}
	// Set global client with loaded token
	config.CLIENT = oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&token))
	fmt.Println("Loaded token from file.")
	return nil
}

func LoadLocalEnvFile() error {
	// TODO #10 : Make Path dynamic for local.env and token.json
	if _, err := os.Stat("C:/Users/priya/OneDrive/Desktop/toyProjects/remindnator/util/local.env"); err == nil {
		if err := godotenv.Load("C:/Users/priya/OneDrive/Desktop/toyProjects/remindnator/util/local.env"); err != nil {
			return err
		}
		config.Refresh()
	} else {
		return err
	}
	return nil
}

func GenerateUUID(length int) string {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
