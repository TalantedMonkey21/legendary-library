package http

import "net/http"

func LecturesCreateHandler (w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Go!"))
}