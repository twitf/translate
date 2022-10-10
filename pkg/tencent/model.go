package tencent

type Result struct {
	SessionUUID string `json:"sessionUuid"`
	Translate   struct {
		ErrCode     int    `json:"errCode"`
		ErrMsg      string `json:"errMsg"`
		SessionUUID string `json:"sessionUuid"`
		Source      string `json:"source"`
		Target      string `json:"target"`
		Records     []struct {
			SourceText string `json:"sourceText"`
			TargetText string `json:"targetText"`
			TraceID    string `json:"traceId"`
		} `json:"records"`
		Full    bool `json:"full"`
		Options struct {
		} `json:"options"`
	} `json:"translate"`
	Dict struct {
		Data []struct {
			Word   string `json:"word"`
			EnHash string `json:"en_hash,omitempty"`
		} `json:"data"`
		ErrCode int    `json:"errCode"`
		ErrMsg  string `json:"errMsg"`
		Type    string `json:"type"`
		Map     struct {
			Life struct {
			} `json:"Life"`
			Is struct {
			} `json:"is"`
			Not struct {
			} `json:"not"`
			Only struct {
			} `json:"only"`
			The struct {
				DetailID string `json:"detailId"`
			} `json:"the"`
			Present struct {
			} `json:"present"`
			But struct {
			} `json:"but"`
			Also struct {
			} `json:"also"`
			Tomorrow struct {
				DetailID string `json:"detailId"`
			} `json:"tomorrow"`
			And struct {
				DetailID string `json:"detailId"`
			} `json:"and"`
			Day struct {
				DetailID string `json:"detailId"`
			} `json:"day"`
			After struct {
				DetailID string `json:"detailId"`
			} `json:"after"`
		} `json:"map"`
	} `json:"dict"`
	Suggest interface{} `json:"suggest"`
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
}

type Config struct {
	Qtv string `json:"qtv"`
	Qtk string `json:"qtk"`
}
