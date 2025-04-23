package main

import (
	"context"
	"fmt"
	"net/http"
	"training-project/internal/config"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

type HelloRequest struct {
	Name string `path:"name" doc:"Tên người muốn chào"`
}

type HelloResponse struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {

	config.InitConfig()
	db := config.ConnectDB()
	defer db.Close()

	// ------------------------------------------------------- Not Huma

	// r := gin.Default()

	// api := r.Group("/api")

	// // Initialize repositories
	// userRepository := repository.NewUserRepository(db)

	// // Initialize services
	// userService := service.NewUserService(userRepository)

	// // Initialize handlers
	// handler.NewUserHandler(api, userService)

	// r.Run(":" + config.AppConfig.AppPort)

	// ------------------------------------------------------- Integrate Huma

	r := gin.Default()

	api := humagin.New(r, huma.DefaultConfig("Huma + Gin API", "v1.0.0"))

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/hello/{name}",
		Summary:     "Greet someone",
		Description: "Returns a greeting message",
	}, HelloHandler)

	r.Run(":" + config.AppConfig.AppPort)

}

func HelloHandler(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	resp := &HelloResponse{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", req.Name)
	return resp, nil
}
