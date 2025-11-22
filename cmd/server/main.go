package main

import (
	"fmt"
	"net/http"

	"github.com/syedrafinulhuq/gotask-engine/internal/config"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.LoadConfig("config.json")
	if err != nil {
		fmt.Println("Config Mismatched and Error :" + err.Error())
	}
	fmt.Println("Starting at server :", cfg.ServerPort)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server Is Running"))
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), nil)
	if err != nil {
		panic(err)
	}
}
