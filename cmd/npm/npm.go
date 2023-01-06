package npm

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

const (
	npmConfigFilename  = ".npmrc"
	yarnConfigFilename = ".yarnrc"
)

var (
	useGlobal bool
	RootCmd   = &cobra.Command{
		Use:   "npm",
		Short: "Create npm config",
	}
)

func init() {
	RootCmd.AddCommand(mirrorCmd)
	RootCmd.AddCommand(pnpmCmd)
}

func resolveConfigFilePath(useGlobal, useYarn bool) (string, error) {
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

	filename := npmConfigFilename
	if useYarn {
		filename = yarnConfigFilename
	}

	return path.Join(dir, filename), nil
}
