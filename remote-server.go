package main

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	http.HandleFunc("/remote_server", RemoteServerHandler)
	log.Info().Msg("server starting in port 8080")
	http.ListenAndServe(":8080", nil)
}

func RemoteServerHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().
		Str("service", "plex-server").
		Msg("recieved")

	w.WriteHeader(http.StatusOK)
}
