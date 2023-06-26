package volcengine

type RequestParam struct {
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
	Text           string `json:"text"`
}

type Result struct {
	Translation      string `json:"translation"`
	DetectedLanguage string `json:"detected_language"`
	Probability      int    `json:"probability"`
	BaseResp         struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}
