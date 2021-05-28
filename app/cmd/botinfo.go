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

type GetMe struct {
	Ok     bool `json:"ok"`
	Result struct {
		Id                      int    `json:"id"`
		IsBot                   bool   `json:"is_bot"`
		FirstName               string `json:"first_name"`
		Username                string `json:"username"`
		CanJoinGroups           bool   `json:"can_join_groups"`
		CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
		SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	} `json:"result"`
	Description string `json:"description"`
}

type BotInfoCommand struct {
	CommonOpts
}

func (c *BotInfoCommand) Execute(_ []string) error {

	url := c.URL + "/bot" + c.Token + "/getme"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	rb, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var getMe GetMe
	err = json.Unmarshal(rb, &getMe)
	if err != nil {
		log.Fatal(err)
	}

	if getMe.Ok != true {
		log.Fatal(getMe.Description)
	}

	fmt.Println("\nBOT INFO:")
	fmt.Println("---------")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "Id: \t%d\n", getMe.Result.Id)
	fmt.Fprintf(w, "Username: \t%s\n", getMe.Result.Username)
	fmt.Fprintf(w, "FirstName: \t%s\n\n", getMe.Result.FirstName)
	w.Flush()

	return nil
}
