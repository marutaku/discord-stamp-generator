/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/marutaku/discord-stamp-generator/internal/stamp"
	"github.com/spf13/cobra"
)

type GenerateCmdOptions struct {
	OutputFilePath string
	FontPath       string
	FontColor      string
	OneLine        bool
}

func createOptionFromCmd(cmd *cobra.Command) (*GenerateCmdOptions, error) {
	options := &GenerateCmdOptions{}
	var err error

	options.OutputFilePath, err = cmd.Flags().GetString("output")
	if err != nil {
		return nil, err
	}
	options.FontPath, err = cmd.Flags().GetString("font-filepath")
	if err != nil {
		return nil, err
	}

	options.FontColor, err = cmd.Flags().GetString("font-color")
	if err != nil {
		return nil, err
	}

	options.OneLine, err = cmd.Flags().GetBool("one-line")
	if err != nil {
		return nil, err
	}
	return options, nil
}

func generateStamp(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify a text.")
		os.Exit(1)
	}
	text := args[0]
	options, err := createOptionFromCmd(cmd)
	if err != nil {
		fmt.Printf("Failed to parse options: %v\n", err)
		os.Exit(1)
	}
	fontColor := options.FontColor
	if fontColor == "" {
		fontColor, err = stamp.GenerateRandomHexColor()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	imageBytes, err := stamp.Generate(text, 360, 360, fontColor, options.FontPath, options.OneLine)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Export image
	if options.OutputFilePath == "" {
		os.Stdout.Write(imageBytes)
	} else {
		err = stamp.Export(imageBytes, options.OutputFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a stamp image",
	Long: `Generate a stamp image from a text.
The generated image is saved as a PNG file and optimized for Discord.
For example:
	stamp generate "Hello, World!" -o hello.png
`,
	Run: generateStamp,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("output", "o", "", "Output file path")
	generateCmd.Flags().StringP("font-filepath", "f", "", "External font path")
	generateCmd.Flags().StringP("font-color", "c", "", "Font color")
	generateCmd.Flags().BoolP("one-line", "", false, "Force to generate a stamp image in one line")
}
