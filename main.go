package main

import (
	"log"
	"net/http"
)

func main() {

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("PORT env is required")
	// }

	// instanceID := os.Getenv("INSTANCE_ID")

	// export PORT=8080
	// export INSTANCE_ID="my first instance"

	port := "9991"
	instanceID := "my first instance"

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		text := "Hello world"
		if instanceID != "" {
			text = text + ". From " + instanceID
		}

		w.Write([]byte(text))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = "0.0.0.0:" + port

	log.Println("server starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
