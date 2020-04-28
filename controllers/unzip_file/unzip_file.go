package unzip_file

import(
  "os"
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/mholt/binding"
  "github.com/sampila/go-receive-file/domain/unzip_file"
  "github.com/sampila/go-receive-file/utils/http_utils"
  "github.com/sampila/go-receive-file/utils/validator_utils"
  "github.com/sampila/go-utils/file_utils"
  "github.com/sampila/go-utils/rest_ok"
  "github.com/sampila/go-utils/rest_errors"
  "github.com/sampila/go-utils/logger"
)


func Post(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  req := new(unzip_file.UnzipFileForm)
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

  if _, err := os.Stat(req.SourcePath); err != nil {
    logger.Info("file not exists")
    restErr := rest_errors.NewBadRequestError(err.Error())
    http_utils.RespondError(w, restErr)
    return
  }

  unzipErr := file_utils.Unzip(req.SourcePath, req.TargetPath)
  if unzipErr != nil {
    logger.Info("Error unzipping file")
    restErr := rest_errors.NewBadRequestError(unzipErr.Error())
    http_utils.RespondError(w, restErr)
  }

  res := rest_ok.NewRestOK("file unzipped",
                            http.StatusOK,
                            false,
                            "",
                            "OK",
                            1)
  http_utils.RespondJson(w, http.StatusOK, res)
}
