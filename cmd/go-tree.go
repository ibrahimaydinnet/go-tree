package cmd

import (
	"fmt"
	"os"

	"github.com/ibrahimaydinnet/go-tree/constant"
	"github.com/ibrahimaydinnet/go-tree/internal"
	"github.com/spf13/cobra"
)

var flags map[string]interface{}

var goTree = &cobra.Command{
	Use:   "go-tree",
	Short: "go-tree is a cli tool which draws a tree of directory structure",
	Long:  `go-tree is a cli tool which draws a tree of directory structure`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.DrawTree(flags)
	},
}

func init() {
	flags = map[string]interface{}{}
	flags[constant.All] = goTree.PersistentFlags().BoolP(constant.All, "a", false, `Flag to list files starts with "."`)
	flags[constant.Color] = goTree.PersistentFlags().BoolP(constant.Color, "c", false, "Flag to colorize output")
	flags[constant.Date] = goTree.PersistentFlags().BoolP(constant.Date, "d", false, "Flag to show modified times")
	flags[constant.Emoji] = goTree.PersistentFlags().BoolP(constant.Emoji, "e", false, "Flag to show file and dir emojis (not supported on windows)")
	flags[constant.Justdir] = goTree.PersistentFlags().BoolP(constant.Justdir, "j", false, "Flag to list just directories")
	flags[constant.Number] = goTree.PersistentFlags().BoolP(constant.Number, "n", false, "Flag to show total numbers of files and directories")
	flags[constant.Mode] = goTree.PersistentFlags().BoolP(constant.Mode, "m", false, "Flag to show permission modes")
	flags[constant.Path] = goTree.PersistentFlags().BoolP(constant.Path, "p", false, "Flag to show fullpaths")
	flags[constant.Size] = goTree.PersistentFlags().BoolP(constant.Size, "s", false, "Flag to show formatted sizes")
	flags[constant.Trim] = goTree.PersistentFlags().BoolP(constant.Trim, "t", false, "Flag to trim empty directories")
	flags[constant.Verbose] = goTree.PersistentFlags().BoolP(constant.Verbose, "v", false, "Flag to show not formatted sizes")
	flags[constant.Win] = goTree.PersistentFlags().BoolP(constant.Win, "w", false, "Flag to list hidden files for windows")
	flags[constant.Find] = goTree.PersistentFlags().StringP(constant.Find, "f", "", "Substring of file names to find")
	flags[constant.Output] = goTree.PersistentFlags().StringP(constant.Output, "o", "", "File path to write output")
	flags[constant.Root] = goTree.PersistentFlags().StringP(constant.Root, "r", fmt.Sprintf("%s%s", ".", constant.Seperator), "Root path of the tree")
	flags[constant.Level] = goTree.PersistentFlags().IntP(constant.Level, "l", 0, "Max level of tree depth")
}

// Execute .. executes cli commands
func Execute() {
	if err := goTree.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
