package license

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type licTplData struct {
	Year string
	Name string
}

func getLicTplData() licTplData {
	name, err := getGitConfigName()
	if err != nil {
		// In situations where git is not installed and username is not set,
		// the name part should be explicitly stated as '<your-name>'
		name = "<your-name>"
	}
	return licTplData{
		Year: fmt.Sprintf("%d", time.Now().Year()),
		Name: name,
	}
}

func getGitConfigName() (string, error) {
	f, err := exec.LookPath("git")
	if err != nil {
		return "", fmt.Errorf("git not found: %w", err)
	}

	n, err := exec.Command(f, "config", "--get", "user.name").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(n)), nil
}
