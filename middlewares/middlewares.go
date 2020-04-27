package middlewares

import(
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func ApiMiddleware(next httprouter.Handle) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    log.Printf("msg: API Logging, Method: %s, URI: %s\n", r.Method, r.RequestURI)
  	header := w.Header()
  	header.Add("Access-Control-Allow-Origin", "*")
  	header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
  	header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
  	header.Add("Access-Control-Allow-Credentials", "true")

  	if r.Method == "OPTIONS" {
  		w.WriteHeader(http.StatusOK)
  		return
  	}
    next(w, r, p)
  }
}
// Middleware for a standard handler returning a "github.com/julienschmidt/httprouter" Handle
func StdToJulienMiddleware(next http.Handler) httprouter.Handle {

    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        // do stuff
        next.ServeHTTP(w, r)
    }
}
