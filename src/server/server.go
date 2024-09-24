package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	responses "template-api-fiber-golang/src/server/response"
)

/*
Responsible to implement the API server logical with management of router.
This server impement the library "Fiber".
*/
type Server struct {
	Port   int
	Router *fiber.Router
	app    *fiber.App
}

// Create a instance of the server setting the port and if the swagger doc must be expose.
func CreateServer(port int) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	// Configure Prefix group
	router := app.Group(getPrefix())
	// Configure logs
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}]${locals:requestid} ${method} ${path} ${status} ${latency}\u200b\n",
		TimeFormat: "02/01/2006:15:04:05",
	}))
	// Configure Swagger
	if os.Getenv("GO_ENV") != "production" {
		doc := LoadDoc(getUrlApi())
		router.Get("/docs/*", swagger.New(swagger.Config{
			URL:         getPrefix() + "swagger/doc.json",
			DeepLinking: false,
		}))
		router.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
			return c.SendString(doc)
		})
	}
	app.Use(recover.New())
	return &Server{Port: port, Router: &router, app: app}
}

// Run server on port configurated.
func (server *Server) StartServer() {
	var error = server.app.Listen(fmt.Sprintf(":%v", server.Port))
	if error != nil {
		log.Fatalln("Error to start server: ", error)
	}
}

// Stop server running.
func (server *Server) StopServer() {
	var error = server.app.Shutdown()
	if error != nil {
		log.Fatalln("Error to stop server: ", error)
	}
}

// Logical about handle internal error when a request throw a Panic.
func errorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	e := ctx.Status(code).JSON(responses.InternalError(err.Error()))
	if e != nil {
		return e
	}
	// Return from handler
	return nil
}

func getPrefix() string {
	prefix, ok := os.LookupEnv("PREFIX_URL")
	if !ok {
		return "/"
	}
	return prefix
}

// Get URL API string either from env variable or a default URL.
func getUrlApi() string {
	urlApi, ok := os.LookupEnv("URL_API")
	if !ok {
		return fmt.Sprintf("http://localhost:%d", getPort())
	}
	return urlApi
}

// Get PORT API number either from env variable or a default PORT.
func getPort() int {
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return 8080
	}
	port, _ := strconv.Atoi(portStr)
	return port
}
