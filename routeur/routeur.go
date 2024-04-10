package routeur

import (
	"MGCS/controller"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Gestionnaire des routes
func Serveur() {
	http.HandleFunc("/index", controller.IndexHandler)
	http.HandleFunc("/contact", controller.ContactHandler)
	http.HandleFunc("/equipe", controller.EquipeHandler)
	http.HandleFunc("/services", controller.ServicesHandler)
	http.HandleFunc("/podcast", controller.PodcastHandler)
	http.HandleFunc("/poles", controller.PolesHandler)
	http.HandleFunc("/solo1", controller.Solo1Handler)
	http.HandleFunc("/solo2", controller.Solo2Handler)
	http.HandleFunc("/solo3", controller.Solo3Handler)
	http.HandleFunc("/solo4", controller.Solo4Handler)
	http.HandleFunc("/solo5", controller.Solo5Handler)
	http.HandleFunc("/courtmetrage", controller.CourtmetrageHandler)
	http.HandleFunc("/treatmentmail", controller.TreatmentMailHandler)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	// Lance le serveur
	runServer()
}

// Fonction qui permet de lancer le serveur
func runServer() {
	port := "localhost:8080"
	url := "http://" + port + "/index"
	go http.ListenAndServe(port, nil)
	fmt.Println("Server is running...")
	time.Sleep(time.Second * 3)
	fmt.Println("If the navigator didn't open on its own, just go to ", url, " on your browser.")
	isRunning := true
	for isRunning {
		fmt.Println("If you want to end the server, type 'stop' here :")
		var command string
		fmt.Scanln(&command)
		if command == "stop" {
			isRunning = false
		}
	}
}
