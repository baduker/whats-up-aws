package whats_up_aws

import (
	"encoding/json"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"time"
)

func removeHTMLTags(item string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(item)
}

func parseNews() AWSNewsData {
	news := AWSNewsData{}
	err := json.Unmarshal(fetchData(), &news)

	if err != nil {
		panic(err)
	}

	return news
}

func show(news AWSNewsData, wrap int) {
	today := time.Now().Format("01-02-2006")

	for index, i := range news.Items {
		date := i.Item.AdditionalFields.PostDateTime.Format("01-02-2006")
		if date == today {
			headline := i.Item.AdditionalFields.Headline
			postBody := wordWrap(removeHTMLTags(i.Item.AdditionalFields.PostBody), wrap)
			link := fmt.Sprintf("https://aws.amazon.com%s", i.Item.AdditionalFields.HeadlineURL)
			fmt.Printf("%d. %s\nPublished: %s\n%s\n\n", index+1, headline, date, link)
			fmt.Printf("%s\n\n", postBody)
		} else {
			fmt.Println("There are no news for today.")
		}
	}
}

func ShowNews() {
	show(parseNews(), 80)
}
