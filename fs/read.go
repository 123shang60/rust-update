package fs

import (
	"os"
	"path/filepath"
)

// GetNowPath 获取命令执行路径
func getNowPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// GetCargoToml 获取cargo.toml文件内容
func GetCargoToml() (string, error) {
	path, err := getNowPath()
	if err != nil {
		return "", err
	}
	readFile, err := os.ReadFile(path + "/Cargo.toml")
	if err != nil {
		return "", err
	}
	return string(readFile), nil
}
