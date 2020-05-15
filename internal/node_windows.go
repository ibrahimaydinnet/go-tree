// +build windows

package internal

import (
	"fmt"
	"path/filepath"
	"runtime"
	"syscall"
)

func isHidden(path, fileName string) bool {
	if runtime.GOOS == "windows" {
		pointer, err := syscall.UTF16PtrFromString(filepath.Join(path, fileName))
		if err != nil {
			fmt.Println(err)
		}
		attributes, err := syscall.GetFileAttributes(pointer)
		if err != nil {
			fmt.Println(err)
		}
		return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
	}
	return false
}
