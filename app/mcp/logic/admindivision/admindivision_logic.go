package admindivision

import (
	"context"
	"errors"

	add "github.com/Andrew-M-C/go.util/china/admindivision"
	jsutil "github.com/Andrew-M-C/go.util/encoding/json"
	"github.com/mark3labs/mcp-go/mcp"
)

type Handler struct{}

const description = `
根据行政区划代码获取行政区划名称, 返回 JSON 格式, 包含以下字段:

- province: 规范化的省级行政区名称
- city: 规范化的市级行政区名称
- county: 规范化的区县级行政区名称
- province_code: 省级行政区代码, 如广东省为 44
- city_code: 市级行政区代码, 如广州市为 01
- county_code: 区县级行政区代码, 如越秀区为 04
`

func (Handler) Description() string {
	return description
}

func (Handler) Parameters() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithString("province",
			mcp.Required(),
			mcp.Description("省级行政区名称"),
		),
		mcp.WithString("city",
			mcp.Description("市级行政区名称"),
		),
		mcp.WithString("county",
			// 非必要
			mcp.Description("县级行政区名称"),
		),
	}
}

func (Handler) HandleMCP(_ context.Context, r mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	nameChain, err := readParams(r)
	if err != nil {
		return nil, err
	}
	adminChain := add.SearchDivisionByName(nameChain...)
	if len(adminChain) < 1 {
		return nil, errors.New("无法找到行政区划代码")
	}

	res := response{
		Province:     adminChain[0].Name(),
		ProvinceCode: adminChain[0].Code(),
	}
	if len(adminChain) > 1 {
		res.City = adminChain[1].Name()
		res.CityCode = adminChain[1].Code()
	}
	if len(adminChain) > 2 {
		res.County = adminChain[2].Name()
		res.CountyCode = adminChain[2].Code()
	}

	s, _ := jsutil.MarshalToString(res)
	return mcp.NewToolResultText(s), nil
}

type response struct {
	Province     string `json:"province"`
	ProvinceCode string `json:"province_code"`
	City         string `json:"city"`
	CityCode     string `json:"city_code"`
	County       string `json:"county,omitempty"`
	CountyCode   string `json:"county_code,omitempty"`
}

func readParams(r mcp.CallToolRequest) ([]string, error) {
	var res []string
	province, _ := r.Params.Arguments["province"].(string)
	if province == "" {
		return nil, errors.New("province 参数缺失")
	}
	res = append(res, province)
	city, _ := r.Params.Arguments["city"].(string)
	if city == "" {
		return res, nil
	}
	res = append(res, city)
	county, _ := r.Params.Arguments["county"].(string)
	if county == "" {
		return res, nil
	}
	return append(res, county), nil
}
