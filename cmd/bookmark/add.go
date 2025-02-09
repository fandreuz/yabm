package bookmark

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fandreuz/yabm/model"
	"github.com/fandreuz/yabm/model/entity"
	"github.com/spf13/cobra"
)

func getWebpageTitle(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("Got HTTP error (%d, %s)", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(doc.Find("title").Text()), nil
}

var AddCmd = &cobra.Command{
	Use:   "add url",
	Short: "Add a new bookmark",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		title, titleErr := getWebpageTitle(url)
		if titleErr != nil {
			return titleErr
		}

		request := entity.BookmarkCreationRequest{Url: url, Title: title}

		bookmark, dbErr := model.CreateBookmark(request)
		if dbErr != nil {
			return dbErr
		}

		fmt.Println(bookmark.String())
		return nil
	},
}
