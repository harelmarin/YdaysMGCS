package controller

import (
	InitTemplate "MGCS/templates"
	"net/http"
)

// Handler pour afficher l'index du site
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "index", nil)
}

// Handler pour afficher les contacts
func ContactHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "contact", nil)
}

// Handler pour afficher l'equipe
func EquipeHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "equipe", nil)
}

// Handler pour afficher les services
func ServicesHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "services", nil)
}

// Handler pour afficher le solo1
func Solo1Handler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "solo1", nil)
}

// Handler pour afficher le solo2
func Solo2Handler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "solo2", nil)
}

// Handler pour afficher le solo3
func Solo3Handler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "solo3", nil)
}

// Handler pour afficher le solo4
func Solo4Handler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "solo4", nil)
}

// Handler pour afficher le solo5
func Solo5Handler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "solo5", nil)
}

// Handler pour afficher la section podcast du site
func PodcastHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "podcast", nil)
}

// Handler pour afficher les Poles du site
func PolesHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "poles", nil)
}

// Handler pour afficher les courts métrage du site
func CourtmetrageHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "courtmetrage", nil)
}

// Handler pour afficher les courts métrage du site
func ClipHandler(w http.ResponseWriter, r *http.Request) {

	InitTemplate.Temp.ExecuteTemplate(w, "clip", nil)
}
func TreatmentMailHandler(w http.ResponseWriter, r *http.Request) {
}
