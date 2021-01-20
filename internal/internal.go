// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"github.com/alimy/mirage/dao"
	"github.com/alimy/mirage/internal/moby"
)

var (
	broker dao.Broker
)

func init() {
	broker = moby.NewBroker()
}

func MyBroker() dao.Broker {
	return broker
}
