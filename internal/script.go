package internal

import (
	"os"
	"os/exec"
  "regexp"
  "strings"
)

func ExecuteScript(scriptText string) error {
	scriptFile := "temp_script.sh"

	err := os.WriteFile(scriptFile, []byte(scriptText), 0755)
	if err != nil {
		return err
	}
	defer os.Remove(scriptFile) // auto delete setelah dijalankan

	cmd := exec.Command("sh", scriptFile)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func ExtractShellCommand(text string) string {
	re := regexp.MustCompile("(?s)```bash\\s*(.*?)\\s*```")
	match := re.FindStringSubmatch(text)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return text // fallback
}
