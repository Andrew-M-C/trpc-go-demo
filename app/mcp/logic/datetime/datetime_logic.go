package datetime

import (
	"context"
	"time"

	jsutil "github.com/Andrew-M-C/go.util/encoding/json"
	tmutil "github.com/Andrew-M-C/go.util/time"
	"github.com/mark3labs/mcp-go/mcp"
)

type Handler struct{}

const description = `
返回当前的时间，以 JSON 格式输出，包含以下字段：

- utc: 格式为 "YYYY-MM-DD HH:MM:SS" 的当前 UTC 时间
- beijing: 格式为 "YYYY-MM-DD HH:MM:SS" 的当前北京时间
- timestamp_sec: 当前时间戳，单位为秒
`

func (Handler) Description() string {
	return description
}

func (Handler) Parameters() []mcp.ToolOption {
	// 无需入参
	return []mcp.ToolOption{}
}

func (Handler) HandleMCP(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 获取当前时间并格式化
	now := time.Now()

	// 构建响应
	res := response{
		UTC:       now.UTC().Format(time.DateTime),
		Beijing:   now.In(tmutil.Beijing).Format(time.DateTime),
		Timestamp: now.Unix(),
	}

	// 转换为 JSON 字符串
	s, _ := jsutil.MarshalToString(res)
	return mcp.NewToolResultText(s), nil
}

type response struct {
	UTC       string `json:"utc"`
	Beijing   string `json:"beijing"`
	Timestamp int64  `json:"timestamp_sec"`
}
