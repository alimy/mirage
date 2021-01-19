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
