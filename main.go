package main

import (
	"github.com/jchavannes/jgo/web"
)

var (
	indexRoute = web.Route{
		Pattern: "/",
		Handler: func(r *web.Response) {
			r.Render()
		},
	}
)

func main() {
	server := web.Server{
		Port: 8256,
		StaticFilesDir: "web",
		TemplatesDir: "web",
		Routes: []web.Route{
			indexRoute,
		},
	}
	server.Run()
}
