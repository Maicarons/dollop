package check

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"os/exec"
	"strings"
)

func Ebitengine() (bool, string, error) {
	// 检查Go版本
	goVersionOutput, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		return false, "", fmt.Errorf("检查Go版本失败：%v", err)
	}
	goVersion := strings.Split(string(goVersionOutput), " ")[2]
	// 检查最小Go版本
	minGoVersion, err := semver.NewVersion("1.16.0")
	if err != nil {
		return false, "", fmt.Errorf("解析最小Go版本失败：%v", err)
	}
	currentGoVersion, err := semver.NewVersion(goVersion[2:])
	if err != nil {
		return false, "", fmt.Errorf("解析当前Go版本失败：%v", err)
	}
	if currentGoVersion.LessThan(*minGoVersion) {
		return false, "", fmt.Errorf("当前Go版本过低，最小版本要求：%s", minGoVersion.String())
	}

	// 检查EbitEngine版本
	ebitEngineVersionOutput, err := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "github.com/hajimehoshi/ebiten/v2").CombinedOutput()
	if err != nil {
		return false, "", fmt.Errorf("检查EbitEngine版本失败：%v", err)
	}
	ebitEngineVersion := strings.TrimSpace(string(ebitEngineVersionOutput))

	return true, ebitEngineVersion, nil
}
