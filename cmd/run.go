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
//
// const (
// 	Yellow = "\033[33m"
// 	Cyan   = "\033[36m"
// 	Green  = "\033[32m"
// 	Red    = "\033[31m"
// 	Reset  = "\033[0m"
// 	Bold   = "\033[1m"
// )
//
// var (
// 	describe bool
// )

var runCmd = &cobra.Command{
	Use:   "run [message]",
	Short: "Run shell command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		var response string
		var err error

		if describe {
			response, err = chat.SendDescriptivePrompt(input)
		} else {
			response, err = chat.SendPrompt(input)
		}

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		clean := strings.TrimSpace(response)
		clean = strings.ReplaceAll(clean, "```bash", "")
		clean = strings.ReplaceAll(clean, "`", "")
		cleanResponse := strings.TrimSpace(clean)
		// fmt.Println("🧠 Capsule AI: \n", cleanResponse)
		fmt.Printf("%s🧠 Capsule AI:%s\n%s%s%s", Bold+Green, Reset, Cyan, cleanResponse, Reset)

		if "I am capsule shell command line interpreter" != cleanResponse {
			if !describe {
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
		}
	},
}

func init() {
	runCmd.Flags().BoolVarP(&describe, "describe", "d", false, "Explain command, Don't run command")
	rootCmd.AddCommand(runCmd)
}
