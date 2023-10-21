package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config interface {
	Addr() string
	Token() string
}

type simpleConfig struct {
	AddrString  string `yaml:"addr" json:"addr"`
	TokenString string `yaml:"token,omitempty" json:"token,omitempty"`
}

func (c *simpleConfig) Addr() string {
	if c.AddrString == "" {
		return ":8080"
	}
	return c.AddrString
}

func (c *simpleConfig) Token() string {
	return c.TokenString
}

func NewConfig(Addr, Token string) Config {
	return &simpleConfig{AddrString: Addr, TokenString: Token}
}

// Start 开启服务
func Start(ctx context.Context, conf Config) error {
	errChan := make(chan error)
	router := gin.Default()
	if conf.Token() != "" {
		router.Use(func(c *gin.Context) {
			if c.GetHeader("Authorization") != "Bearer "+conf.Token() {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
		})
	}
	svc := &http.Server{
		Addr:    conf.Addr(),
		Handler: router,
	}
	// 异常通过chan 传递
	go func() {
		if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// panic(err)
			errChan <- err
		}
	}()
	select {
	case <-ctx.Done():
		return svc.Shutdown(ctx)
	case err := <-errChan:
		return err
	}
}
