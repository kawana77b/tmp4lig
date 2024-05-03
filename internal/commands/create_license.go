package commands

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kawana77b/tmp4lig/internal/license"
	"github.com/kawana77b/tmp4lig/internal/prompt"
	"github.com/logrusorgru/aurora"
)

type CreateLicenseCommand struct {
	fs      embed.FS
	saveDir string
}

func (c *CreateLicenseCommand) Set(props CommandProps) {
	c.fs = props.Fs
	c.saveDir = props.SaveDir
}

func (c *CreateLicenseCommand) Execute() error {
	ans, err := prompt.NewSelectQuestion().Ask("Select a license", license.GetLicNames())
	if err != nil {
		// if the user cancels the prompt, just exit
		return nil
	}

	info, ok := license.GetLicInfoByName(ans)
	if !ok {
		return nil
	}

	data, err := license.NewLicenseFs(c.fs).ReadFile(info.Path())
	if err != nil {
		return err
	}

	savePath := filepath.Join(c.saveDir, "LICENSE")
	os.WriteFile(savePath, data, 0644)

	fmt.Printf("%s %s\n", aurora.Cyan("created:"), savePath)

	return nil
}
