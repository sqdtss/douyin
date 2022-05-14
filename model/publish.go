package model

type PublishResponse struct {
	Response
	VideoInfoList []VideoInfo `json:"video_list"`
}
