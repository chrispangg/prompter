# Prompter - A Simple CLI Tool for OpenAI API with Templates

Prompter is a command-line tool designed to interact with the OpenAI API using predefined templates. It supports various options to customize the API requests.

## Features

-   Interact with OpenAI API using CLI
-   Support for predefined templates
-   Adjustable model, temperature, and streaming options

## Installation

```sh
go install github.com/chrispangg/prompter@latest
```

## Usage

Ask a question to the OpenAI API. Reads input from stdin if no query is provided.

```sh
prompter ask "Explain the theory of relativity"
```

Use pbpaste to pipe clipboard content:

```sh
pbpaste | prompter ask
```

Override the model and temperature with flags:

```sh
pbpaste | prompter ask --api-key your_api_key_here --model "gpt-3.5-turbo" --temperature 0.9 --template templates/example.tmpl
```

### Options

-   `--api-key` or `OPENAI_API_KEY`: Set your OpenAI API key.
-   `--model` or `-m`: Specify the model to use (default: `gpt-3.5-turbo`).
-   `--temperature` or `-t`: Set the sampling temperature (default: `0.1`).
-   `--template`: Path to the template file to use.
-   `--streaming`: Enable or disable streaming (default: `true`).

### Example with Template

```sh
prompter ask "Explain the theory of relativity" --template=templates/write_essay.tmpl
```

## Templates

Templates are used to format the input query. Example templates are provided in the `templates` directory.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.

## Acknowledgements

This project uses the following libraries:

-   [go-openai](https://github.com/sashabaranov/go-openai)
-   [cobra](https://github.com/spf13/cobra)
