package controllers

import (
	"net/http"
	"runtime"

	"github.com/dimfeld/httptreemux"
	"github.com/godiscourse/godiscourse/api/config"
	"github.com/godiscourse/godiscourse/api/controllers/admin"
	"github.com/godiscourse/godiscourse/api/durable"
	"github.com/godiscourse/godiscourse/api/views"
)

// RegisterRoutes register all routes
func RegisterRoutes(database *durable.Database, router *httptreemux.TreeMux) {
	router.GET("/_hc", health)
	registerUser(database, router)
	registerCategory(database, router)
	registerTopic(database, router)
	registerComment(database, router)
	admin.RegisterAdminRoutes(database, router)
}

func health(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	views.RenderResponse(w, r, map[string]string{
		"build":      config.BuildVersion + "-" + runtime.Version(),
		"developers": "https://live.godiscourse.com",
	})
}
