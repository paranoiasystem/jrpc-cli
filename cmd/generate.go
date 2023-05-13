package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jrpc-cli/generate"
	"jrpc-cli/util"
	"os"
)

var input string
var output string
var language string

func init() {
	generateCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input file name")
	generateCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output file name")
	generateCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "Language to generate code for")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate JSON-RPC Schema complaint interface code",
	Long:  `Generate JSON-RPC Schema complaint interface code.`,
	Run: func(cmd *cobra.Command, args []string) {
		if input == "" {
			fmt.Println("input file name is required")
			return
		}

		if output == "" {
			fmt.Println("output file name is required")
			return
		}

		if language == "" {
			fmt.Println("language is required")
			return
		}

		json, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("error reading file: ", err)
			return
		}

		var code string

		switch language {
		case "go":
			fmt.Println("Generating Go code\n Not available yet")
		case "java":
			fmt.Println("Generating Java code\n Not available yet")
		case "python":
			fmt.Println("Generating Python code\n Not available yet")
		case "typescript":
			code, err = generate.Typescript(json)
			if err != nil {
				fmt.Println("error generating TypeScript code: ", err)
				return
			}
			fmt.Println("TypeScript code generated successfully")
		default:
			fmt.Println("Language not supported")
			return
		}

		err2 := util.Writefile(output, code)
		if err2 != nil {
			fmt.Println("error writing to file: ", err2)
			return
		}

	},
}
