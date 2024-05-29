package main

import (
	"os"
	"prompter/internal/app"

	"github.com/spf13/cobra"
)

var (
	apiKey      string
	modelName   string
	temperature float32
	template    string
	streaming   bool
	rootCmd     = &cobra.Command{
		Use:   "prompter",
		Short: "Prompter is a tool to interact with OpenAI API.",
		Run: func(cmd *cobra.Command, args []string) {
			// Default action if no subcommands are provided
		},
	}
)

func init() {
	rootCmd.AddCommand(askCmd)
	rootCmd.AddCommand(cmd2)
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "OpenAI API key")

	askCmd.Flags().StringVarP(&modelName, "model", "m", "gpt-3.5-turbo", "Model name to use (default: gpt-3.5-turbo)")
	askCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.1, "Sampling temperature to use (default: 0.1)")
	askCmd.Flags().StringVar(&template, "template", "", "Path to the template file to use")
	askCmd.Flags().BoolVar(&streaming, "streaming", true, "Enable or disable streaming (default: true)")
}

var askCmd = &cobra.Command{
	Use:   "ask [query]",
	Short: "Ask a question to the OpenAI API. Reads input from stdin if no query is provided.",
	Run: func(cmd *cobra.Command, args []string) {
		app.ExecuteAsk(args, apiKey, modelName, temperature, template, streaming)
	},
}

var cmd2 = &cobra.Command{
	Use:   "cmd2",
	Short: "Test only",
	Run: func(cmd *cobra.Command, args []string) {
		app.ExecuteCmd2()
	},
}

func main() {
	// Fetch the API key from environment variable if not set via flag
	if apiKey = os.Getenv("OPENAI_API_KEY"); apiKey == "" {
		println("API key not set. Provide it via --api-key flag or OPENAI_API_KEY environment variable.")
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
