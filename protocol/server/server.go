package server

import (
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsfs embed.FS

//go:embed static/*
var embedDirStatic embed.FS

func RunServer(httpAddr, liveUrl string) {

	hash := time.Now().Unix()

	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{Views: engine})

	//app.Static("/static", "views/static")
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
		Browse:     true,
	}))

	app.Get("/live", func(c *fiber.Ctx) error {
		// Render index - start with views directory
		return c.Render("views/index", fiber.Map{
			"Title": "Live Stream",
			"Hash":  hash,
			"Url":   liveUrl,
		})
	})

	log.Fatal(app.Listen(httpAddr))
}
