package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Request struct {
	Contents Contents `json:"contents"`
}

type Contents struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type Response struct {
	Candidates []Candidates `json:"candidates"`
}

type Candidates struct {
	Content Content `json:"content"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Gemini struct {
}

func (g *Gemini) Generate(text string) (string, error) {
	GEMINI_API_KEY := os.Getenv("GEMINI_API_KEY")
	GEMINI_BASE_URL := os.Getenv("GEMINI_BASE_URL")
	GEMINI_MODEL := os.Getenv("GEMINI_MODEL")

	parts := Part{Text: text}
	gemini := Contents{
		Parts: []Part{parts},
	}

	contents := map[string]Contents{}
	contents["contents"] = gemini

	jsonBody, err := json.Marshal(contents)

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/models/%s:generateContent",
		GEMINI_BASE_URL,
		GEMINI_MODEL)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("ContentType", "application/json")
	req.Header.Set("X-goog-api-key", GEMINI_API_KEY)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	json.Unmarshal(body, &response)

	return response.Candidates[0].Content.Parts[0].Text, nil
}
