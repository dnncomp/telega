package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
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

type GetMeCommand struct {
	CommonOpts
}

func (c *GetMeCommand) Execute(_ []string) error {

	url := c.URL + "/bot" + c.Token + "/getme"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	json1, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var getme GetMe
	err = json.Unmarshal(json1, &getme)
	if err != nil {
		log.Fatal(err)
	}

	if getme.Ok != true {
		log.Fatal(getme.Description)
	}

	color.Danger.Println("\n  GETME COMMAND:\n")
	magenta := color.FgMagenta.Render

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "  Id: \t%s\n", magenta(getme.Result.Id))
	fmt.Fprintf(w, "  Username: \t%s\n", magenta(getme.Result.Username))
	fmt.Fprintf(w, "  FirstName: \t%s\n", magenta(getme.Result.FirstName))
	w.Flush()

	return nil
}
