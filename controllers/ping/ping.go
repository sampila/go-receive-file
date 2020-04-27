package ping

import(
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/sampila/go-receive-file/utils/http_utils"
)

func PingGet(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  http_utils.RespondJson(w, http.StatusOK, "pong")
}
