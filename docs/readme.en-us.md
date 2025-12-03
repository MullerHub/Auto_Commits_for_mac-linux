# 🌙 Luna — AI Commit Generator

Luna generates concise commit messages for your staged files using Google Gemini 2.0 Flash API.

## ✨ Features

- **Per-file commits**: one commit for each staged file
- **Gemini 2.0 Flash**: AI-powered summaries from file diffs
- **Conventional prefixes**: adds prefix if missing
- **Length control**: target < 60 chars, configurable max (default 72)
- **Smart filtering**: ignores common binaries and images
- **Optional emojis**: enable with `-e`

## How it works

1. Collects staged files via `git diff --cached --name-only`
2. Sends each diff to `gemini-2.0-flash:generateContent`
3. If response lacks known prefix, randomly selects from: `chore:`, `refactor:`, `feat:`, `fix:`, `docs:`, `test:`, `etc..`
4. If `-e` is active, adds random emoji
5. Truncates to `maxCommitLength` and commits with `git commit -m <message> -- <file>`

## Requirements

- macOS | Linux | UNIX-like systems
- Git installed and available in `PATH`
- Google Gemini API key (`https://aistudio.google.com/app/apikey`)

## Installation

### Option A — Use prebuilt binary (`bin/luna`)

1. Copy `bin/luna` to a system directory, for example `/usr/local/bin`:

```bash
sudo cp bin/luna /usr/local/bin/
sudo chmod +x /usr/local/bin/luna
```

Or add the folder containing the binary to your `PATH` (add the line to `~/.zshrc` or `~/.bashrc`):

```bash
export PATH="$PATH:/path/to/your/folder"
source ~/.zshrc
```

### Option B — Build from source (Go)

```bash
go build -o ./bin/luna main.go
```

Or use the helper script:

```bash
./build.sh
```

## Configuration

Luna reads configuration from project and global files:

- Project: `.lunacfg` (in repository root or nearest parent)
- Global: `.lunarc` (in user home directory)

Priority:

- API key: Global → Project → Default
- Other settings: Project → Default

Default settings (from code):

- `ignoredPatterns`: `*.exe`, `*.dll`, `*.png`, `*.jpg`, `*.jpeg`, `*.gif`, `*.bin`
- `commitPrefixes`: `chore:`, `refactor:`, `feat:`, `fix:`, `docs:`, `test:`
- `maxCommitLength`: `72`
- `defaultEmoji`: `false`

### Set your API key

```bash
LunaApikey YOUR_GEMINI_KEY
```

This saves the key in your global `.lunarc`. Reopen terminal after setting.

## Usage

Run Luna inside a Git repository with staged changes.

### Commands and aliases

- `LunaHelp` | `lh` | `-lh`: Show help
- `LunaCommit` | `lc` | `-lc`: Generate and commit per-file messages
- `LunaApikey <YOUR_KEY>` | `lkey <YOUR_KEY>` | `-lkey <YOUR_KEY>`: Set API key
- `LunaConfig` | `config` | `-config` with subcommands:
  - `init`: Create `.lunacfg` in current directory
  - `show`: Print merged configuration
  - `edit`: Placeholder (not implemented yet)

You can call them as executable arguments (e.g., `Luna -c`, `Luna lc`, `Luna -lh`) or directly as command names if your shell exposes them.

### Typical flow

```bash
git add .
Luna -c          # or: Luna lc, or: LunaCommit
```

### Optional emojis

```bash
Luna -lc -e       # enable emojis in messages
```

## Example output

```
Generating commit for file: src/main.go
Committed src/main.go with message:
🚀 feat: add user authentication system

Generating commit for file: README.md
Committed README.md with message:
📝 docs: update installation instructions
```

## Notes

- Luna skips common binary/image files
- If model returns empty response, fallback is `update <file>`
- Supported prefixes: `feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `chore:`
- `maxCommitLength` is enforced (default 72)

## Troubleshooting

- Error: `Set API key using LunaApikey first`
  - Run `LunaApikey YOUR_KEY` and reopen terminal
- Error running Git commands
  - Ensure you're in a Git repository and Git is installed
- No staged changes
  - Run `git add .` or stage specific files
- API key not working
  - Verify key is valid and has access to Gemini 2.0 Flash

---

Made with ❤️ by hax — version 1.3 (Beta)
