package entities

import "time"

type BaseRequest struct {
	CreationTime time.Time `json:"creation_time"`
}
