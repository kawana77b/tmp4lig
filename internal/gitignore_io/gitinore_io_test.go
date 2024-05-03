package gitignoreio_test

import (
	"strings"
	"testing"

	gitignoreio "github.com/kawana77b/tmp4lig/internal/gitignore_io"
)

func Test_GetFileBytes(t *testing.T) {
	gio := gitignoreio.NewGitignoreIO()
	bytes, err := gio.GetGitignoreData("node", "react")
	if err != nil {
		t.Error(err)
	}

	if len(bytes) == 0 {
		t.Error("empty bytes")
	}

	if strings.Contains(string(bytes), "node_modules/") == false {
		t.Error("node_modules/ is not included")
	}
}

func Test_GetList(t *testing.T) {
	gio := gitignoreio.NewGitignoreIO()
	list, err := gio.GetList()
	if err != nil {
		t.Error(err)
	}

	if len(list) == 0 {
		t.Error("empty list")
	}
}
