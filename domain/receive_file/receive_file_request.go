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
	ThemeFile  *multipart.FileHeader `json:"theme_file" validate:"required"`
}

func (f *ReceiveFileForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.StoreCode: "store_code",
		&f.StoreName: "store_name",
		&f.TargetPath: "target_path",
		&f.Type: "type",
		&f.ThemeFile: "theme_file",
	}
}
