package receive_file

import(
  "os"
  "io"
  //"io"
  //"io/ioutil"
  "fmt"
  "net/http"
  //"strconv"
  //"archive/zip"
  //"bytes"
  "path/filepath"
  "log"

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
  if !reqValid {
    logger.Info("invalid parameter")
    restErr := rest_errors.NewBadRequestError(msgErr)
    http_utils.RespondError(w, restErr)
    return
  }

  var fullPath string

  path := fmt.Sprintf("%s", req.TargetPath)
  if _, err := os.Stat(path); os.IsNotExist(err) {
    os.MkdirAll(path, 0755)
  }

  if req.ThemeFile != nil {
    var handler io.ReadCloser
    var err error
    filename := req.ThemeFile.Filename
    if handler, err = req.ThemeFile.Open();err == nil {

      fileLocation := filepath.Join(path, filepath.Base(filename))
      log.Println(fileLocation)
      targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
      defer targetFile.Close()
      if err != nil {
        logger.Info("target file not found")
        restErr := rest_errors.NewBadRequestError(err.Error())
        http_utils.RespondError(w, restErr)
        return
      }

      if _, err := io.Copy(targetFile, handler); err != nil {
        logger.Info("copy file error")
        restErr := rest_errors.NewBadRequestError(err.Error())
        http_utils.RespondError(w, restErr)
        return
      }

      fullPath = fileLocation
    }
  }

  res := rest_ok.NewRestOK("file received",
                            http.StatusOK,
                            false,
                            "",
                            fullPath,
                            1)
  http_utils.RespondJson(w, http.StatusCreated, res)
}
