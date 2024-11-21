// -------------------------------------------
// @file      : root_path.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/21 下午5:24
// -------------------------------------------

package misc

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultRootPath = "./"
	RootLocater     = ".root_locater"
)

// IsFileExist 文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Lstat(filename)
	return !os.IsNotExist(err)
}

// IsRootDir 判断是否根目录
// 使用这个函数的前提是在项目根目录下有一个.root_locater文件
func IsRootDir(path string) bool {
	locaterPath := filepath.Join(path, RootLocater)
	if IsFileExist(locaterPath) {
		return true
	}
	return false
}

// GetRootPath 获取根目录
// 使用这个函数的前提是在项目根目录下有一个.root_locater文件
func GetRootPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		return DefaultRootPath
	}
	path, err := filepath.Abs(cwd)
	if err != nil {
		return DefaultRootPath
	}
	for depth := 32; depth > 0 && !IsRootDir(path); depth-- {
		idx := strings.LastIndexByte(path, filepath.Separator)
		if idx >= 0 {
			path = path[:idx]
		} else {
			return DefaultRootPath
		}
	}
	return path
}

// GetPathInRootDir 获取根目录下的文件路径
// 使用这个函数的前提是在项目根目录下有一个.root_locater文件
func GetPathInRootDir(filename string) string {
	rootDir := GetRootPath()
	return filepath.Join(rootDir, filename)
}
