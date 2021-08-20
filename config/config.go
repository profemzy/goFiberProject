package config

import (
	"os"
)

// Config function to get env values from key
func Config(key string) string {
	// load .env file
	//err := godotenv.Load(".env")
	//if err != nil {
	//	fmt.Print("Error loading .env file")
	//}
	return os.Getenv(key)
}
