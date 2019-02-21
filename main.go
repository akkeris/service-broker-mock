package main

import (
       "github.com/go-martini/martini"
        "github.com/martini-contrib/render"
)

func main() {
        m := martini.Classic()
        m.Use(render.Renderer())
        m.Get("/v1/:service/plans", plans)
        m.Run()
}

func plans(params martini.Params, r render.Render) {
        plans := make(map[string]interface{})
        plans["n/a"] = "Service Not Available in this Region"
        r.JSON(200, plans)
}
