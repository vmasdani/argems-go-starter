package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

// Brotli Compression
type brotliResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w brotliResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func BrotliCompressorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "br")
		br := brotli.NewWriter(w)
		defer br.Close()
		brw := brotliResponseWriter{Writer: br, ResponseWriter: w}
		next.ServeHTTP(brw, r)
	})
}
