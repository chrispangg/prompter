package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"prompter/internal/openai"
)

func readFromStdin() (string, error) {
	var input strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return input.String(), nil
}

func ExecuteAsk(args []string, apiKey, model string, temperature float32, templateFile string, stream bool) {
	var query string
	if len(args) == 0 {
		fmt.Println("Reading input from stdin... Press Ctrl+D to end input.")
		// Read from stdin
		stdinQuery, err := readFromStdin()
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}
		query = stdinQuery
	} else {
		// Read from args
		query = strings.Join(args, " ")
	}

	if query == "" {
		fmt.Println("Please provide a question.")
		return
	}

	if stream {
		openai.MakeAPICallStream(query, apiKey, model, temperature, templateFile)
	} else {
		response := openai.MakeAPICall(query, apiKey, model, temperature, templateFile)
		fmt.Println(response)
	}
}
