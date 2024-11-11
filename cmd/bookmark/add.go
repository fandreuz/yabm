package bookmark

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/fandreuz/yabm/model"
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

	return doc.Find("title").Text(), nil
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bookmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("'add' expects only one argument")
		}

		url := args[0]
		title, titleErr := getWebpageTitle(url)
		if titleErr != nil {
			return titleErr
		}

		request := model.BookmarkCreationRequest{Url: url, Title: title}

		bookmark, dbErr := model.AddBookmark(request)
		if dbErr != nil {
			return dbErr
		}
		fmt.Println(bookmark)

		return nil
	},
}
