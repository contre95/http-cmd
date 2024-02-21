package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

var SHADERS_PATH string = "/home/contre/.config/hypr/shaders"

func handleListShaders(w http.ResponseWriter, r *http.Request) {
	// Define the command to execute
	cmd := exec.Command("ls", SHADERS_PATH)
	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to execute command", http.StatusInternalServerError)
		return
	}
	// Write the command output as the HTTP response
	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func handleApplyShader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Get the command from the request body
	shader := r.FormValue("shader")
	if shader == "" {
		http.Error(w, "No command provided", http.StatusBadRequest)
		return
	}
	// Define the command to execute
	execCmd := exec.Command("sh", "-c", fmt.Sprintf("hyprctl keyword decoration:screen_shader %s/%s", SHADERS_PATH, shader))
	// Execute the command
	output, err := execCmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to execute command", http.StatusInternalServerError)
		return
	}

	// Write the command output as the HTTP response
	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func handleResetShader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Define the command to execute
	execCmd := exec.Command("sh", "-c", "hyprctl keyword decoration:screen_shader [[EMPTY]]")
	// Execute the command
	output, err := execCmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to execute command", http.StatusInternalServerError)
		return
	}

	// Write the command output as the HTTP response
	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func main() {
	// Define the HTTP endpoint
	http.HandleFunc("/shaders", handleListShaders)
	http.HandleFunc("/shade", handleApplyShader)
	http.HandleFunc("/reset", handleResetShader)

	// Start the HTTP server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
