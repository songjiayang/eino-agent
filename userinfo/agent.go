package userinfo

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type Agent struct {
	runnable compose.Runnable[[]*schema.Message, []*schema.Message]
}

func NewAgent() *Agent {
	ctx := context.Background()
	chatModel, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
		BaseURL: os.Getenv("OPENAI_API_URL"),
		Model:   os.Getenv("MODEL_ID"),
		APIKey:  os.Getenv("OPENAI_API_KEY"),
	})
	if err != nil {
		log.Fatalf("NewChatModel failed, err=%v", err)
		return nil
	}

	tools, toolInfos := UserInfoTools(ctx)
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		log.Fatalf("BindTools failed, err=%v", err)
		return nil
	}

	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: tools,
	})
	if err != nil {
		log.Fatalf("NewToolNode failed, err=%v", err)
		return nil
	}

	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	agent, err := chain.Compile(ctx)
	if err != nil {
		log.Fatalf("chain.Compile failed, err=%v", err)
		return nil
	}

	return &Agent{
		runnable: agent,
	}
}

func (a *Agent) Invoke(ctx context.Context, content string) {
	_, err := a.runnable.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: content,
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "no tool call found in input") {
			fmt.Println("暂不支持该操作")
			return
		}
		log.Printf("agent.Invoke failed, err=%v", err)
	}
}
