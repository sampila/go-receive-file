package receive_file

import(
  "os"
  "io"
  "io/ioutil"
  //"fmt"
  "net/http"
  //"strconv"
  "archive/zip"
  "bytes"

  "github.com/julienschmidt/httprouter"
  "github.com/mholt/binding"
  "github.com/sampila/go-receive-file/domain/receive_file"
  "github.com/sampila/go-receive-file/utils/http_utils"
  "github.com/sampila/go-receive-file/utils/validator_utils"
  "github.com/sampila/go-utils/rest_ok"
  "github.com/sampila/go-utils/rest_errors"
  "github.com/sampila/go-utils/logger"
)

func Post(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  req := new(receive_file.ReceiveFileForm)
  bindErr := binding.Bind(r, req)
  if bindErr != nil {
    logger.Info("invalid parameter")
    restErr := rest_errors.NewBadRequestError("error invalid parameter")
    http_utils.RespondError(w, restErr)
    return
  }

  reqValid, msgErr := validator_utils.ValidateInputs(req)
  if reqValid == false {
    logger.Info("invalid parameter")
    restErr := rest_errors.NewBadRequestError(msgErr)
    http_utils.RespondError(w, restErr)
    return
  }

  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
      // err
  }

  fr, err := zip.NewReader(bytes.NewReader(body), r.ContentLength)
  if err != nil {
      // err
  }
  for _, zf := range fr.File {
      dst, err := os.Create(zf.Name)
      if err != nil {
          // err
      }
      defer dst.Close()
      src, err := zf.Open()
      if err != nil {
          // err
      }
      defer src.Close()

      io.Copy(dst, src)
  }

  res := rest_ok.NewRestOK("file received",
                            http.StatusOK,
                            false,
                            "",
                            "OK",
                            1)
  http_utils.RespondJson(w, http.StatusCreated, res)
}
