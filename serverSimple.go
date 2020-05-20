package socketio

import (
	"net/http"

	engineio "github.com/talkfun/go-engine.io"
)

// 简化版的 socketio，主要不需要各事件接口，直接暴露相关的连接出去
type ServerSimple struct {
	Eio		*engineio.Server
}

// NewServerSimple return a ServerSimple.
func NewServerSimple(c *engineio.Options) (*ServerSimple, error) {
	eio, err := engineio.NewServer(c)
	if err != nil {
		return nil, err
	}
	return &ServerSimple{
		Eio: eio,
	}, nil
}

// Close closes server.
func (s *ServerSimple)Close() error {
	return s.Eio.Close()
}

func (s *ServerSimple)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Eio.ServeHTTP(w, r)
}
