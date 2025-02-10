# eino-agent

一个基于 [Eino](https://github.com/cloudwego/eino) 简易 agent 实现，能够完成公司员工信息增删改查，支持 Kimi 和 Deepseek 等各种兼容 OpenAI 的大模型.

## 运行程序

1. 添加大语言模型相关环境变量

修改 env.sh.example，修改后记得 source

2. 运行程序

```
go run main.go
```


## 使用体验

```
欢迎使用员工信息 Agent, 支持用户信息的增删改查，输入 'exit' 退出程序。

请输入操作: 请查询 songjiayang 信息
2025/02/10 21:22:36 姓名: songjiayang, 年龄: 34, 部门: 软件研发

请输入操作: 更新songjiayang 信息，年龄35，部门测试
2025/02/10 21:22:55 songjiayang updated, 年龄: 35, 部门: 测试

请输入操作: 请查询 songjiayang 信息
2025/02/10 21:26:04 姓名: songjiayang, 年龄: 35, 部门: 测试

请输入操作: 添加用户，姓名小红，年龄40，部门：HR
2025/02/10 21:23:29 小红 added, 年龄: 40, 部门: HR

请输入操作: 删除 songjiayang 用户
2025/02/10 21:23:43 songjiayang deleted

请输入操作: 请查询 songjiayang 信息
2025/02/10 21:25:04 songjiayang not found

请输入操作:   今天天气如何?
暂不支持该操作
```
