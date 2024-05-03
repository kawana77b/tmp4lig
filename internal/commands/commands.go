package commands

var comamnds = map[string]ICommand{
	"license":   &CreateLicenseCommand{},
	"gitignore": &CreateGitIgnoreCommand{},
}

// returns a command by name and a boolean value indicating whether the command exists
func GetCommand(name string) (ICommand, bool) {
	cmd, ok := comamnds[name]
	return cmd, ok
}
