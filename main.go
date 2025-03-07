package main

import (
	"log"
	"math/rand"
	"net/http"

	"secret-santa/models"
	"secret-santa/routers"

	"github.com/gin-gonic/gin"
	gossr "github.com/natewong1313/go-react-ssr"
)

var APP_ENV string

func main() {
	g := gin.Default()
	g.StaticFile("favicon.ico", "./frontend/public/favicon.ico")
	g.Static("/assets", "./frontend/public")
	engine, err := gossr.New(gossr.Config{
		AppEnv:             APP_ENV,
		AssetRoute:         "/assets",
		FrontendDir:        "./frontend/src",
		GeneratedTypesPath: "./frontend/src/generated.d.ts",
		LayoutCSSFilePath:  "Home.css",
		PropsStructsPath:   "./models/props.go",
	})
	if err != nil {
		log.Fatal("Failed to init go-react-ssr")
	}

	g.GET("/", func(c *gin.Context) {
		c.Writer.Write(engine.RenderRoute(gossr.RenderConfig{
			File:  "Home.tsx",
			Title: "Gin example app",
			MetaTags: map[string]string{
				"og:title":    "Gin example app",
				"description": "Hello world!",
			},
			Props: &models.IndexRouteProps{
				InitialCount: rand.Intn(100),
			},
		}))
	})
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	routers.SetupUserRoute(g)
	routers.SetupPoolRouter(g)
	g.Run()
}
