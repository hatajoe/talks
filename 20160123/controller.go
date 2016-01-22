package membercontroller

import (
	"common"
	"models/member"
	"net/http"

	"github.com/goji/context"
	"github.com/zenazn/goji/web"
)

func GetMember(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := context.FromC(c)
	uuid, err := ctx.Value("UUID").(string)
	if err != nil {
		common.ResponseError(err, http.StatusBadRequest, w)
		return
	}
	// type ResponseJson map[string]interface{}
	(&common.ResponseJson{
		"member": member.GetMember(uuid),
	}).Write(w)
}
