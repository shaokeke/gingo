package user

import "github.com/songcser/gingo/utils"

type Response struct {
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Enable int            `json:"enable"`
	Geom   utils.GeoPoint `json:"geom"`
}
