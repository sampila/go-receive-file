package app

import(
  "fmt"
  "net/http"
  "time"

  "github.com/sampila/jsonconfig/jsonconfig"
  "github.com/sampila/go-receive-file/router"
  "github.com/NYTimes/gziphandler"
  "github.com/sampila/go-utils/logger"
)

var (
  r = router.New()
  config jsonconfig.Configuration
)

func StartApplication(){
  if err := jsonconfig.Load("./config/config.json", &config); err != nil {
    panic(err)
  }

  mapUrls()

  server := new(http.Server)
	server.Handler = gziphandler.GzipHandler(r)
	server.ReadTimeout = config.Server.ReadTimeout * time.Second
	server.WriteTimeout = config.Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%v", config.Server.HTTPPort)
	//if config.Configuration().Log.Verbose {
	//  log.Printf("Starting server at %s \n", server.Addr)
	//}

	/*errm := gomail.SendMail("alipxsulistio@gmail.com","Mail From Go", "Body email")
	if errm != nil {
		log.Println("error send email")
	}*/
  logger.Info(fmt.Sprintf("starting server at port: %d", config.Server.HTTPPort))

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
