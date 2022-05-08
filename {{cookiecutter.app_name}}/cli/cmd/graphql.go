package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginzap "github.com/gin-contrib/zap"
	"github.com/spf13/cobra"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent"
	_ "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/runtime"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/db"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/http/handlers"
	mw "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/http/middlewares"
	"go.uber.org/zap"
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "GraphQL server",
	Run:   graphqlServer,
}

func init() {
	rootCmd.AddCommand(graphqlCmd)
}

func graphqlServer(cmd *cobra.Command, args []string) {
	var client *ent.Client
	var err error

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client, err = db.GetClient(Conf)
	if err != nil {
		logger.Fatal("database error", zap.Error(err))
	}
	defer client.Close()

	// Setting up Gin
	router := gin.New()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger, true))
	
	//router.Use(gin.Recovery())
	router.Use(mw.GinContextToContextMiddleware())

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // remove this, to AllowOrigin (see next line)
		//AllowOrigins:     []string{"https://my-grapqhlserver-domain.com", "http://localhost:3000", "http://localhost:4200"},
		AllowMethods:     []string{"POST", "OPTION"}, // "GET", "DELETE", "PATCH", "REDIRECT"
		AllowHeaders:     []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// set trusted proxies - see gingonic
	// router.TrustedPlatform = gin.PlatformCloudflare
	// router.SetTrustedProxies([]string{"192.168.80.0/24"}) // replace this with your ip range

	router.POST("/query", handlers.GraphqlHandler(client, Conf, nil))

	router.GET("/", handlers.PlaygroundHandler())
	logger.Info("{{cookiecutter.app_name}}", zap.String("GraphQL playground listens on /", (*Conf).Server.GetHostWithPort()))

	if err := router.Run((*Conf).Server.GetHostWithPort()); err != nil {
		logger.Fatal("graphql server terminated", zap.Error(err))
	}
}
