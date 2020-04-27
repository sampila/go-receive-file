package app

import(
  "github.com/sampila/go-receive-file/middlewares"
  "github.com/sampila/go-receive-file/controllers/ping"
  "github.com/sampila/go-receive-file/controllers/receive_file"

)

func mapUrls(){
  r.GET("/ping",middlewares.ApiMiddleware(ping.PingGet))
  //r.GET("/show/:code",middlewares.ApiMiddleware(file_storage.Get))
  //v1 := r.Group("/v1",middlewares.AuthMiddleware,middlewares.ApiMiddleware)
  v1 := r.Group("/v1",middlewares.ApiMiddleware)
  //v1.GET("/file/:code",file_storage.Get)
  v1.POST("/receive",receive_file.Post)
  //v1.DELETE("/file/:id",file_storage.Delete)
}
