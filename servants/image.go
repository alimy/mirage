// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/alimy/mirage/internal"
	"github.com/alimy/mirage/internal/dao"
	"github.com/alimy/mirage/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type imageSrv struct {
	api.UnimplementedImageServant
	base

	broker dao.Broker
}

func (s *imageSrv) ListImage(c *gin.Context) {
	images, err := s.broker.ListImage()
	s.resp(c, images, err)
}

func (s *imageSrv) ImageInfo(c *gin.Context) {
	id := c.Param("imageId")
	info, err := s.broker.ImageInfo(id)
	s.resp(c, info, err)
}

func (s *imageSrv) DeleteImage(c *gin.Context) {
	id := c.Param("imageId")
	force, _ := strconv.ParseBool(c.Param("force"))
	err := s.broker.DeleteImage(id, force)
	s.reply(c, err)
}

func (s *imageSrv) TagImage(c *gin.Context) {
	source := c.Query("source")
	tag := c.Query("tag")
	err := s.broker.TagImage(source, tag)
	s.reply(c, err)
}

func (s *imageSrv) SaveImage(c *gin.Context) {
	id := c.Param("imageId")
	if reader, err := s.broker.SaveImage(id); err != nil {
		c.Header("Content-Type", "application/force-download")
		c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%s.tar.gz", id))
		c.Header("Content-Transfer-Encoding", "binary")
		c.Status(http.StatusOK)
		buf := &bytes.Buffer{}
		buf.ReadFrom(reader)
		buf.WriteTo(c.Writer)
	} else {
		s.error(c, err)
	}
}

func (s *imageSrv) PullImage(c *gin.Context) {
	ref := strings.Trim(c.Query("refStr"), " ")
	reader, err := s.broker.PullImage(ref)
	if err != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		s.ok(c, buf.Bytes())
	} else {
		s.error(c, err)
	}
}

func newImageSrv() api.Image {
	return &imageSrv{
		base:   base{},
		broker: internal.MyBroker(),
	}
}
