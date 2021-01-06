package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(Image))
}

type Image struct {
	GetImageList Get `mir:"/api/image"`
	GetImageInfo Get `mir:"/api/image/info/:imageId"`
	DeleteImage  Get `mir:"/api/image/remove/:imageId/:forge"`
	TagImage     Get `mir:"/api/image/tag"`
	SaveImage    Get `mir:"/api/image/save/:imageId"`
	PullImage    Get `mir:"/api/image/pull"`
}
