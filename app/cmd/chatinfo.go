package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

type Update struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Message struct {
			From struct {
				FirstName    string `json:"first_name"`
				Username     string `json:"username"`
				LanguageCode string `json:"language_code"`
			} `json:"from"`
			Chat struct {
				Id    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
		} `json:"message"`
	} `json:"result"`
	Description string `json:"description"`
}

type ChatInfoCommand struct {
	CommonOpts
}

func (c *ChatInfoCommand) Execute(_ []string) error {

	url := c.URL + "/bot" + c.Token + "/getUpdates"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	rb, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var up Update
	err = json.Unmarshal(rb, &up)
	if err != nil {
		log.Fatal(err)
	}

	if up.Ok != true {
		log.Fatal(up.Description)
	}

	fmt.Println("\nCHAT INFO:")
	fmt.Println("----------")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	for _, r := range up.Result {

		c := r.Message.Chat
		f := r.Message.From
		fmt.Fprintf(w, "Title:\t%s   \tId:\t%d   \tType:\t%s   \tFrom:\t%s\n",
			c.Title, c.Id, c.Type, f.Username)
	}
	w.Flush()

	return nil
}
