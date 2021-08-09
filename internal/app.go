package internal

import (
	"ai4am.com/go-restapi/internal/config"
	"ai4am.com/go-restapi/internal/routes"
	"ai4am.com/go-restapi/pkg/middlewares"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// Application entry point
type App struct {
	Config config.Configuration
	App    *gin.Engine
}

// Starts the application
//
// parameters:
//		logFilePath: Log file path
func (a *App) Start(logFilePath string) {
	gin.SetMode(gin.ReleaseMode)
	a.App = gin.New()
	config.SetupDB(a.Config.Database)
	a.SetLogger(logFilePath)
	a.App.Use(gin.Recovery())
	a.SetMiddlewares()
	a.SetRoutes()
}

// Sets logger for the application
//
// parameters:
//		logFilePath: Log file path
func (a *App) SetLogger(logFilePath string) {
	f, _ := os.Create(logFilePath)
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)
	a.App.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
}

// Sets APP Middlewares
func (a *App) SetMiddlewares() {
	a.App.Use(middlewares.CORS())
	a.App.NoMethod(middlewares.NoMethodHandler())
	a.App.NoRoute(middlewares.NoRouteHandler())
}

// Set APP Routes & Group Routes
func (a *App) SetRoutes() {
	api := a.App.Group("/api")
	// ================== API Routes
	routes.SetAuthRoutes(api)
	routes.SetUsersRoutes(api)
	routes.SetTasksRoutes(api)
	// ================== Docs Routes
	a.App.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}