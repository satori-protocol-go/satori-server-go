package application

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori-protocol-go/satori-server-go/pkg/config"
)

type Instance interface {

	// Start 开启服务
	Start() error
}

type ginInstance struct {
	ctx context.Context
	cfg config.Cfg
}

func (instacne *ginInstance) Start() error {
	errChan := make(chan error)
	router := gin.Default()
	if instacne.cfg.Token() != "" {
		router.Use(func(c *gin.Context) {
			if c.GetHeader("Authorization") != "Bearer "+instacne.cfg.Token() {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
		})
	}
	svc := &http.Server{
		Addr:    instacne.cfg.Addr(),
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
	case <-instacne.ctx.Done():
		return svc.Shutdown(instacne.ctx)
	case err := <-errChan:
		return err
	}
}

func NewInstance(ctx context.Context, conf config.Cfg) (Instance, error) {
	return &ginInstance{
		ctx: ctx,
		cfg: conf,
	}, nil
}
