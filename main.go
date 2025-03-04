package main

import (
	"fmt"
	"net/http"
	"resources/internal/config"
	statistics "resources/internal/stats"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func getStatsWithUptime(upd websocket.Upgrader, uptime time.Duration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upd.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("Error upgrading connection: %v\n", err)
			return
		}
		defer c.Close()
		for {
			message, err := statistics.GetAllStats()
			if err != nil {
				fmt.Printf("Error getting stats: %v\n", err)
				break
			}
			err = c.WriteMessage(1, message)
			if err != nil {
				fmt.Printf("Error writing message: %v\n", err)
				break
			}
			time.Sleep(uptime)
		}
	}
}

func main() {
	cfg := config.NewConfigService()
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			allowedOrigins := make(map[string]bool, len(cfg.Server.AllowedOrigins))
			for _, origin := range strings.Split(cfg.Server.AllowedOrigins, ",") {
				allowedOrigins[origin] = true
			}
			origin := r.Header.Get("Origin")
			fmt.Println("request origin", origin)
			return allowedOrigins[origin]
		},
	}
	http.HandleFunc("/ws/", getStatsWithUptime(upgrader, cfg.Server.Uptime))
	if err := http.ListenAndServe(
		fmt.Sprintf(
			"%s:%d",
			cfg.Server.Host,
			cfg.Server.Port,
		),
		nil,
	); err != nil {
		panic(err)
	}
}
