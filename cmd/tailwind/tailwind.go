package tailwind

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

const (
	configFilename = "tailwind.config.cjs"
	configContent  = `/* eslint-env node */
/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
	theme: {
		columns: {
			'auto': 'auto',
			'1': '1',
			'2': '2',
			'3': '3',
			'4': '4',
			'5': '5',
			'6': '6',
			'7': '7',
			'8': '8',
			'9': '9',
			'10': '10',
			'11': '11',
			'12': '12',
			'3xs': '256px', // 16rem
			'2xs': '288px', // 18rem
			'xs': '320px', // 20rem
			'sm': '384px', // 24rem
			'md': '448px', // 28rem
			'lg': '512px', // 32rem
			'xl': '576px', // 36rem
			'2xl': '672px', // 42rem
			'3xl': '768px', // 48rem
			'4xl': '896px', // 56rem
			'5xl': '1024px', // 64rem
			'6xl': '1152px', // 72rem
			'7xl': '1280px', // 80rem
		},
		spacing: {
			px: '1px',
			0: '0px',
			0.5: '2px', // 0.125rem
			1: '4px', // 0.25rem
			1.5: '6px', // 0.375rem
			2: '8px', // 0.5rem
			2.5: '10px', // 0.625rem
			3: '12px', // 0.75rem
			3.5: '14px', // 0.875rem
			4: '16px', // 1rem
			5: '20px', // 1.25rem
			6: '24px', // 1.5rem
			7: '28px', // 1.75rem
			8: '32px', // 2rem
			9: '36px', // 2.25rem
			10: '40px', // 2.5rem
			11: '44px', // 2.75rem
			12: '48px', // 3rem
			14: '56px', // 3.5rem
			16: '64px', // 4rem
			20: '80px', // 5rem
			24: '96px', // 6rem
			28: '112px', // 7rem
			32: '128px', // 8rem
			36: '144px', // 9rem
			40: '160px', // 10rem
			44: '176px', // 11rem
			48: '192px', // 12rem
			52: '208px', // 13rem
			56: '224px', // 14rem
			60: '240px', // 15rem
			64: '256px', // 16rem
			72: '288px', // 18rem
			80: '320px', // 20rem
			96: '384px', // 24rem
		},
		borderRadius: {
			'none': '0px',
			'sm': '2px', // 0.125rem
			'DEFAULT': '4px', // 0.25rem
			'md': '6px', // 0.375rem
			'lg': '8px', // 0.5rem
			'xl': '12px', // 0.75rem
			'2xl': '16px', // 1rem
			'3xl': '24px', // 1.5rem
			'full': '9999px',
		},
		fontSize: {
			'xs': ['12px', { lineHeight: '16px' }], // 0.75rem 1rem
			'sm': ['14px', { lineHeight: '20px' }], // 0.875rem 1.25rem
			'base': ['16px', { lineHeight: '24px' }], // 1rem 1.5rem
			'lg': ['18px', { lineHeight: '28px' }], // 1.125rem 1.75rem
			'xl': ['20px', { lineHeight: '28px' }], // 1.25rem 1.75rem
			'2xl': ['24px', { lineHeight: '32px' }], // 1.5rem 2rem
			'3xl': ['30px', { lineHeight: '36px' }], // 1.875rem 2.25rem
			'4xl': ['36px', { lineHeight: '40px' }], // 2.25rem 2.5rem
			'5xl': ['48px', { lineHeight: '1' }], // 3rem
			'6xl': ['60px', { lineHeight: '1' }], // 3.75rem
			'7xl': ['72px', { lineHeight: '1' }], // 4.5rem
			'8xl': ['96px', { lineHeight: '1' }], // 6rem
			'9xl': ['128px', { lineHeight: '1' }], // 8rem
		},
		lineHeight: {
			none: '1',
			tight: '1.25',
			snug: '1.375',
			normal: '1.5',
			relaxed: '1.625',
			loose: '2',
			3: '12px', // .75rem
			4: '16px', // 1rem
			5: '20px', // 1.25rem
			6: '24px', // 1.5rem
			7: '28px', // 1.75rem
			8: '32px', // 2rem
			9: '36px', // 2.25rem
			10: '40px', // 2.5rem
		},
		maxWidth: ({ theme, breakpoints }) => ({
			'none': 'none',
			'0': '0px', // 0rem
			'xs': '320px', // 20rem
			'sm': '384px', // 24rem
			'md': '448px', // 28rem
			'lg': '512px', // 32rem
			'xl': '576px', // 36rem
			'2xl': '672px', // 42rem
			'3xl': '768px', // 48rem
			'4xl': '896px', // 56rem
			'5xl': '1024px', // 64rem
			'6xl': '1152px', // 72rem
			'7xl': '1280px', // 80rem
			'full': '100%',
			'min': 'min-content',
			'max': 'max-content',
			'fit': 'fit-content',
			'prose': '65ch',
			...breakpoints(theme('screens')),
		}),
	},
	plugins: [],
}`
)

var (
	RootCmd = &cobra.Command{
		Use:   "tailwind",
		Short: "Create tailwind config",
		RunE:  execCmd,
	}
)

func execCmd(cmd *cobra.Command, args []string) error {
	filePath, err := resolveConfigFilePath(configFilename)
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

func resolveConfigFilePath(filename string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(dir, filename), nil
}

func genConfigContent() string {
	strBuilder := &strings.Builder{}

	strBuilder.WriteString(configContent)
	strBuilder.WriteString("\n")

	return strBuilder.String()
}
