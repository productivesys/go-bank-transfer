package action

import "net/http"

//HealthCheck é handler para verificar se a API está UP
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
