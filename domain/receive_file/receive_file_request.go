package receive_file

import(
  "net/http"
  "mime/multipart"

  "github.com/mholt/binding"
)

type ReceiveFileForm struct{
  StoreCode  string                `json:"store_code" validate:"required"`
  StoreName  string                `json:"store_name" validate:"required"`
  TargetPath string                `json:"target_path" validate:"required"`
  Type       string                `json:"type" validate:"required"`
	File      *multipart.FileHeader `json:"file" validate:"required"`
}

func (f *ReceiveFileForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.StoreCode: "store_code",
		&f.StoreName: "store_name",
		&f.TargetPath: "target_path",
		&f.Type: "type",
		&f.File: "file",
	}
}
