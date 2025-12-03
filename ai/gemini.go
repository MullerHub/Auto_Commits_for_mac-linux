package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const API_URL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type RequestBody struct {
	Contents []Content `json:"contents"`
}

type Candidate struct {
	Content Content `json:"content"`
}

type Response struct {
	Candidates []Candidate `json:"candidates"`
}

func CallGemini(apiKey, diff string) string {
	body := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: fmt.Sprintf(
						"Generate a short, concise, one-line commit message for the following diff. "+
							"Keep it under 60 characters and include type like chore:, refactor:, feat:, fix:, docs:, test:\n%s", diff)},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", API_URL+"?key="+apiKey, bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println("Deu ruim no Gemini:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return ""
	}

	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		return strings.TrimSpace(response.Candidates[0].Content.Parts[0].Text)
	}
	return ""
}
