package main

import (
	"log"
	"os"

	"github.com/JJDoneAway/addressbook/controllers"
	"github.com/JJDoneAway/addressbook/middleware"
	"github.com/JJDoneAway/addressbook/middleware/slog_gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

// @title           Appetite POC App Example: Address Book
// @version         1.0
// @description     This is a simple show case app for the Appetite POC

// @Accept json
// @Produce json

// @contact.name   	Johannes Höhne
// @contact.email  	Johannes.Hoehne@mail.schwarz
// @contact.url     https://vcs.sys.schwarz/hoehnejo

// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html

// @host 		  	localhost:8080
// @schemes 		http

// @tag.name        addresses
// @tag.description Managing your addresses

// @tag.name        status
// @tag.description status of the app

// @tag.name        metrics
// @tag.description prometheus metrics
func main() {
	if ok := godotenv.Load(); ok != nil {
		slog.Error("ENVIRONMENT", "Message", "Error loading .env file", "Error", ok)
	}

	router := getGin()

	// playground to sniff some http traffic
	// router.Any("webhook", func(c *gin.Context) {

	// 	body, _ := io.ReadAll(c.Request.Body)
	// 	fmt.Printf("HOST: '%v' | METHOD: '%v' | URL: '%v' | BODY: '%v' \n", c.Request.Host, c.Request.Method, c.Request.URL, string(body))
	// 	c.String(200, "OK")
	// })

	addCors(router)

	middleware.RegisterSwagger(router)

	middleware.RegisterPrometheus(router)

	controllers.AddStatus(router)

	controllers.AddAddressRouts(router)

	go middleware.AddDummies()

	log.Fatal(router.Run(":8080"))

}
func addCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	router.Use(cors.New(config))
}

func getGin() *gin.Engine {
	switch os.Getenv("ODJ_DEP_BACKEND_GIN_MODE") {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)
	default:
		slog.Error("ENV", "MESSAGE", "ODJ_DEP_BACKEND_GIN_MODE is not set as environment")
		gin.SetMode(gin.ReleaseMode)
	}

	level := slog.LevelError
	switch os.Getenv("ODJ_DEP_BACKEND_LOGGING") {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "ERROR":
		level = slog.LevelError
	default:
		slog.Error("ENV", "MESSAGE", "ODJ_DEP_BACKEND_LOGGING is not set as environment")
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	logger := slog.New(h)
	slog.SetDefault(logger)
	gin.DefaultWriter = slog.NewLogLogger(logger.Handler(), level).Writer()
	gin.DefaultErrorWriter = slog.NewLogLogger(logger.Handler(), level).Writer()

	engine := gin.New()
	engine.Use(slog_gin.New(logger))
	return engine
}
