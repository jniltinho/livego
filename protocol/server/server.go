package server

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//go:embed views static
var FS embed.FS

func RunServerGin(httpAddr, hlsUrl string) {
	hash := time.Now().Unix()
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(FS, "views/*.html"))
	router.SetHTMLTemplate(templ)

	fe, _ := fs.Sub(FS, "static")
	router.StaticFS("/static", http.FS(fe))

	router.GET("/live/:room", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":  "Live Stream",
			"Hash":   hash,
			"HlsUrl": hlsUrl,
			"Room":   c.Param("room"),
		})
	})

	log.Fatal(router.Run(httpAddr))
}
