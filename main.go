package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Access environment variables
    apiKey := os.Getenv("JENKINS_API_KEY")
    apiUrl := os.Getenv("JENKINS_API_URL")

    // Use the variables in your application
    fmt.Println("API Key:", apiKey)
    fmt.Println("API URL:", apiUrl)

    // Your application logic goes here
}

