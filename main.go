package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Payload struct {
	Message string `json:"message"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Solo se permiten solicitudes POST", http.StatusMethodNotAllowed)
		return
	}

	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	if payload.Message == "" {
		payload.Message = "World"
	}

	response := fmt.Sprintf("Hello, %s!", payload.Message)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": response})
}

func sendRequest() {

	payload := Payload{Message: "Bryan"}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error al convertir payload a JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/webhook", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error en la solicitud: %v", err)
	}
	defer resp.Body.Close()

	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Error al leer la respuesta del servidor: %v", err)
	}

	fmt.Println("Respuesta del servidor:", response["response"])
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	go func() {
		fmt.Println("Servidor escuchando en http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	time.Sleep(1 * time.Second)

	sendRequest()
}
