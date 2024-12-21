package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"auth-jwt/database"
	"auth-jwt/handler"
	"auth-jwt/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Helper function to setup the Fiber app
func setupApp() *fiber.App {
	app := fiber.New()
	// Add routes to your app...
	app.Post("/user", handler.CreateUser)
	app.Get("/user/:id", handler.GetUser)
	app.Put("/user/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)
	return app
}

func createTestUser(app *fiber.App, username, password string) (*model.User, string) {
	user := model.User{Username: username, Password: password, Email: username + "@example.com"}
	hash, _ := handler.HashPassword(password)
	user.Password = hash

	database.DB.Create(&user)

	// Generate JWT for the user
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		// Handle error accordingly; here we could panic for testing purposes
		panic("Error signing token: " + err.Error())
	}

	return &user, tokenString
}

func TestCreateUser(t *testing.T) {
	app := setupApp()

	// Test data
	userData := map[string]string{
		"username": "testUser",
		"password": "password123",
		"email":    "testuser@example.com",
	}

	data, _ := json.Marshal(userData)

	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status code %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestGetUser(t *testing.T) {
	app := setupApp()
	user, token := createTestUser(app, "testuser", "password123")

	req := httptest.NewRequest(http.MethodGet, "/user/"+strconv.Itoa(int(user.ID)), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := app.Test(req, -1)

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status code %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}

// Additional tests for UpdateUser and DeleteUser can be added similarly
