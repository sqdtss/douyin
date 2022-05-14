package model

type FeedResponse struct {
	Response
	VideoInfoList []VideoInfo `json:"video_list"`
	NextTime      int64       `json:"next_time"`
}
