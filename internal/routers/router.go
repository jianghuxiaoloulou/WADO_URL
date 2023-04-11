package routers

import (
	_ "WowjoyProject/WADO_URL/docs"
	v1 "WowjoyProject/WADO_URL/internal/routers/api/v1"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/wado" {
		v1.FileDownload(w, r)
		return
	}
	http.NotFound(w, r)
	return

}
