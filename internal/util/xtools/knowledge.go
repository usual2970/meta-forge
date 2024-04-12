package xtools

import (
	"context"
	"regexp"
	"strings"

	"github.com/usual2970/meta-forge/internal/util/zhipu"

	"github.com/tmc/langchaingo/callbacks"
)

type Knowledge struct {
	apiKey           string
	CallbacksHandler callbacks.Handler
}

func NewKnowledge(apiKey string) *Knowledge {
	return &Knowledge{
		apiKey: apiKey,
	}
}

func (c Knowledge) Name() string {
	return "LocalKnowledge"
}

func (c Knowledge) Description() string {
	return `"业务系统相关的本地知识库"
			"当你要回答问题时，优先从本工具中获取答案"
			"当你自身的知识库中没有相关的答案时，或其它工具无法获取答案时，也要使用本工具尝试获取答案"
			"输入参数就是要回答的问题"`
}

func (c Knowledge) Call(ctx context.Context, input string) (string, error) {
	k := zhipu.NewKnowledge(c.apiKey)

	resp, err := k.Invoke(input)
	if err != nil {
		return "", err
	}

	// 创建正则表达式匹配 Markdown 链接
	linkPattern := regexp.MustCompile(`!\[[^\]]*\]\((http[s]?://[^\)]*)\)|\[[^\]]*\]\((http[s]?://[^\)]*)\)`) // 匹配 ![alt text](url) 和 [text](url)

	// 使用 ReplaceAllStringFunc 提取并返回 URL
	result := linkPattern.ReplaceAllStringFunc(resp.Data.Content, extractURL)

	return result, nil
}

func extractURL(match string) string {
	// Markdown 链接格式是 [text](url)，我们需要提取括号内的 URL
	parts := strings.Split(match, "(")
	if len(parts) > 1 {
		urlPart := parts[1]
		// 去掉结尾的 ')'
		url := urlPart[:len(urlPart)-1]
		return url
	}
	// 如果不是有效的 Markdown 链接格式，返回原字符串
	return match
}
