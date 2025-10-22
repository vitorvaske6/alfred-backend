package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vitorvaske6/alfred-backend/src/utils"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var localizerUtil = utils.NewLocalizerUtil()
var envVars = utils.NewEnvironmentVariables()

func Home(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	fmt.Println(lang)
	accept := r.Header.Get("Accept-Language")
	localizerUtil.SetLocalizer(lang, accept)

	welcomeConfig := i18n.LocalizeConfig{
		MessageID: "home.welcome_message",
	}
	appConfig := i18n.LocalizeConfig{
		MessageID: "app.name",
		TemplateData: map[string]string{
			"Name": envVars.GetEnv("NAME"),
		},
	}
	versionConfig := i18n.LocalizeConfig{
		MessageID: "app.version",
		TemplateData: map[string]string{
			"Version": envVars.GetEnv("VERSION"),
		},
	}

	fmt.Fprintf(w, "<h1>%s</h1><p>%s</p><p>%s</p>", localizerUtil.Localize(welcomeConfig), localizerUtil.Localize(appConfig), localizerUtil.Localize(versionConfig))
}

func HandleRequests() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("\nServer running on http://localhost:8080")
	HandleRequests()
}
