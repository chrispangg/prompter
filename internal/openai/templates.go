package openai

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

func LoadTemplate(templatePath string) (string, error) {
	absPath, err := filepath.Abs(templatePath)
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
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
