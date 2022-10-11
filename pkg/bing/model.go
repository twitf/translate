package bing

type Result []struct {
	DetectedLanguage struct {
		Language string  `json:"language"`
		Score    float64 `json:"score"`
	} `json:"detectedLanguage"`
	Translations []struct {
		Text            string `json:"text"`
		Transliteration struct {
			Text   string `json:"text"`
			Script string `json:"script"`
		} `json:"transliteration"`
		To      string `json:"to"`
		SentLen struct {
			SrcSentLen   []int `json:"srcSentLen"`
			TransSentLen []int `json:"transSentLen"`
		} `json:"sentLen"`
	} `json:"translations"`
}

type Config struct {
	Key   string
	Token string
	IG    string
	IID   string
}
