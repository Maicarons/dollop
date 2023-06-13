package check

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Go bool is checked result, string is version, error is error message.
func Go() (bool, string, error) {

	// 检查Go版本
	versionOutput, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		return false, "", fmt.Errorf(localize("检查Go版本失败：%v", "Failed to check Go version: %v", err.Error()))
	}
	version := strings.Split(string(versionOutput), " ")[2]

	// 检查GOPATH
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return false, version, fmt.Errorf(localize("GOPATH未设置", "GOPATH is not set"))
	}

	// 检查GOROOT
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		return false, version, fmt.Errorf(localize("GOROOT未设置", "GOROOT is not set"))
	}

	// 检查GOPATH是否存在
	if _, err := os.Stat(gopath); os.IsNotExist(err) {
		return false, version, fmt.Errorf(localize("GOPATH不存在：%v", "GOPATH does not exist: %v", err.Error()))
	}

	// 检查GOROOT是否存在
	if _, err := os.Stat(goroot); os.IsNotExist(err) {
		return false, version, fmt.Errorf(localize("GOROOT不存在：%v", "GOROOT does not exist: %v", err.Error()))
	}

	// 检查GOPATH/bin是否在PATH中
	pathEnv := os.Getenv("PATH")
	gopathBin := filepath.Join(gopath, "bin")
	if !strings.Contains(pathEnv, gopathBin) {
		switch runtime.GOOS {
		case "windows":
			// 在Windows上尝试将GOPATH/bin添加到PATH中
			err := exec.Command("setx", "PATH", fmt.Sprintf(`%s;%s`, pathEnv, gopathBin)).Run()
			if err != nil {
				return false, version, fmt.Errorf(localize("无法将GOPATH/bin添加到PATH中：%v", "Failed to add GOPATH/bin to PATH: %v", err.Error()))
			}
		case "linux", "darwin":
			// 在Linux和MacOS上尝试将GOPATH/bin添加到PATH中
			shell := os.Getenv("SHELL")
			if shell == "" {
				shell = "/bin/bash"
			}
			err := exec.Command(shell, "-c", fmt.Sprintf(`echo 'export PATH=$PATH:%s' >> ~/.bashrc`, gopathBin)).Run()
			if err != nil {
				return false, version, fmt.Errorf(localize("无法将GOPATH/bin添加到PATH中：%v", "Failed to add GOPATH/bin to PATH: %v", err.Error()))
			}
		default:
			return false, version, fmt.Errorf(localize("不支持的操作系统：%s", "Unsupported operating system: %s", runtime.GOOS))
		}
	}
	return true, version, nil
}

// localize函数将错误消息本地化为当前语言环境
func localize(en string, otherLangs ...string) string {
	lang := os.Getenv("LANG")
	switch {
	case strings.HasPrefix(lang, "zh"):
		return otherLangs[0]
	default:
		return en
	}
}
