package npm

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	useLoose    bool
	looseConfig = "strict-peer-dependencies=false"
	useFlat     bool
	flatConfig  = "shamefully-hoist=true"
	pnpmCmd     = &cobra.Command{
		Use:   "pnpm",
		Short: "Create pnpm config",
		RunE:  execPnpmCmd,
	}
)

func init() {
	bindPnpmCmdFlags()
}

func bindPnpmCmdFlags() {
	pnpmCmd.Flags().BoolVarP(&useGlobal, "global", "g", false, "apply to the global config")
	pnpmCmd.Flags().BoolVarP(&useLoose, "loose", "l", false, "cancel peer dependencies detection")
	pnpmCmd.Flags().BoolVarP(&useFlat, "flat", "f", false, "create flat dependencies")
}

func execPnpmCmd(cmd *cobra.Command, args []string) error {
	content := genPnpmConfigContent(useLoose, looseConfig, useFlat, flatConfig)
	if len(content) == 0 {
		return nil
	}

	useYarn = false
	filePath, err := resolveConfigFilePath(useGlobal, useYarn)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}

func genPnpmConfigContent(useLoose bool, looseConfig string, useFlat bool, flatConfig string) string {
	strBuilder := &strings.Builder{}

	if useLoose {
		strBuilder.WriteString(looseConfig)
		strBuilder.WriteString("\n")
	}

	if useFlat {
		strBuilder.WriteString(flatConfig)
		strBuilder.WriteString("\n")
	}

	return strBuilder.String()
}
