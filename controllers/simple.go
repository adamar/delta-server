
package controllers

import (
	"net/http"
	"github.com/unrolled/render"
	)


func SimplePage(w http.ResponseWriter, req *http.Request, template string, r *render.Render) {
        r.HTML(w, http.StatusOK, template, nil)
}

