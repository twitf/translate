package alibaba

type Csrf struct {
	Token         string `json:"token"`
	ParameterName string `json:"parameterName"`
	HeaderName    string `json:"headerName"`
}

type Result struct {
	RequestID      string `json:"requestId"`
	Success        bool   `json:"success"`
	HTTPStatusCode int    `json:"httpStatusCode"`
	Code           string `json:"code"`
	Message        string `json:"message"`
	Data           struct {
		TranslateText  string `json:"translateText"`
		DetectLanguage string `json:"detectLanguage"`
	} `json:"data"`
}
