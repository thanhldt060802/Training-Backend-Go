package route

import (
	"training-project/internal/handler"
	"training-project/internal/repository"
	"training-project/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func RegisterRoutes(r *gin.Engine, db *bun.DB) {
	// Init repositories
	userRepo := repository.NewUserRepository(db)

	// Init services
	userServ := service.NewUserService(userRepo)

	// Init handlers
	userHandler := handler.NewUserHandler(userServ)

	// Define the routes
	api := r.Group("/api")
	{
		userAPI := api.Group("/users")
		{
			userAPI.GET("/", userHandler.GetAllUsers)
		}
	}
}
