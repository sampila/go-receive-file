package middlewares

import(
  "net/http"
  "encoding/json"

  "github.com/julienschmidt/httprouter"
  "github.com/sampila/go-oauth/oauth"
  "github.com/sampila/go-utils/rest_errors"
  "github.com/sampila/britenesia-api/utils/http_utils"
  "github.com/sampila/go-utils/logger"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    if err := oauth.AuthenticateRequest(r); err != nil {
  		w.Header().Set("Content-Type", "application/json")
  		w.WriteHeader(err.Status())
  		if a := json.NewEncoder(w).Encode(err); a != nil {
  			logger.Info("Error json: " + a.Error())
  		}
		  return
	  }

    userId := oauth.GetCallerId(r)
  	if userId == 0 {
  		respErr := rest_errors.NewUnauthorizedError("invalid access token")
  		http_utils.RespondError(w, respErr)
  		return
  	}

    logger.Info("Auth middlewares")
    next(w, r, p)
  }
}
