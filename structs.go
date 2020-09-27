package whats_up_aws

import "time"

// Known Discord API Endpoints.
var (
	EndpointDiscord      = "https://discordapp.com/"
	EndpointAPI          = EndpointDiscord + "api/"
	EndpointWebhooks     = EndpointAPI + "webhooks/"
	EndpointWebhookToken = func(wID, token string) string { return EndpointWebhooks + wID + "/" + token }
)

// RandomGirl is a container for kindgirls.com's model data
type RandomGirl struct {
	URL       string
	Name      string
	ThumbNail string
	Galleries string
	Gallery   string
	Teaser    string
}

// MessageEmbedThumbnail is a part of a MessageEmbed struct.
type MessageEmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// MessageEmbedVideo is a part of a MessageEmbed struct.
type MessageEmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// MessageEmbedFooter is a part of a MessageEmbed struct.
type MessageEmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// MessageEmbedImage is a part of a MessageEmbed struct.
type MessageEmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// MessageEmbedProvider is a part of a MessageEmbed struct.
type MessageEmbedProvider struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

// MessageEmbedAuthor is a part of a MessageEmbed struct.
type MessageEmbedAuthor struct {
	URL          string `json:"url,omitempty"`
	Name         string `json:"name,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// MessageEmbedField is a part of a MessageEmbed struct.
type MessageEmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

// An MessageEmbed stores data for message embeds.
type MessageEmbed struct {
	URL         string                 `json:"url,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Timestamp   string                 `json:"timestamp,omitempty"`
	Color       int                    `json:"color,omitempty"`
	Footer      *MessageEmbedFooter    `json:"footer,omitempty"`
	Image       *MessageEmbedImage     `json:"image,omitempty"`
	Thumbnail   *MessageEmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *MessageEmbedVideo     `json:"video,omitempty"`
	Provider    *MessageEmbedProvider  `json:"provider,omitempty"`
	Author      *MessageEmbedAuthor    `json:"author,omitempty"`
	Fields      []*MessageEmbedField   `json:"fields,omitempty"`
}

// WebHookParams is the main struct used as payload for discord webhook
type WebHookParams struct {
	Content   string          `json:"content,omitempty"`
	Username  string          `json:"username,omitempty"`
	AvatarURL string          `json:"avatar_url,omitempty"`
	TTS       bool            `json:"tts,omitempty"`
	File      string          `json:"file,omitempty"`
	Embeds    []*MessageEmbed `json:"embeds,omitempty"`
}

// AWSNewsData is the main struct that represents the AWS API response
type AWSNewsData struct {
	Metadata struct {
		Count     int `json:"count"`
		TotalHits int `json:"totalHits"`
	} `json:"metadata"`
	FieldTypes struct {
		RelatedBlog  string `json:"relatedBlog"`
		PostBody     string `json:"postBody"`
		ModifiedDate string `json:"modifiedDate"`
		HeadlineURL  string `json:"headlineUrl"`
		PostDateTime string `json:"postDateTime"`
		PostSummary  string `json:"postSummary"`
		Headline     string `json:"headline"`
		ContentType  string `json:"contentType"`
	} `json:"fieldTypes"`
	Items []struct {
		Tags []struct {
			TagNamespaceID string `json:"tagNamespaceId"`
			CreatedBy      string `json:"createdBy"`
			Name           string `json:"name"`
			DateUpdated    string `json:"dateUpdated"`
			Locale         string `json:"locale"`
			LastUpdatedBy  string `json:"lastUpdatedBy"`
			DateCreated    string `json:"dateCreated"`
			Description    string `json:"description"`
			ID             string `json:"id"`
		} `json:"tags"`
		Item struct {
			CreatedBy        string  `json:"createdBy"`
			Locale           string  `json:"locale"`
			Author           string  `json:"author"`
			DateUpdated      string  `json:"dateUpdated"`
			Score            float64 `json:"score"`
			Name             string  `json:"name"`
			NumImpressions   int     `json:"numImpressions"`
			DateCreated      string  `json:"dateCreated"`
			AdditionalFields struct {
				PostBody     string    `json:"postBody"`
				ModifiedDate time.Time `json:"modifiedDate"`
				HeadlineURL  string    `json:"headlineUrl"`
				PostDateTime time.Time `json:"postDateTime"`
				PostSummary  string    `json:"postSummary"`
				ContentType  string    `json:"contentType"`
				Headline     string    `json:"headline"`
			} `json:"additionalFields"`
			ID            string `json:"id"`
			DirectoryID   string `json:"directoryId"`
			LastUpdatedBy string `json:"lastUpdatedBy"`
		} `json:"item"`
	} `json:"items"`
}
