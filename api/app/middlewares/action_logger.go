package middlewares

import "github.com/go-raptor/raptor"

type ActionLogger struct {
	raptor.Middleware
}

func (m *ActionLogger) New(c *raptor.Context) error {
	return c.Next()
}
