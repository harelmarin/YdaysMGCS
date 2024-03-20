package main

import (
	"MGCS/routeur"
	InitTemplate "MGCS/templates"
)

func main() {
	InitTemplate.InitTemplate()
	routeur.Serveur()
}
