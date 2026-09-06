package webrtc

type RTPRtxParameters struct {
	SSRC SSRC `json:"ssrc"`
}

type RTPFecParameters struct {
	SSRC SSRC `json:"ssrc"`
}

type RTPCodingParameters struct {
	RID         string           `json:"rid"`
	SSRC        SSRC             `json:"ssrc"`
	PayloadType PayloadType      `json:"payloadType"`
	RTX         RTPRtxParameters `json:"rtx"`
	FEC         RTPFecParameters `json:"fec"`
}
