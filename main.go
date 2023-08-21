package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	route := echo.New()
	route.POST("api/user_registration", UserRegistration)
	route.Start(":8080")
}

type UserRegistrationForm struct {
	Email            string `json:"email" validate:"required,email"`
	Username         string `json:"username" validate:"required,min=5,max=30"`
	NamaLengkap      string `json:"nama_lengkap" validate:"required"`
	JenisKelamin     string `json:"jenis_kelamin" validate:"required,oneof=L P"`
	TanggalLahir     string `json:"tanggal_lahir" validate:"required,datetime=2006-01-02"`
	StatusPernikahan string `json:"status_pernikahan" validate:"required,oneof=menikah belum_menikah"`
	Penghasilan      int    `json:"penghasilan" validate:"required_if=StatusPernikahan menikah"`
	Usia             int    `json:"usia" validate:"gt=15"`
}

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Errors   interface{} `json:"errors,omitempty"`
}

func UserRegistration(c echo.Context) error {
	user := new(UserRegistrationForm)
	if err := c.Bind(user); err != nil {
		return err
	}

	validate := *validator.New()

	err := validate.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		errorList := make(map[string]string)

		for _, e := range errors {
			var errMsg string
			field, _ := reflect.TypeOf(*user).FieldByName(e.StructField())
			fieldName := field.Tag.Get("json")

			switch e.Tag() {
			case "required":
				errMsg = fmt.Sprintf("%s tidak boleh kosong", fieldName)
			case "email":
				errMsg = fmt.Sprintf("%s bukan email yang valid", fieldName)
			case "oneof":
				errMsg = fmt.Sprintf("%s harus berupa %s", fieldName, e.Param())
			case "datetime":
				errMsg = fmt.Sprintf("%s bukan berupa tanggal yang valid", fieldName)
			case "min":
				errMsg = fmt.Sprintf("%s minimal %s karakter", fieldName, e.Param())
			case "max":
				errMsg = fmt.Sprintf("%s maksimal %s karakter", fieldName, e.Param())
			case "required_if":
				errMsg = fmt.Sprintf("%s harus diisi jika %s", fieldName, e.Param())
			case "gt":
				errMsg = fmt.Sprintf("%s harus lebih besar dari %s", fieldName, e.Param())
			}
			errorList[fieldName] = errMsg
		}

		return c.JSON(http.StatusBadRequest, Response{
			Status:   400,
			Errors:   errorList,
			Messages: "Request ditolak",
		})
	}

	// masukkan logic setelah proses validasi selesai dan tidak ada error

	return c.JSON(http.StatusOK, Response{
		Status:   200,
		Messages: "Sukses",
	})
}
