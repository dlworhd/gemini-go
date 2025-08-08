package main

import (
	"github.com/dlworhd/gemini-go/model"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	gemini := model.Gemini{}

	prompt := "Hello World!"
	gemini.Generate(prompt)
}
