# 🧠 Capsule Shell

> An AI-powered CLI assistant that helps you generate and run shell commands interactively, safely, and with context awareness.

---

## ✨ Features

- 🗣️ **Natural Language to Shell Command**: Convert plain text prompts into shell commands.
- 🧠 **Descriptive Mode**: Get a detailed explanation before executing any command.
- 🔒 **Safe Execution**: Confirmation prompt before running any shell command.
- 🧪 **Execution Flag**: Only execute commands when `-x` flag is explicitly set.
- 🧹 **Command Extraction**: Extracts only valid Bash commands from AI responses, ignoring descriptions or comments.

---

## 🚀 Getting Started

### 🔧 Requirements

- Go 1.21+
- Internet access (to call OpenRouter API)

### 📦 Installation

```bash
git clone https://github.com/ak4bento/capsule-shell.git
cd capsule-shell
go mod tidy
```

---

## 🛠️ Usage

### 💬 Basic Prompt (No execution)

```bash
go run main.go "create a folder for golang learning"
```

You’ll see a friendly AI response containing a suggested command.

---

### 🔍 Descriptive Mode

```bash
go run main.go -d "delete all .tmp files in this directory"
```

You’ll receive a description **with** the shell command in a markdown code block.

---

### ⚡ Execute Prompt Result (after confirmation)

```bash
go run main.go -x "create backup folder and copy all .log files into it"
```

This will:
- Ask Capsule AI for a command,
- Show it to you,
- Ask: "Run this command? (y/n)",
- Then execute **only the extracted Bash code** if you approve.

---

### 🔍 + ⚡ Combine Descriptive and Execution

```bash
go run main.go -x -d "create backup folder and move all .bak files"
```

You’ll get the explanation and the command will be executed **only** if confirmed.

---

## 🧠 Behind the Scenes

Capsule Shell uses:

- OpenRouter API (supporting OpenAI, DeepSeek, etc.)
- `cobra` for CLI argument parsing
- Smart shell script extractor to ensure **only** Bash code is executed

---

## 🧱 Project Structure

```
.
├── cmd/
│   └── run.go        # CLI logic
├── chat/
│   └── chat.go      # Send prompt to OpenRouter
├── internal/
│   └── language.go     # Language shell command from AI response
│   └── readonly.go     # Render readonly not execute shell command from AI response
│   └── script.go       # Extracts and execute shell command from AI response
│   └── ui.go           # User Interface terminal AI response
├── main.go
└── go.mod
```

---

## 🔐 Security

This project does **not auto-execute** any command by default.

You **must** use the `-x` flag and confirm execution.  
Only the valid shell block (`bash`) inside the AI’s response is executed.

---

## ✍️ License

MIT © 2025 — [@ak4bento](https://github.com/ak4bento)
