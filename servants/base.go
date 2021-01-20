// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/mirage/internal/json"
	"github.com/gin-gonic/gin"
)

type response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type render struct {
	data interface{}
}

type base struct{}

func (r *render) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	jsonBytes, err := json.Marshal(r.data)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func (r *render) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
}

func (base) resp(c *gin.Context, data interface{}, err error) {
	if err == nil {
		c.Render(http.StatusOK, &render{
			data: &response{
				Code: "OK",
				Data: data,
			},
		})
	} else {
		c.Render(http.StatusInternalServerError, &render{
			data: &response{
				Code: "ERROR",
				Msg:  err.Error(),
			},
		})
	}
}

func (base) reply(c *gin.Context, err error) {
	if err == nil {
		c.Render(http.StatusOK, &render{
			data: &response{
				Code: "OK",
			},
		})
	} else {
		c.Render(http.StatusInternalServerError, &render{
			data: &response{
				Code: "ERROR",
				Msg:  err.Error(),
			},
		})
	}
}

func (base) ok(c *gin.Context, data interface{}) {
	c.Render(http.StatusOK, &render{
		data: &response{
			Code: "OK",
			Data: data,
		},
	})
}

func (base) success(c *gin.Context) {
	c.Render(http.StatusOK, &render{
		data: &response{
			Code: "OK",
		},
	})
}

func (base) error(c *gin.Context, err error) {
	c.Render(http.StatusInternalServerError, &render{
		data: &response{
			Code: "ERROR",
			Msg:  err.Error(),
		},
	})
}

func (base) failure(c *gin.Context, msg string) {
	c.Render(http.StatusInternalServerError, &render{
		data: &response{
			Code: "ERROR",
			Msg:  msg,
		},
	})
}
