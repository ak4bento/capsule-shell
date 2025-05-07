package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ak4bento/capsule-shell/chat"
	"github.com/ak4bento/capsule-shell/internal"
	"github.com/spf13/cobra"
)

const (
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Green  = "\033[32m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
	Bold   = "\033[1m"
)

var (
	describe bool
	execute  bool
)

var rootCmd = &cobra.Command{
	Use:   "capsule-shell",
	Short: "Capsule Shell - CLI AI Assistant",
	Long:  `Capsule Shell is AI CLI assistant to help you run shell commands.`,
	Args:  cobra.ArbitraryArgs, // Accept any number of arguments
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("%s⚠️  Empty Prompt, Enter the command or question for AI capsule.%s\n", Bold+Red, Reset)
			fmt.Println("Example: capsule-shell \"make a folder config and file index.js\"")
			return
		}

		input := strings.Join(args, " ")

		var response string
		var err error

		if describe {
			response, err = chat.SendDescriptivePrompt(input)
		} else {
			response, err = chat.SendPrompt(input)
		}

		if err != nil {
			fmt.Printf("%s❌ Error: %s%s\n", Bold+Red, err.Error(), Reset)
			os.Exit(1)
		}

		clean := strings.TrimSpace(response)
		if !describe {
			clean = strings.ReplaceAll(clean, "```bash", "")
			clean = strings.ReplaceAll(clean, "`", "")
		}
		cleanResponse := strings.TrimSpace(clean)

		fmt.Printf("%s🧠 Capsule AI:%s\n%s%s%s", Bold+Green, Reset, Cyan, cleanResponse, Reset)

		if cleanResponse == "I am capsule shell command line interpreter" {
			return
		}

		if execute {
			fmt.Printf("\n%s🚀 Running command...%s", Bold+Green, Reset)

			fmt.Printf("\n%s🔒 Running the command? (y/n): %s", Bold+Yellow, Reset)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if strings.ToLower(input) != "y" {
				fmt.Printf("%s🚫 Command execution canceled.%s\n", Bold+Red, Reset)
			} else {
				// Extract shell command from response
				fmt.Printf("%s🚀 Extracting shell command...%s", Bold+Green, Reset)
				shellCommand := internal.ExtractShellCommand(cleanResponse)
				err := internal.ExecuteScript(shellCommand)
				if err != nil {
					fmt.Printf("%sFailed to run command%s\n%s", Bold+Red, Reset, err)
				}
			}
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&describe, "describe", "d", false, "Explain command without running it")
	rootCmd.Flags().BoolVarP(&execute, "execute", "x", false, "Run command after explanation")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
