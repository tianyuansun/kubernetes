package filters

import (
	"fmt"
	"net/http"
)

const maxReqPerConn = 10

func WithMaxReqPerConn(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("suntianyuan: req headers are %v\n", r.Header)
		localAddr := r.Context().Value(http.LocalAddrContextKey)
		if ct, ok := localAddr.(interface{ Increment() int }); ok {
			if ct.Increment() >= maxReqPerConn {
				fmt.Printf("suntianyuan: Connection from %s to %+v should close\n", r.RemoteAddr, localAddr)
				//coon, _, _ := c.Writer.Hijack()
				//_ = coon.Close()
				//return
				w.Header().Set("Connection", "close")
			} else {
				fmt.Printf("Connection from %s to %+v should Keepalive\n", r.RemoteAddr, localAddr)
				handler.ServeHTTP(w, r)
			}
		}
	})
}