package handler

import (
	"applaza-backend/model"
	"applaza-backend/service"
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"net/http"
	"regexp"
	"time"
)

var mySigningKye = []byte("secret")

func signInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one sign in request")
	w.Header().Set("Content-Type", "text/plain")

	// Get User information from client
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client %v\n", err)
		return
	}

	exists, err := service.CheckUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to read user from Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to read user data from Elasticsearch %v\n", err)
		return
	}
	if !exists {
		http.Error(w, "User doesn't exist or inputs wrong password", http.StatusUnauthorized)
		fmt.Printf("User doesn't exist or inputs wrong password\n")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(mySigningKye)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Printf("Failed to generate token %v\n", err)
		return
	}
	w.Write([]byte(tokenString))
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one sign up request")
	w.Header().Set("Content-Type", "text/plain")

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client %v\n", err)
		return
	}
	if user.Username == "" || user.Password == "" || regexp.MustCompile(`^[a-z0-9]$`).MatchString(user.Username) {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		fmt.Println("Invalid username or password")
		return
	}

	success, err := service.AddUser(&user)
	if err != nil {
		http.Error(w, "Failed to save user to Elasticsearch", http.StatusInternalServerError)
		fmt.Println("Failed to save user to Elasticsearch")
		return
	}
	if !success {
		http.Error(w, "User already exists", http.StatusBadRequest)
		fmt.Println("User already exists")
		return
	}
	fmt.Printf("User is added successfully: %s\n", user.Username)
}
