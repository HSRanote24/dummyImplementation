package main

import (
	"bytes"
	"dummyImplementation/internal/database"
	"dummyImplementation/internal/middleware"
	"dummyImplementation/internal/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type KeycloakTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func getKeycloakToken() (string, error) {

	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	url := os.Getenv("KEYCLOAK_URL")
	clientID := os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	username := os.Getenv("KEYCLOAK_USERNAME")
	password := os.Getenv("KEYCLOAK_PASSWORD")

	if url == "" || clientID == "" || clientSecret == "" || username == "" || password == "" {
		return "", fmt.Errorf("missing required environment variables for Keycloak configuration")
	}

	data := []byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=password&username=%s&password=%s",
		clientID, clientSecret, username, password,
	))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to fetch token: %s", string(body))
	}

	var tokenResponse KeycloakTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	app := fiber.New()

	if err := database.Connect(); err != nil {
		panic(err)
	}

	token, err := getKeycloakToken()
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch Keycloak token: %v", err))
	}
	fmt.Printf("Keycloak Access Token: %s\n", token)

	app.Use(middleware.ErrorHandler)

	routes.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
