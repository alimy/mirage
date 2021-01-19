package dao

import "github.com/docker/docker/api/types"

type Version struct {
	ClientVer string        `json:"client"`
	ServerVer types.Version `json:"server"`
}
