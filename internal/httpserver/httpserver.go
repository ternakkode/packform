package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ternakkode/packform-backend/internal/httpserver/middleware"
	"github.com/ternakkode/packform-backend/internal/httpserver/router"
)

type HttpServer struct {
	handler    *gin.Engine
	httpServer *http.Server
}

func InitAndStart(
	appName string, httpServerPort string,
	additionalMiddlewares ...gin.HandlerFunc,
) *HttpServer {

	switch os.Getenv("ENV") {
	case "PRODUCTION", "PROD":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	ginHandler := gin.Default()
	ginHandler.Use(middleware.CORSMiddleware())

	routeManager := router.NewRouteManager(ginHandler)
	routeManager.Register()

	srv := &http.Server{
		Addr:    ":" + httpServerPort,
		Handler: ginHandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server %s\n", err)
		}
	}()

	return &HttpServer{
		handler:    ginHandler,
		httpServer: srv,
	}
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
