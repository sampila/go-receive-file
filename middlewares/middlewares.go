package middlewares

import(
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/jpillora/ipfilter"
  "github.com/tomasen/realip"
)

type ipFilter struct{
  next    http.Handler
  filterIP *ipfilter.IPFilter
}

func IPfilterMiddleware(next http.Handler, ipf *ipfilter.IPFilter) *ipFilter {
	return &ipFilter{next: next, filterIP : ipf}
}

func (m *ipFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//use remote addr as it cant be spoofed
  log.Printf("msg:Logging, Method: %s, URI: %s\n", r.Method, r.RequestURI)
  ip := realip.FromRequest(r)
	//show simple forbidden text
  log.Printf(ip)
	if !m.filterIP.Allowed(ip) {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

  header := w.Header()
  header.Add("Access-Control-Allow-Origin", "*")
  header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
  header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
  header.Add("Access-Control-Allow-Credentials", "true")
  //w.Header().Set("Content-Type", "application/json")
  //w.Header().Set("Access-Control-Allow-Origin", "*")
  //w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
  //w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
  //w.Header().Set("Access-Control-Allow-Credentials", "true")
  if r.Method == "OPTIONS" {
    w.WriteHeader(http.StatusOK)
    return
  }
	//success!
	m.next.ServeHTTP(w, r)
}

func ApiMiddleware(next httprouter.Handle) httprouter.Handle{
  return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    log.Printf("msg: API Logging, Method: %s, URI: %s\n", r.Method, r.RequestURI)
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
