package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/songjiayang/eino-agent/userinfo"
)

func main() {
	agent := userinfo.NewAgent()
	ctx := context.Background()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("欢迎使用员工信息 Agent, 支持用户信息的增删改查，输入 'exit' 退出程序。")
	inputTips := "\n请输入操作: "
	for {
		fmt.Print(inputTips)
		if !scanner.Scan() {
			fmt.Println("读取输入失败，程序退出。")
			break
		}

		input := scanner.Text()

		switch strings.ToLower(input) {
		case "exit":
			fmt.Println("欢迎再次使用，再见。")
			return
		default:
			agent.Invoke(ctx, strings.Replace(input, inputTips, "", 1))
		}
	}
}
