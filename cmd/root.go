package cmd

import (
	// "bufio"
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
	satire   bool
)

var rootCmd = &cobra.Command{
	Use:   "capsule-shell [args]",
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

		response, err = FlagingOptions(input)

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

		if cleanResponse == "I am capsule shell command line interpreter" || satire {
			internal.RenderOnly(cleanResponse)
			return
		}
		// show UI
		approved := true

		if execute {
			approved = internal.RunUI(cleanResponse)
			if approved {
				// Extract shell command from response
				shellCommand := internal.ExtractShellCommand(cleanResponse)
				err := internal.ExecuteScript(shellCommand)
				if err != nil {
					fmt.Printf("%sFailed to run command%s\n%s", Bold+Red, Reset, err)
				}
			}
		} else {
			internal.RenderOnly(cleanResponse)
		}
	},
}

func FlagingOptions(input string) (response string, err error) {
	if describe {
		response, err = chat.SendDescriptivePrompt(input)
	} else {
    response, err = chat.SendMainPrompt(input)
  }

  if satire {
    response, err = chat.SendSatiricalPrompt(input)
  }
	return response, err
}

func init() {
	rootCmd.Flags().BoolVarP(&describe, "describe", "d", false, "Explain command without running it")
	rootCmd.Flags().BoolVarP(&execute, "execute", "x", false, "Run command after explanation")
	rootCmd.Flags().BoolVarP(&satire, "satire", "s", false, "This is jokes for developer")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
