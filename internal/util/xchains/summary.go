package xchains

import (
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
)

const _stuffSummarizationTemplate = `对下面的内容做一个简要的总结:


"{{.context}}"


总结:`

func LoadStuffSummarization(llm llms.LLM) chains.StuffDocuments {
	llmChain := chains.NewLLMChain(llm, prompts.NewPromptTemplate(
		_stuffSummarizationTemplate, []string{"context"},
	))

	return chains.NewStuffDocuments(llmChain)
}
