package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var botID = os.Getenv("BOT_ID")
var port = os.Getenv("PORT")

// GroupMeMessage is a single message in a GroupMe chat
type GroupMeMessage struct {
	Text string `json:"text"`
}

var eeyores = map[string][]byte{
	"DO YOU":       []byte(fmt.Sprintf("{\"bot_id\": \"%s\", \"text\": \"%s\"}", botID, "Wish I could say yes, but I can't.")),
	"GOOD MORNING": []byte(fmt.Sprintf("{\"bot_id\": \"%s\", \"text\": \"%s\"}", botID, "If it is a good morning, which I doubt.")),
	"DID YOU":      []byte(fmt.Sprintf("{\"bot_id\": \"%s\", \"text\": \"%s\"}", botID, "Most likely lose it again, anyway.")),
	"HOW LONG":     []byte(fmt.Sprintf("{\"bot_id\": \"%s\", \"text\": \"%s\"}", botID, "Days. Weeks. Months. Who knows?")),
}

func groupMeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var m GroupMeMessage
	err := decoder.Decode(&m)
	if err != nil {
		return
	}
	t := strings.ToUpper(m.Text)
	if strings.Contains(t, "@ARYA") {
		for k, v := range eeyores {
			if strings.Contains(t, k) {
				r, err := http.NewRequest("POST", "https://api.groupme.com/v3/bots/post", bytes.NewBuffer(v))
				if err != nil {
					return
				}

				r.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				if _, err := client.Do(r); err != nil {
					return
				}

				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/groupme", groupMeHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
