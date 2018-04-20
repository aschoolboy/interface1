package JsonResponse

import (
	"net/http"
)

func JsonResponse(json []byte, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
