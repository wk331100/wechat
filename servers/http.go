package servers

import (
    "wechat/servers/https"
)

type HttpServer struct {
    router    *https.ApiRouter
    bootstrap *https.Bootstrap
}

func (h *HttpServer) Run() {
    routeMap := h.router.Map()
    h.bootstrap.Run(routeMap)
}
