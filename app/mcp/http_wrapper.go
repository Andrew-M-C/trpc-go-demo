package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/Andrew-M-C/go-bytesize"
	"trpc.group/trpc-go/trpc-go/log"
)

type wrappedHTTP struct {
	http.Handler
}

func (wh *wrappedHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Infof("ServeHTTP: %s", r.URL)
	log.Infof("ServeHTTP: Method: %s", r.Method)
	log.Infof("ServeHTTP: Headers: %v", r.Header)
	log.Infof("ServeHTTP: Content-Length: %v", r.ContentLength)

	// 设置 CORS 头部允许所有来源
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 如果是 OPTIONS 请求，直接返回成功
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	in, _ := io.ReadAll(r.Body)
	if len(in) > 0 {
		log.Infof("read %v:\n%s", bytesize.Base10(len(in)), in)
	} else {
		log.Info("read nothing")
	}

	r.Body = io.NopCloser(bytes.NewBuffer(in))
	writerWrapper := &writerWrapper{w: w}
	wh.Handler.ServeHTTP(writerWrapper, r)
	return nil
}

// 监控返回值
type writerWrapper struct {
	w http.ResponseWriter
}

func (w *writerWrapper) Write(p []byte) (int, error) {
	log.Infof("write: \n%s", p)
	return w.w.Write(p)
}

func (w *writerWrapper) WriteHeader(statusCode int) {
	w.w.WriteHeader(statusCode)
}

func (w *writerWrapper) Header() http.Header {
	return w.w.Header()
}

func (w *writerWrapper) Flush() {
	flusher, _ := w.w.(http.Flusher)
	flusher.Flush()
}
