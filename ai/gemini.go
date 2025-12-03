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

func CallGemini(apiKey, diff string) (string, error) {
	prompt := fmt.Sprintf(
		"You are a git commit assistant. Analyze the diff below and generate a SINGLE, short, concise commit message (max 60 chars). "+
			"Use Conventional Commits (feat:, fix:, chore:, refactor:, etc). "+
			"CRITICAL RULES: 1. Output ONLY the raw text message. 2. Do NOT use markdown or code blocks. 3. Do NOT repeat the diff.\n\nDiff:\n%s", diff)

	body := RequestBody{
		Contents: []Content{
			{Parts: []Part{{Text: prompt}}},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("erro json: %v", err)
	}

	req, err := http.NewRequest("POST", API_URL+"?key="+apiKey, bytes.NewReader(jsonData))
	if err != nil {
		return "", fmt.Errorf("erro req: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro net: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro leitura: %v", err)
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("api erro %d: %s", resp.StatusCode, string(data))
	}

	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return "", fmt.Errorf("erro parse: %v", err)
	}

	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		text := response.Candidates[0].Content.Parts[0].Text
		// Limpeza de segurança: removemos qualquer markdown que a IA tentar colocar
		text = strings.ReplaceAll(text, "```", "")
		text = strings.ReplaceAll(text, "diff", "")
		return strings.TrimSpace(text), nil
	}

	return "", fmt.Errorf("sem resposta da IA")
}
