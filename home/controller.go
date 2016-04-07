package home

import (
	"net/http"
	"github.com/pborges/log"
	"github.com/pborges/mvc"
	"path"
)

const TemplatePrefix string = "tmpl/home/"

func RegisterController() (c *Controller) {
	log.WithField("controller", "home").Info("register controller")
	c = new(Controller)
	return c
}

type Controller struct {
}

func (this *Controller)Index(w http.ResponseWriter, r *http.Request) {
	model := mvc.NewViewModel()
	mvc.RenderLayout(model, path.Join(TemplatePrefix, "index.tmpl.html"))(w, r)
}
func (this *Controller)About(w http.ResponseWriter, r *http.Request) {
	model := mvc.NewViewModel()
	mvc.RenderLayout(model, path.Join(TemplatePrefix, "about.tmpl.html"))(w, r)
}
