package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type AIRequest struct {
	CountryName string `json:"countryName"`
}

type AIResponse struct {
	Description string `json:"description"`
}

func HandleAI(c *gin.Context) {
	var req AIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	prompt := "Give a short 1-2 sentence description about the country " + req.CountryName + " suitable for travelers."

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, // or GPT-4 if you want
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful travel guide.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch AI response"})
		return
	}

	answer := resp.Choices[0].Message.Content

	c.JSON(http.StatusOK, AIResponse{Description: answer})
}
