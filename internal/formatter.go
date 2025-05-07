package internal

import (
	"fmt"
	"strings"
)

const bubbleTop = "╭─────────────────────────────────────────────╮"
const bubbleBot = "╰─────────────────────────────────────────────╯"

func PrintUserBubble(content string) {
	fmt.Println("\n🧑‍💻 Kamu:")
	printBubble(content)
}

func PrintBotBubble(content string) {
	fmt.Println("\n🤖 Capsule Shell:")
	printBubble(content)
}

func printBubble(content string) {
	lines := strings.Split(content, "\n")
	// fmt.Println(bubbleTop)
	for _, line := range lines {
		fmt.Println(line)
	}
	// fmt.Println(bubbleBot)
}
