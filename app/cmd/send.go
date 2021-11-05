package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	OK   = "\xF0\x9F\x86\x97"
	SOS  = "\xF0\x9F\x86\x98"
	CHEK = "\xE2\x9C\x85"
	EYE  = "\xF0\x9F\x91\x80"
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
	err := Send(c)
	if err != nil {
		return err
	}
	return nil
}

func Send(c *SendCommand) error {
	url, token, cid, message := c.URL, c.Token, c.Cid, c.Message

	url = url + "/bot" + token + "/sendMessage"
	text := strings.Replace(message, "\\n", "\n", -1)
	text = strings.Replace(text, "%F0%9F%86%97", OK, -1)
	text = strings.Replace(text, "%F0%9F%86%98", SOS, -1)
	text = strings.Replace(text, "%E2%9C%85", CHEK, -1)
	text = strings.Replace(text, "%F0%9F%91%80", EYE, -1)

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
		c.Log.Fatal(err)
	}

	rb, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		c.Log.Fatal(err)
	}

	var m Message
	err = json.Unmarshal(rb, &m)
	if err != nil {
		c.Log.Fatal(err)
	}

	if m.Ok != true {
		c.Log.Fatal(m.Description)
	}

	c.Log.Printf("[DEBUG] From: %s, MessageId: %d, Text: \"%s\"",
		m.Result.From.FirstName, m.Result.MessageId, m.Result.Text)
	return nil
}
