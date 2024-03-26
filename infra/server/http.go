package server

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/sznborges/to_do_list/application/service"
	"github.com/sznborges/to_do_list/config"
	"github.com/sznborges/to_do_list/infra/controller"
	"github.com/sznborges/to_do_list/infra/database"
	"github.com/sznborges/to_do_list/infra/logger"
	"github.com/sznborges/to_do_list/infra/repository"
	"github.com/sznborges/to_do_list/infra/route"
	"github.com/sznborges/to_do_list/shutdown"
	response "github.com/sznborges/to_do_list/application/dto"
)

var corsConfig = cors.Config{
	AllowOrigins:     "*",
	AllowMethods:     "GET,POST",
	AllowCredentials: true,
	MaxAge:           2592000, //1 month
}

var staticConfig = fiber.Static{
	Compress:       false,
	ByteRange:      false,
	Browse:         false,
	Download:       true,
	Index:          "layout.csv",
	CacheDuration:  10 * time.Second,
	MaxAge:         604800, // 7 days
	ModifyResponse: nil,
	Next:           nil,
}

// @title Midas Service
// @version v1

// @host main.payment-midas.dev.gcp.gruposbf.com.br
// @schemes http https

// @securityDefinitions.apikey AuthKey
// @in header
// @name Authorization
func StartHTTP() {
	app := fiber.New(fiber.Config{
		ReadTimeout:           config.GetDuration("HTTP_SERVER_READ_TIMEOUT_MILLIS"),
		WriteTimeout:          config.GetDuration("HTTP_SERVER_WRITE_TIMEOUT_MILLIS"),
		DisableStartupMessage: true,
	})
	app.Use(func(c *fiber.Ctx) error {
		defer handleUnexpectedError(c)
		return c.Next()
	})
	app.Use(compress.New())
	if config.GetString("ENV") != "prd" {
		app.Get("/", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/docs/swagger/index.html", http.StatusMovedPermanently)
		}))
		app.Get("/docs/swagger/swagger.json", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./docs/swagger/swagger.json")
		}))
		app.Get("/docs/swagger/*", adaptor.HTTPHandlerFunc(httpSwagger.Handler(
			httpSwagger.URL("/docs/swagger/swagger.json"),
		)))
		corsConfig.AllowOrigins = "*"
	}
	app.Use(cors.New(corsConfig))
	// Health Check
	app.Get("/health", HealthCheck)
	v1 := app.Group("/api/v1")

	// Authentication for /api/v1
	keyToken := config.GetString("HTTP_AUTH_TOKEN")
	v1.Use(keyauth.New(keyauth.Config{
		AuthScheme: "Basic",
		Validator: func(c *fiber.Ctx, key string) (bool, error) {
			if key == keyToken {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(controller.ErrorResponse(fmt.Errorf("Unauthorized")))
		},
	}))
	// Controllers
	route.TaskRouter(v1, *controller.NewTaskController(service.NewTask(repository.NewTaskRepository(database.NewPostgresConnection()))))
	// Static files
	serverStatic(v1)
	// Graceful shutdown
	shutdown.Subscribe(func(ctx context.Context) error {
		addr := fmt.Sprint(":", config.GetString("HTTP_PORT"))
		logger.Logger.Infof("tasks api is ready at %s", addr)
		return app.Listen(addr)
	}, func(ctx context.Context) error {
		return app.Shutdown()
	})
}

// @Summary	Server static files
// @Description	Server static files
// @Tags	Static
// @Accept	*/*
// @Produce	text/plain
// @Router	/api/v1/static	[get]
func serverStatic(v1 fiber.Router) {
	v1.Static("/static", "./static", staticConfig)
}

// @Summary	Show the status of server
// @Description	Show the status of server
// @Tags	Root
// @Accept	*/*
// @Produce	json
// @Success	200	{object}	map[string]string
// @Router	/health	[get]
func HealthCheck(ctx *fiber.Ctx) error {
	res := map[string]string{
		"data": "server is up and running",
	}
	err := ctx.JSON(res)
	if err != nil {
		return err
	}
	return nil
}

func handleUnexpectedError(c *fiber.Ctx) {
	r := recover()
	if r != nil {
		logger.Logger.WithField("error", fmt.Sprintf("%v", r)).WithField("stack", string(debug.Stack())).Errorf("unexpected error handling request %s", c.Route().Path)
		if err, ok := r.(error); ok {
			_ = response.WriteError(c, err)
			return
		}
		c.Status(http.StatusInternalServerError)
	}
}
