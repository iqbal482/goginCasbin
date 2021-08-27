package routes

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"goginCasbin/controllers"
	"goginCasbin/middleware"
	"goginCasbin/repository"
	"gorm.io/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	// Initialize casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter in rbac_model casbin
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	//add policy
	if hasPolicy := enforcer.HasPolicy("doctor", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "report", "read")
	}
	if hasPolicy := enforcer.HasPolicy("doctor", "report", "write"); !hasPolicy {
		enforcer.AddPolicy("admin", "report", "write")
	}
	if hasPolicy := enforcer.HasPolicy("patient", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("user", "report", "read")
	}

	userRepository := repository.NewUserRepository(db)

	userController := controllers.NewUserController(userRepository)

	apiRoutes := httpRouter.Group("/api")

	{
		apiRoutes.POST("/register", userController.AddUser(enforcer))
		apiRoutes.POST("/signin", userController.Login)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", middleware.Authorize("report", "read", enforcer), userController.GetAllUser)
		userProtectedRoutes.GET("/:user", middleware.Authorize("report", "read", enforcer), userController.GetUser)
		userProtectedRoutes.PATCH("/:user", middleware.Authorize("report", "write", enforcer), userController.UpdateUser)
		userProtectedRoutes.DELETE("/:user", middleware.Authorize("report", "write", enforcer), userController.DeleteUser)
	}

	httpRouter.Run()

}