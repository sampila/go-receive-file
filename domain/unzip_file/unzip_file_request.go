package unzip_file

import(
  "net/http"

  "github.com/mholt/binding"
)

type UnzipFileForm struct{
  StoreCode  string     `json:"store_code" validate:"required"`
  StoreName  string     `json:"store_name" validate:"required"`
  SourcePath string     `json:"source_path" validate:"required"`
	TargetPath string     `json:"target_path" validate:"required"`
  Type       string     `json:"type" validate:"required"`
}

func (f *UnzipFileForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.StoreCode: "store_code",
		&f.StoreName: "store_name",
    &f.SourcePath: "source_path",
		&f.TargetPath: "target_path",
    &f.Type: "type",
	}
}
