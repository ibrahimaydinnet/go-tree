package internal

import (
	"fmt"
	"github.com/ibrahimaydinnet/go-tree/constant"
	"io"
	"math"
	"os"
	"strings"
)

func search(files []os.FileInfo, subStr string) []os.FileInfo {
	result := []os.FileInfo{}
	for _, f := range files {
		if f.IsDir() || strings.Contains(f.Name(), subStr) {
			result = append(result, f)
		}
	}
	return result
}

func justDirs(files []os.FileInfo) []os.FileInfo {
	dirs := []os.FileInfo{}
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f)
		}
	}
	return dirs
}

func appendSeperator(path string) string {
	if !strings.HasSuffix(path, constant.Seperator) {
		return fmt.Sprintf("%s%s", path, constant.Seperator)
	}
	return path
}

func formatSize(s int64) string {
	GB := int64(1024 * 1024 * 1024)
	MB := int64(1024 * 1024)
	KB := int64(1024)
	unit := "B"
	amount := 0.
	if s > GB {
		amount = math.Round(float64(s/GB)*100) / 100
		unit = "G"
	}
	if s > MB {
		amount = math.Round(float64(s/MB)*100) / 100
		unit = "M"
	}
	if s > KB {
		amount = math.Round(float64(s/KB)*100) / 100
		unit = "K"
	}
	return fmt.Sprintf("%v%v", amount, unit)
}

func exceptHiddens(all, win bool, path string, files []os.FileInfo) []os.FileInfo {
	result := []os.FileInfo{}
	for _, file := range files {
		if !all && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if !win && isHidden(path, file.Name()) {
			continue
		}
		result = append(result, file)
	}
	return result
}

func writeToFile(output, outputFile string) (err error) {
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	if _, err = io.WriteString(file, output); err != nil {
		return
	}
	if err = file.Sync(); err != nil {
		return
	}
	return
}

func colorize(color, str string) string {
	return fmt.Sprintf("%s%s%s", color, str, resetColor)
}
