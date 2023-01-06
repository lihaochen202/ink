package npm

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	useYarn       bool
	mirrorConfigs = [...][2]string{
		{"registry", "https://registry.npmmirror.com/"},
		{"disturl", "https://npmmirror.com/mirrors/node/"},
		{"sass_binary_site", "https://npmmirror.com/mirrors/node-sass/"},
		{"electron_mirror", "https://npmmirror.com/mirrors/electron/"},
		{"python_mirror", "https://npmmirror.com/mirrors/python/"},
	}
	mirrorCmd = &cobra.Command{
		Use:   "mirror",
		Short: "Create npm mirror config",
		RunE:  execMirrorCmd,
	}
)

func init() {
	bindMirrorCmdFlags()
}

func bindMirrorCmdFlags() {
	mirrorCmd.Flags().BoolVarP(&useGlobal, "global", "g", false, "apply to the global config")
	mirrorCmd.Flags().BoolVarP(&useYarn, "yarn", "y", false, "whether to use yarn")
}

func execMirrorCmd(cmd *cobra.Command, args []string) error {
	filePath, err := resolveConfigFilePath(useGlobal, useYarn)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	content := genMirrorConfigContent(mirrorConfigs, useYarn)
	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}

func genMirrorConfigContent(configs [5][2]string, useYarn bool) string {
	strBuilder := &strings.Builder{}

	split := "="
	if useYarn {
		split = " "
	}

	for _, config := range configs {
		key, value := config[0], config[1]
		strBuilder.WriteString(key)
		strBuilder.WriteString(split)
		strBuilder.WriteString(value)
		strBuilder.WriteString("\n")
	}

	return strBuilder.String()
}
