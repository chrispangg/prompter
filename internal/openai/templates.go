package openai

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

const (
	templateBaseURL  = "https://raw.githubusercontent.com/chrispangg/prompter/main/templates/"
	localTemplateDir = "templates"
)

func downloadTemplate(templateName string) (string, error) {
	url := templateBaseURL + templateName + ".tmpl"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download template: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download template: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read template body: %w", err)
	}

	// Save the template locally for future use
	localPath := filepath.Join(localTemplateDir, templateName+".tmpl")
	if err := os.MkdirAll(localTemplateDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create local template directory: %w", err)
	}
	if err := os.WriteFile(localPath, body, 0644); err != nil {
		return "", fmt.Errorf("failed to save template locally: %w", err)
	}

	return string(body), nil
}

func LoadTemplate(templateName string) (string, error) {
	localPath := filepath.Join(localTemplateDir, templateName+".tmpl")
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		// Template not found locally, download it
		return downloadTemplate(templateName)
	}

	// Load the template from the local file
	data, err := os.ReadFile(localPath)
	if err != nil {
		return "", fmt.Errorf("failed to read local template: %w", err)
	}

	return string(data), nil
}

func FillTemplate(templateStr string, data map[string]string) (string, error) {
	tmpl, err := template.New("queryTemplate").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var filledTemplate bytes.Buffer
	if err := tmpl.Execute(&filledTemplate, data); err != nil {
		return "", err
	}

	return filledTemplate.String(), nil
}
