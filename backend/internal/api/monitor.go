package api

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Olá do pacote Api 🚀")
}