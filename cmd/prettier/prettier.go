package prettier

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

const (
	configFilename = ".prettierrc"
	configContent  = `{
	"printWidth": 100,
	"tabWidth": 2,
	"semi": false,
	"singleQuote": true,
	"quoteProps": "consistent",
	"jsxSingleQuote": true,
	"trailingComma": "es5",
	"arrowParens": "avoid"
}`
)

var (
	useGlobal bool
	RootCmd   = &cobra.Command{
		Use:   "prettier",
		Short: "Create prettier config",
		RunE:  execCmd,
	}
)

func init() {
	bindCmdFlags()
}

func bindCmdFlags() {
	RootCmd.Flags().BoolVarP(&useGlobal, "global", "g", false, "apply to the global config")
}

func execCmd(cmd *cobra.Command, args []string) error {
	filePath, err := resolveConfigFilePath(configFilename, useGlobal)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	content := genConfigContent()
	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}

func resolveConfigFilePath(filename string, useGlobal bool) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if useGlobal {
		dir, err = os.UserHomeDir()
		if err != nil {
			return "", err
		}
	}

	return path.Join(dir, filename), nil
}

func genConfigContent() string {
	strBuilder := &strings.Builder{}

	strBuilder.WriteString(configContent)
	strBuilder.WriteString("\n")

	return strBuilder.String()
}
