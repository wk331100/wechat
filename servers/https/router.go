package https

import (
    "wechat/servers/https/controllers"
    "wechat/servers/https/libs"
)

type ApiRouter struct {
}

func (api *ApiRouter) Map() []*libs.RouteMap {
    route := new(libs.Route)

    //配置静态路由
    indexController := new(controllers.IndexController)
    route.Get("/index", indexController.Index)

    return route.Map
}
