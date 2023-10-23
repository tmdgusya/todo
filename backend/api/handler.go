package api

import (
	"net/http"
	"regexp"
)

type paramHanlder func(http.ResponseWriter, *http.Request, []string)

type paramRoute struct {
	method  string
	re      *regexp.Regexp
	handler paramHanlder
}

type route struct {
	method  string
	re      *regexp.Regexp
	handler http.HandlerFunc
}

type RegexpServeMux struct {
	routes      []*route
	paramRoutes []*paramRoute
}

func (r *RegexpServeMux) AddParamRoute(method string, pattern *regexp.Regexp, handler paramHanlder) {
	r.paramRoutes = append(r.paramRoutes, &paramRoute{method, pattern, handler})
}

func (r *RegexpServeMux) AddRoute(method string, pattern *regexp.Regexp, handler http.HandlerFunc) {
	r.routes = append(r.routes, &route{method, pattern, handler})
}

func (r *RegexpServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	found := false

	for i := range r.routes {
		if r.routes[i].method == req.Method {
			if r.routes[i].re.MatchString(req.URL.Path) {
				r.routes[i].handler.ServeHTTP(w, req)
				found = true
				return
			}
		}
	}

	for i := range r.paramRoutes {
		if r.paramRoutes[i].method == req.Method {
			if matches := r.paramRoutes[i].re.FindStringSubmatch(req.URL.Path); len(matches) > 0 {
				r.paramRoutes[i].handler(w, req, matches)
				found = true
				return
			}
		}
	}

	if !found {
		http.NotFound(w, req)
	}
}
