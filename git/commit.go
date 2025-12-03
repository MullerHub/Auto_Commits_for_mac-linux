package git

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"

	"luna/ai"
	"luna/config"
)

var emojis = []string{"✨", "🛠️", "🐛", "🔥", "📝", "🚀", "🔧", "🎨", "🔒", "💄"}

func GenerateCommitMessage(apiKey, diff, filename string, cfg config.Config, includeEmoji bool) string {
	commitMsg, err := ai.CallGemini(apiKey, diff)

	// Se der erro, avisamos no terminal e usamos o fallback
	if err != nil {
		fmt.Printf("Gemini error for %s: %v\n", filename, err)
		commitMsg = ""
	}

	if commitMsg == "" {
		commitMsg = "update " + filename
	}

	hasPrefix := false
	for _, p := range cfg.CommitPrefixes {
		if strings.HasPrefix(strings.ToLower(commitMsg), strings.ToLower(p)) {
			hasPrefix = true
			break
		}
	}

	if !hasPrefix && len(cfg.CommitPrefixes) > 0 {
		// Removido rand.Seed antigo
		prefix := cfg.CommitPrefixes[rand.Intn(len(cfg.CommitPrefixes))]
		commitMsg = prefix + " " + commitMsg
	}

	if includeEmoji {
		emoji := emojis[rand.Intn(len(emojis))]
		commitMsg = fmt.Sprintf("%s %s", emoji, commitMsg)
	}

	if len(commitMsg) > cfg.MaxCommitLength {
		commitMsg = commitMsg[:cfg.MaxCommitLength-3] + "..."
	}

	return commitMsg
}

func CommitFile(filename, commitMsg string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", commitMsg, "--", filename)
	out, err := cmd.CombinedOutput()
	if err != nil {
		// Aproveitando para corrigir aquele outro erro de mensagem oculta
		return string(out), fmt.Errorf("git failed: %s (Err: %v)", string(out), err)
	}
	return string(out), nil
}
