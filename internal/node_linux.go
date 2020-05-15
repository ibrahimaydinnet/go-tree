// +build !windows

package internal

func isHidden(path, fileName string) bool {
	return false
}
