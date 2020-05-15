package internal

import (
	"bytes"
	"fmt"
	"os"
	"runtime"

	"github.com/ibrahimaydinnet/go-tree/constant"

	colorable "github.com/mattn/go-colorable"
)

const (
	fileEmoji  = "\U0001f4c4"
	dirEmoji   = "\U0001f4c1"
	file       = "file"
	dir        = "dir"
	resetColor = "\x1b[0m"
)

var lineColors = []string{
	"\x1b[31m", // red
	"\x1b[32m", // green
	"\x1b[33m", // yellow
}

var nameColors = map[string]string{
	file: "\x1b[35m", // magenta
	dir:  "\x1b[36m", // cyan
}

// Tree ..
type tree struct {
	root  node
	flags map[string]interface{}
}

// DrawTree .. draws a tree map
func DrawTree(flags map[string]interface{}) {
	if hasEmoji := *(flags[constant.Emoji].(*bool)); runtime.GOOS == "windows" && hasEmoji {
		fmt.Println("no emoji support on windows!")
		os.Exit(1)
	}
	rootPath := *(flags[constant.Root].(*string))
	info, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		fmt.Println("no such directory:", rootPath)
		os.Exit(1)
	}
	tree := tree{node{nil, nil, 0, false, rootPath, map[string]int{dir: 0, file: 0}, info}, flags}
	tree.draw()

	os.Exit(0)
}

func (t tree) draw() {
	if err := t.root.buildTree(t.flags); err != nil {
		fmt.Println(err)
	}

	if outputFile := *(t.flags[constant.Output].(*string)); outputFile != "" {
		buf := new(bytes.Buffer)
		t.root.draw(buf, t.flags)
		if err := writeToFile(buf.String(), outputFile); err != nil {
			fmt.Println(err)
		}
	} else {
		t.root.draw(colorable.NewColorableStdout(), t.flags)
	}

	if hasNumber := *(t.flags[constant.Number].(*bool)); hasNumber {
		fmt.Printf("\ntotal directories: %v, total files: %v\n", t.root.total[dir], t.root.total[file])
	}
}
