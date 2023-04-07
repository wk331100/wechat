package libs

const MethodGet = "GET"
const MethodPost = "POST"
const MethodPut = "PUT"
const MethodDelete = "DELETE"
const MethodRequest = "REQUEST"

type Route struct {
    Prefix string
    Map    []*RouteMap
}

type RouteMap struct {
    Path     string
    Function interface{}
    Method   string
}

type FieldParam struct {
    Path     string
    function interface{}
}

func (r *Route) Group(name string, f func()) {
    r.Prefix = name
    f()
    r.Prefix = ""
}

func (r *Route) Get(path string, function interface{}) {
    r.setField(path, function, MethodGet)
}

func (r *Route) Post(path string, function interface{}) {
    r.setField(path, function, MethodPost)
}

func (r *Route) Put(path string, function interface{}) {
    r.setField(path, function, MethodPut)
}

func (r *Route) Delete(path string, function interface{}) {
    r.setField(path, function, MethodDelete)
}

func (r *Route) Any(path string, function interface{}) {
    r.setField(path, function, MethodRequest)
}

func (r *Route) setField(path string, function interface{}, method string) {
    if r.Prefix != "" {
        path = r.Prefix + path
    }
    field := &RouteMap{
        Path:     path,
        Function: function,
        Method:   method,
    }
    r.Map = append(r.Map, field)
}
