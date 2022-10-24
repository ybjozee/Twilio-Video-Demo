package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"video_app/helper"
	"video_app/model"
)

func main() {
	loadEnv()

	c := cors.Default()
	handler := http.HandlerFunc(roomHandler)

	fmt.Printf("Starting server at port 8000\n")

	if err := http.ListenAndServe(":8000", c.Handler(handler)); err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func roomHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	var room model.Room

	json.NewDecoder(request.Body).Decode(&room)

	err := room.Validate()

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		response["message"] = err.Error()
	} else {
		response["jwt"] = helper.GenerateToken(room.Name)
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	writer.Write(jsonResponse)
}
