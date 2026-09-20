package handler

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "API, restaurante online")

}
