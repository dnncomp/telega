package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	SOS = "\xF0\x9F\x86\x98"
	OK  = "\xE2\x9C\x85"
	EYE = "\xF0\x9F\x91\x80"
)

type Message struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int    `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			Id    int64  `json:"id"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
	Description string `json:"description"`
}

type SendCommand struct {
	Cid     string `short:"c" long:"cid" env:"TELEGA_CHAT_ID" required:"true" description:"chat id"`
	Message string `short:"m" long:"message" required:"true" default:"hi" description:"text message"`

	CommonOpts
}

func (c *SendCommand) Execute(_ []string) error {
	err := Send(c.URL, c.Token, c.Cid, c.Message)
	if err != nil {
		return err
	}
	return nil
}

func Send(url, token, cid, message string) error {

	url = url + "/bot" + token + "/sendMessage"

	text := strings.Replace(message, "\\n", "\n", -1)
	text = OK + "  " + text

	payload, err := json.Marshal(
		map[string]string{
			"chat_id": cid,
			"text":    text,
			//"parse_mode": "Markdown",
			//"parse_mode": "MarkdownV2",
			"parse_mode": "HTML",
		})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	rb, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var m Message
	err = json.Unmarshal(rb, &m)
	if err != nil {
		log.Fatal(err)
	}

	if m.Ok != true {
		log.Fatal(m.Description)
	}

	fmt.Println("  MessageId:", m.Result.MessageId)
	return nil
}
