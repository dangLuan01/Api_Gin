package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandlerValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationError {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn %s", e.Field(), e.Param())
			case "slug":
				errors[e.Field()] = fmt.Sprintf("%s phải là một slug hợp lệ", e.Field())
			case "required":
				errors[e.Field()] = fmt.Sprintf("%s là trường bắt buộc", e.Field())
			case "min":
				errors[e.Field()] = fmt.Sprintf("%s phải có ít nhất %s ký tự", e.Field(), e.Param())
			case "max":
				errors[e.Field()] = fmt.Sprintf("%s không được vượt quá %s ký tự", e.Field(), e.Param())
			case "url":
				errors[e.Field()] = fmt.Sprintf("%s phải là một URL hợp lệ", e.Field())
			case "minInt":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn %s", e.Field(), e.Param())
			case "maxInt":
				errors[e.Field()] = fmt.Sprintf("%s không được lớn hơn %s", e.Field(), e.Param())
			case "file_ext":
				exts := strings.Split(e.Param(), " ")
				errors[e.Field()] = fmt.Sprintf("%s phải có phần mở rộng là %s", e.Field(), strings.Join(exts, ", "))
			}
		}
		return gin.H{"errors": errors}
	}
	return gin.H{
		"error": "Validation failed",
		"details": err.Error(),
	}
}
func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	v.RegisterValidation("minInt", func(fl validator.FieldLevel) bool {
		min, _ := strconv.ParseInt(fl.Param(), 10, 64)
		if fl.Field().Int() < min  {
			return false
		}
		return true
	})
	v.RegisterValidation("maxInt", func(fl validator.FieldLevel) bool {
		max, _ := strconv.ParseInt(fl.Param(), 10, 64)
		if fl.Field().Int() > max  {
			return false
		}
		return true
	})
	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		fileName := fl.Field().String()
		ext := fileName[strings.LastIndex(fileName, ".") + 1 : ]
		validExts := strings.Fields(fl.Param())
		for _, validExt := range validExts {
			if strings.EqualFold(ext, validExt) {
				return true
			}
		}

		return false
	})

	return nil

}