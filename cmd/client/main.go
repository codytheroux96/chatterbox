package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select an option: \n1. Register\n2. Login\n3. Exit")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("Enter Username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Println("Enter Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			if username == "" || password == "" {
				fmt.Println("Error: Username and Password cannot be empty")
				return
			}
			if len(password) < 8 {
				fmt.Println("Error: Password must be at least 8 characters")
			}

			payload := map[string]string{
				"username": username,
				"password": password,
			}

			jsonBody, err := json.Marshal(payload)
			if err != nil {
				fmt.Println("Failed to serialize payload:", err)
				return
			}

			resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(jsonBody))
			if err != nil {
				fmt.Println("Error: Failed to send the request:", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
				fmt.Println("Registration Successful")
			} else {
				body, _ := io.ReadAll(resp.Body)
				fmt.Printf("Error: %s\n", string(body))
			}
		case "2":
			fmt.Println("Enter Username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Println("Enter Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			if username == "" || password == "" {
				fmt.Println("Error: Username and Password cannot be empty")
				continue
			}

			payload := map[string]string{
				"username": username,
				"password": password,
			}

			jsonBody, err := json.Marshal(payload)
			if err != nil {
				fmt.Println("Failed to serialize payload:", err)
				return
			}

			resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonBody))
			if err != nil {
				fmt.Println("Error: Failed to send the request:", err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				fmt.Printf("Login Successful. Token: %s\n", string(body))
			} else {
				body, _ := io.ReadAll(resp.Body)
				fmt.Printf("Error: %s\n", string(body))
			}
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
