package domain

type Header struct {
	Time     int64  `json:"time"`
	PrevHash string `json:"prev_hash"`
}
