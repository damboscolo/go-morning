package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	env_token   string = os.Getenv("TELEGRAM_BOT_TOKEN")
	env_chat_id string = os.Getenv("TELEGRAM_GROUP_ID")
)

func SendTelegramMessage(text string) {
	url := "https://api.telegram.org/bot" + env_token + "/sendMessage"

	body := map[string]string{
		"chat_id":    env_chat_id,
		"text":       text,
		"parse_mode": "markdown",
	}
	json_data, err := json.Marshal(body)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	res_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("SendTelegramMessage response=", string(res_body))
}
