package user

import (
	"net/http"

	contacts "github.com/wshaman/contacts-stub"

	"github.com/wshaman/course-rest/lib"
)

func Search(w http.ResponseWriter, r *http.Request) {
	p := contacts.LoadPersistent()
	num := r.URL.Query().Get("number")
	res, err := p.FindByPhone(num)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, res)
}
