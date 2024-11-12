package main

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServerMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 정적 분석 오류를 회피하기 위해 명시적으로 반환값을 버림
		_, _ = w.Write([]bute(`{"status": "ok"}`))
	})
	return mux
}