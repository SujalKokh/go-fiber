package responses

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// JSON : json response function
func JSON(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{"data": data})
}

// ErrorJSON : json error response function
func ErrorJSON(c *fiber.Ctx, statusCode int, err error) error {
	var errorObject = make(map[string]string)
	trans := Validator()
	if val, ok := err.(validator.ValidationErrors); ok {
		errMap := val.Translate(trans)
		for _, v := range errMap {
			errorObject["message"] = v
			break
		}
		return c.Status(statusCode).JSON(fiber.Map{"error": errorObject})

	} else {
		var raw interface{}

		in, _ := json.Marshal(err)

		if len(in) > 5 {
			if err1 := json.Unmarshal(in, &raw); err1 != nil {
				if e, ok := err1.(*json.SyntaxError); ok {
					fmt.Printf("syntax error at byte offset %d", e.Offset)
				}
				fmt.Printf("sakura response: %q", in)
			}
			return c.Status(statusCode).JSON(fiber.Map{"error": raw, "message": err.Error()})
		} else {
			errorObject["message"] = err.Error()
			return c.Status(statusCode).JSON(fiber.Map{"error": errorObject})
		}

	}
}

// SuccessJSON : json error response function
func SuccessJSON(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{"msg": data})
}

// JSONCount : json response function
func JSONCount(c *fiber.Ctx, statusCode int, data interface{}, count int) error {
	return c.Status(statusCode).JSON(fiber.Map{"data": data, "count": count})
}

func NoContent(c *fiber.Ctx) error {
	return c.Send(nil)
}

func Validator() ut.Translator {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} is a required field", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})

		_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
			return ut.Add("email", "{0} must be a valid email", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("email", fe.Field())
			return t
		})
	}
	return trans
}
