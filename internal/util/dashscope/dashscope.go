package dashscope

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	xhttp "github.com/usual2970/meta-forge/internal/util/http"
)

type Dashscope struct {
	apiKey string
}

type DashscopeOption func(d *Dashscope)

func WithApiKey(apiKey string) DashscopeOption {
	return func(d *Dashscope) {
		d.apiKey = apiKey
	}
}

func New(opts ...DashscopeOption) *Dashscope {
	rs := &Dashscope{}
	for _, opt := range opts {
		opt(rs)
	}
	return rs

}

type MultimodalGenerationReq struct {
	Model      string                    `json:"model"`
	Input      MultimodalGenerationInput `json:"input"`
	Parameters map[string]interface{}    `json:"parameters"`
}

type MultimodalGenerationInput struct {
	Messages []MultimodalGenerationMessage `json:"messages"`
}

type MultimodalGenerationMessage struct {
	Role    string                               `json:"role"`
	Content []MultimodalGenerationMessageContent `json:"content"`
}

type MultimodalGenerationMessageContent struct {
	Text  string `json:"text,omitempty"`
	Audio string `json:"audio,omitempty"`
}

type MultimodalGenerationResp struct {
	Output    MultimodalGenerationOutput `json:"output"`
	Usage     MultimodalGenerationUsage  `json:"usage"`
	RequestID string                     `json:"request_id"`
	Code      string                     `json:"code"`
	Message   string                     `json:"message"`
}

type MultimodalGenerationOutput struct {
	Choices []MultimodalGenerationChoice `json:"choices"`
}

type MultimodalGenerationUsage struct {
	OutputTokens int `json:"output_tokens"`
	InputTokens  int `json:"input_tokens"`
	AudioTokens  int `json:"audio_tokens"`
}

type MultimodalGenerationChoice struct {
	FinishReason string                      `json:"finish_reason"`
	Message      MultimodalGenerationMessage `json:"message"`
}

func (d *Dashscope) MultimodalGeneration(req *MultimodalGenerationReq) (*MultimodalGenerationResp, error) {

	body, _ := json.Marshal(req)
	resp, err := xhttp.Req("https://dashscope.aliyuncs.com/api/v1/services/aigc/multimodal-generation/generation", http.MethodPost, bytes.NewBuffer(body), d.getHeader(), xhttp.WithTimeout(time.Minute))
	if err != nil {
		return nil, err
	}

	rs := &MultimodalGenerationResp{}
	if err := json.Unmarshal(resp, rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func (d *Dashscope) getHeader() map[string]string {

	return map[string]string{
		"Authorization": "Bearer " + d.apiKey,
		"Content-Type":  "application/json",
	}
}
