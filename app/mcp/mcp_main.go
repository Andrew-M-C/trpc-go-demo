package main

import (
	"github.com/Andrew-M-C/trpc-go-demo/app/mcp/logic/admindivision"
	"github.com/Andrew-M-C/trpc-go-utils/log"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	trpc "trpc.group/trpc-go/trpc-go"
	thttp "trpc.group/trpc-go/trpc-go/http"
	trpcserver "trpc.group/trpc-go/trpc-go/server"
)

func main() {
	svc := trpc.NewServer()
	mcpSvr := server.NewMCPServer(
		"ip-mcp",
		"1.0.0",
	)

	addAdminDivisionTool(mcpSvr)
	serveHTTP(svc, mcpSvr)
}

func addAdminDivisionTool(s *server.MCPServer) {
	h := admindivision.Handler{}
	options := []mcp.ToolOption{
		mcp.WithDescription(h.Description()),
	}
	options = append(options, h.Parameters()...)

	tool := mcp.NewTool("admin_division_query", options...)
	s.AddTool(tool, h.HandleMCP)
}

// serveHTTP 启动HTTP服务
func serveHTTP(svc *trpcserver.Server, mcpSvr *server.MCPServer) {
	// 创建SSE服务器
	// SSE端点会自动变为 /mcp/sse
	// 消息端点会自动变为 /mcp/message
	sseServer := server.NewSSEServer(mcpSvr, server.WithBasePath("/mcp"))

	wrappedHTTP := &wrappedHTTP{Handler: sseServer}
	thttp.HandleFunc("/mcp/sse", wrappedHTTP.ServeHTTP)
	thttp.HandleFunc("/mcp/message", wrappedHTTP.ServeHTTP)
	thttp.RegisterNoProtocolService(svc.Service("trpc.amc.demo.mcp"))

	if err := svc.Serve(); err != nil {
		log.Errorf("TRPC server error: '%v'", err)
	}
}
