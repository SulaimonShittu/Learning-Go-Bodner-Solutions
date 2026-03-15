package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/time", webHandler{})

	server := http.Server{
		Addr:         ":9090",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      serveMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type webHandler struct{}

func (wh webHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		wh.getTime(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), 405)
	}
}

func ipLogger(r *http.Request) {
	ip := r.RemoteAddr
	hOptions := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, hOptions)
	logg := slog.New(jsonHandler)
	logg.Info("user ip address", ip)
}

func (wh webHandler) getTime(w http.ResponseWriter, r *http.Request) {
	ipLogger(r)
	acptVals := r.Header.Values("Accept")
	output := "text"
	for _, v := range acptVals {
		if v == "application/json" {
			output = "json"
		}
	}
	if output == "text" {
		w.Header().Set("Content-Type", "text/html")
		currentTime := time.Now().Format(time.RFC3339)
		_, err := w.Write([]byte(
			fmt.Sprintf(`<h1>Your current time is : %v</h1>`, currentTime),
		))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if output == "json" {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		curTimeJSON := struct {
			DayOfWeek  string `json:"day_of_week"`
			DayOfMonth int    `json:"day_of_month"`
			Month      string `json:"month"`
			Year       int    `json:"year"`
			Hour       int    `json:"hour"`
			Minute     int    `json:"minute"`
			Second     int    `json:"second"`
		}{
			DayOfWeek:  time.Now().Weekday().String(),
			DayOfMonth: time.Now().Day(),
			Month:      time.Now().Month().String(),
			Year:       time.Now().Year(),
			Hour:       time.Now().Hour(),
			Minute:     time.Now().Minute(),
			Second:     time.Now().Second(),
		}
		err := enc.Encode(curTimeJSON)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
}
