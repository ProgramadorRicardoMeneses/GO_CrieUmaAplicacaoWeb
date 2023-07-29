package main

import (
	"net/http"

	"github.com/ProgramadorRicardoMeneses/GO_CrieUmaAplicacaoWeb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
