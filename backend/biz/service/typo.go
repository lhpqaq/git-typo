package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/go-git/go-git/v5"
)

// const proxyURL = "socks5://127.0.0.1:7890"

// cloneRepo 克隆指定的 Git 仓库到临时目录
func cloneRepo(gitUrl string) (string, error) {
	repoURL := gitUrl
	fmt.Println("repoURL: ", repoURL)
	dir := path.Join(os.TempDir(), "repo")

	// 删除临时目录（如果存在）
	err := os.RemoveAll(dir)
	if err != nil {
		return "", fmt.Errorf("failed to remove directory: %v", err)
	}

	// 创建临时目录
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// 克隆 Git 仓库
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:             repoURL,
		InsecureSkipTLS: true,
	})
	if err != nil {
		fmt.Println("clone repo error: ", err)
		return "", err
	}

	return dir, nil
}

// CheckTypo 检查 Git 仓库中的拼写错误
func CheckTypo(gitUrl string) (string, error) {
	dir, err := cloneRepo(gitUrl)
	if err != nil {
		return "", err
	}

	// 切换到克隆的目录
	err = os.Chdir(dir)
	if err != nil {
		return "", err
	}

	// // 执行 typos 命令
	cmd := exec.Command("typos") // 这里假设 typos 是可执行命令
	cmd.Env = append(cmd.Env, "TERM=xterm-256color")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // 捕获错误输出

	err = cmd.Run()
	if err == nil {
		return "Check Pass!", nil
	}

	// 获取输出，处理颜色转义字符
	output := out.String()
	fmt.Println(output)
	// 这里可以添加颜色转义字符的处理逻辑
	return output, nil
}

// func runTyposCommand() (string, error) {
// 	// 创建命令
// 	cmd := exec.Command("typos")

// 	// 启动 PTY 会话，连接命令的输入输出
// 	ptmx, err := pty.Start(cmd)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer ptmx.Close()

// 	// 读取命令输出
// 	buf := make([]byte, 1024)
// 	var output string
// 	for {
// 		n, err := ptmx.Read(buf)
// 		if err != nil {
// 			break
// 		}
// 		output += string(buf[:n])
// 	}
// 	output = strings.TrimSpace(output)
// 	output = strings.Replace(output, "\n", "", 0)                        // 去掉开头结尾的多余空白字符
// 	output = regexp.MustCompile(`\n{2,}`).ReplaceAllString(output, "\n") // 将多余的连续换行符替换为单个换行

// 	// 输出命令结果
// 	return output, nil
// }
