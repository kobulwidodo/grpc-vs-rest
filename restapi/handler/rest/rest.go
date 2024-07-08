package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restapi/business/usecase"
	"sync"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var once = &sync.Once{}

type REST interface {
	Run()
}

type rest struct {
	http *gin.Engine
	uc   *usecase.Usecase
}

func Init(uc *usecase.Usecase) REST {
	r := &rest{}
	once.Do(func() {
		gin.SetMode(gin.ReleaseMode)

		httpServ := gin.Default()

		r = &rest{
			http: httpServ,
			uc:   uc,
		}

		r.http.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"*"},
			AllowMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
		}))

		// Set Recovery
		r.http.Use(gin.Recovery())

		r.Register()
	})

	return r
}

func (r *rest) Run() {
	port := ":8080"

	server := &http.Server{
		Addr:    port,
		Handler: r.http,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(fmt.Sprintf("Serving HTTP error: %s", err.Error()))
		}
	}()
	fmt.Println(fmt.Sprintf("Listening and Serving HTTP on %s", server.Addr))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10000)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	log.Println("Server exiting")
}

func (r *rest) Register() {
	publicApi := r.http.Group("/public")
	publicApi.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})

	api := r.http.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/users", r.GetListUsers)

	v1.GET("/test-large-data", r.GetLargeData)
}
