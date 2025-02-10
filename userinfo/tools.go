package userinfo

import (
	"context"
	"log"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

func UserInfoTools(ctx context.Context) ([]tool.BaseTool, []*schema.ToolInfo) {
	addTool, _ := utils.InferTool("add_employee", "添加用户信息，包括 username, age, department 字段", AddFunc)
	findTool, _ := utils.InferTool("find_employee", "查看用户信息", FindFunc)
	updateTool, _ := utils.InferTool("update_employee", "更新用户信息，包括 age, department", UpdateFunc)
	deleteTool, _ := utils.InferTool("delete_employee", "删除用户信息", DeleteFunc)

	tools := []tool.BaseTool{
		addTool,
		findTool,
		updateTool,
		deleteTool,
	}

	var toolInfos []*schema.ToolInfo
	for _, tool := range tools {
		info, err := tool.Info(ctx)
		if err != nil {
			log.Fatalf("get ToolInfo failed, err=%v", err)
		}
		toolInfos = append(toolInfos, info)
	}

	return tools, toolInfos
}

type UserInfo struct {
	Username   string `json:"username" jsonschema:"description=username of the employee"`
	Age        int    `json:"age" jsonschema:"description=age of the employee"`
	Department string `json:"department" jsonschema:"description=department of the employee"`
}

var (
	userDb = map[string]UserInfo{
		"songjiayang": {
			Username:   "songjiayang",
			Age:        34,
			Department: "软件研发",
		},
	}
)

func FindFunc(_ context.Context, params *UserInfo) (resp string, err error) {
	info, ok := userDb[params.Username]
	if ok {
		log.Printf("姓名: %s, 年龄: %d, 部门: %s", info.Username, info.Age, info.Department)
	} else {
		log.Printf("%s not found", params.Username)
	}
	return
}

func AddFunc(_ context.Context, params *UserInfo) (resp string, err error) {
	userDb[params.Username] = *params
	log.Printf("%s added, 年龄: %d, 部门: %s", params.Username, params.Age, params.Department)

	return
}

func UpdateFunc(_ context.Context, params *UserInfo) (resp string, err error) {
	if _, ok := userDb[params.Username]; !ok {
		log.Printf("%s not found", params.Username)
		return
	}

	userDb[params.Username] = *params
	log.Printf("%s updated, 年龄: %d, 部门: %s", params.Username, params.Age, params.Department)

	return
}

func DeleteFunc(_ context.Context, params *UserInfo) (resp string, err error) {
	if _, ok := userDb[params.Username]; !ok {
		log.Printf("%s not found", params.Username)
		return
	}

	delete(userDb, params.Username)
	log.Printf("%s deleted", params.Username)

	return
}
