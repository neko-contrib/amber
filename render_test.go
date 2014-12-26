package amber

import (
	"github.com/rocwong/neko"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

type person struct {
	Name string
	Age  int
}

func Test_Render(t *testing.T) {
	Convey("Normal Render", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer())
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("index", person{Name: "Neko", Age: 11})
		})
		m.ServeHTTP(w, req)
		So(w.Body.String(), ShouldEqual, "<div>layout</div>\n<div>Neko - 11</div>\n")
	})

	Convey("Compact Render", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		m := neko.New()
		m.Use(Renderer(&Options{BaseDir: "views", Extension: ".amber", PrettyPrint: false}))
		m.GET("/", func(ctx *neko.Context) {
			ctx.Render("index", person{Name: "Neko", Age: 11}, 404)
		})
		m.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, 404)
		So(w.Body.String(), ShouldEqual, "<div>layout</div><div>Neko - 11</div>\n")
	})

}
