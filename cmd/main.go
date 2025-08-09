package main

import (
	"github.com/dlworhd/gemini-go/model"
	"github.com/joho/godotenv"
)

// GEMINI_BASE_URL=https://generativelanguage.googleapis.com/v1beta
// GEMINI_API_KEY=your-api-key
// GEMINI_MODEL=gemini-2.0-flash

func main() {
	godotenv.Load()
	gemini := model.Gemini{}

	prompt := "your prompt"
	gemini.Generate(prompt)
}
