package validator_utils

import(
  "fmt"
  "strconv"
  "github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
  "gopkg.in/go-playground/validator.v9"
  en_translations "gopkg.in/go-playground/validator.v9/translations/en"
  "github.com/sampila/go-utils/logger"
)

func ValidateInputs(dataSet interface{}) (bool, string){

  translator := en.New()
  uni := ut.New(translator, translator)

  // this is usually known or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, found := uni.GetTranslator("en")
	if !found {
		logger.Info("translator not found")
	}

  v := validator.New()

  if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		logger.Error("error while registering translator",err)
	}

  _ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

  _ = v.RegisterTranslation("notblank", trans, func(ut ut.Translator) error {
		return ut.Add("notblank", "{0} cannot be empty value", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("notblank", fe.Field())
		return t
	})

  _ = v.RegisterTranslation("numeric", trans, func(ut ut.Translator) error {
		return ut.Add("numeric", "{0} is contain(s) invalid character", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("numeric", fe.Field())
		return t
	})

  _ = v.RegisterTranslation("name", trans, func(ut ut.Translator) error {
		return ut.Add("name", "{0} cannot contains special character", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("name", fe.Field())
		return t
	})

  _ = v.RegisterValidation("notblank", func(fl validator.FieldLevel) bool{
    return fl.Field().String() != " "
  })

  _ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		  return len(fl.Field().String()) > 6
  })

  _ = v.RegisterValidation("name", func(fl validator.FieldLevel) bool {
    //TO-DO : check to db if exists
    return len(fl.Field().String()) > 12
  })

  _ = v.RegisterValidation("numeric", func(fl validator.FieldLevel) bool {
    _, err := strconv.ParseFloat(fl.Field().String(), 64)
    return err == nil
  })



  err := v.Struct(dataSet)

  if err != nil {
      //Validation syntax is invalid
      if err,ok := err.(*validator.InvalidValidationError);ok{
          panic(err)
      }
      //errors := make([]interface{}, len(err.(validator.ValidationErrors)))
      //Validation errors occurred
      for _,err := range err.(validator.ValidationErrors){
          //If json tag doesn't exist, use lower case of name
          //name := strings.ToLower(err.StructField())
          if err != nil {
            return false, fmt.Sprintf(err.Translate(trans))
          }
      }

      //return false,errors
  }
  return true,""
}
