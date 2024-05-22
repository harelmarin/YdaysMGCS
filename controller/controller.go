package controller

import (
	InitTemplate "MGCS/templates"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Charger les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Récupérer les valeurs du formulaire
	name := r.FormValue("nom")
	surname := r.FormValue("prenom")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Créer le corps de l'email
	body := fmt.Sprintf("Nom: %s\nPrénom: %s\nEmail: %s\n\nMessage:\n%s", name, surname, email, message)

	// Configuration SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Destinataires
	to := []string{"yubi2a1812@gmail.com"}

	// En-têtes et corps du message
	msg := []byte("To: yubi2a1812@gmail.com\r\n" +
		"Subject: Nouveau message du site MGCS \r\n" +
		"\r\n" +
		body + "\r\n")

	// Envoyer l'email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, to, msg)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Rediriger après succès
	http.Redirect(w, r, "/contact", http.StatusSeeOther)
}
