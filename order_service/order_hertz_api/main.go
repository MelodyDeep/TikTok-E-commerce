// Code generated by hertz generator.

package main

import (
	"fmt"
	"order/order_hertz_api/config"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	conf := config.GetConfig()
	h := server.New(server.WithHostPorts(fmt.Sprintf("%s:%d", conf.Host, conf.Port)))

	register(h)
	h.Spin()
}
