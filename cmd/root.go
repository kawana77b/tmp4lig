package cmd

import (
	"embed"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/kawana77b/tmp4lig/internal/commands"
	"github.com/kawana77b/tmp4lig/internal/prompt"
	"github.com/spf13/cobra"
)

var Version string = "0.0.0"

var Fs embed.FS

type rootProps struct {
	fileType string
	saveDir  string
}

var rootP = &rootProps{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tmp4lig",
	Short: "Template Output Tool for license and gitignore.",
	Long: `Template Output Tool for license and gitignore.

Example:
  # outputs a .gitignore or LICENSE file to the current current directory.
	 $ tmp4lig
  # outputs a .gitignore file to the current current directory.
	 $ tmp4lig gitignore
  # outputs a LICENSE file to the current current directory.
	 $ tmp4lig license

For more information, see:
  https://github.com/kawana77b/tmp4lig
	`,
	Args:    cobra.MaximumNArgs(1),
	PreRunE: preRunRoot,
	RunE:    runRoot,
}

func Execute() {
	rootCmd.Version = Version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func preRunRoot(cmd *cobra.Command, args []string) error {
	// set the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	rootP.saveDir = wd

	// set the file type
	var validFileTypes = []string{"license", "gitignore"}
	var fileType string
	if len(args) == 0 {
		ans, err := prompt.NewSelectQuestion().Ask("Select a file type", validFileTypes)
		if err != nil {
			// if the user cancels the prompt, just exit
			os.Exit(0)
		}
		fileType = ans
	} else {
		arg := args[0]
		if !slices.Contains(validFileTypes, arg) {
			return fmt.Errorf("invalid argument: %s, argument must be one of %s", arg, strings.Join(validFileTypes, ", "))
		}
		fileType = arg
	}
	rootP.fileType = fileType

	return nil
}

func runRoot(cmd *cobra.Command, args []string) error {
	command, ok := commands.GetCommand(rootP.fileType)
	if !ok {
		return fmt.Errorf("fileType not found: %s", rootP.fileType)
	}

	command.Set(commands.CommandProps{
		Fs:      Fs,
		SaveDir: rootP.saveDir,
	})
	err := command.Execute()
	if err != nil {
		return err
	}

	return nil
}
