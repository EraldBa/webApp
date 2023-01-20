package helpers

import (
	"github.com/EraldBa/webApp/pkg/config"
	"log"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

//func ClientError(w http.ResponseWriter, status int) {
//	app.InfoLog.Println("Client error with status:", status)
//	http.Error(w, http.StatusText(status), status)
//}

func ServerError(w http.ResponseWriter, err error) {
	app.ErrorLog.Printf("%s\n%s\n", err.Error(), debug.Stack())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "user_id")
}

func ErrorCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}
