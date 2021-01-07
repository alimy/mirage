// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

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
