package main

// Imports
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Définition de la variable d'environnement PING_LISTEN_PORT
	port := os.Getenv("PING_LISTEN_PORT")
	if port == "" {
		port = "8080"
	}
	// Définition du serveur HTTP
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			// Si la méthode de la requête n'est pas GET, retourne une erreur 404
			http.NotFound(w, r)
			return
		}
		// Si la méthode est GET, retourne les headers HTTP de la requête
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r.Header)
	})
	// Démarrage du serveur HTTP
	fmt.Println("Serveur sur :" + port)
	http.ListenAndServe(":"+port, nil)
}
