package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIRequest struct {
	CountryName string `json:"countryName"`
}

type AIResponse struct {
	Description string `json:"description"`
}

// Structure of Ollama API request
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// Structure of Ollama API response
type OllamaResponse struct {
	Response string `json:"response"`
}

func HandleAI(c *gin.Context) {
	var req AIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	prompt := "Give a short 2-3 sentence description about the country " + req.CountryName + " suitable for travelers."

	ollamaReq := OllamaRequest{
		Model:  "llama3", // or "mistral", "phi", etc. depending on what you have downloaded
		Prompt: prompt,
	}

	reqBody, err := json.Marshal(ollamaReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to build request"})
		return
	}

	resp, err := http.Post("http://host.docker.internal:11434/api/generate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contact AI model"})
		return
	}
	defer resp.Body.Close()

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse AI response"})
		return
	}

	c.JSON(http.StatusOK, AIResponse{Description: ollamaResp.Response})
}
