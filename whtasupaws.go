package whats_up_aws

import "encoding/json"

func parseNews() AWSNewsData {
	news := AWSNewsData{}
	err := json.Unmarshal(FetchData(), &news)

	if err != nil {
		panic(err)
	}

	return news
}
