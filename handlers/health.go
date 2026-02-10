package healthCheck

import "net/http"

// HealthCheck godoc
// @Summary Health check
// @Description API health status
// @Tags Health
// @Produce plain
// @Success 200 {string} string "ok"
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
