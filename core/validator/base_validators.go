package validator

import (
	"fmt"
	"reflect"
	"strings"
	"strconv"

	validate "github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)


func BaseValidator(givenStruct validate.Validator, errors *validate.Errors) {
	// Получаем значение структуры (с разыменовыванием поинтера через Elem())
	structValue := reflect.ValueOf(givenStruct).Elem()
	// получаем кол-во полей в структуре
	structNumFields := structValue.NumField()

	var field reflect.StructField
	var tag string
	// перебираем поля и проверяем теги структуры
	for i:=0; i < structNumFields; i++ {
		// полная информация по полю (название, тип, значение и т.д.)
		field = structValue.Type().Field(i)

		// значение тега validate
		tag = field.Tag.Get("myvalid")
		if tag != "" {
			// перебираем значения тега validate
			for _, tagValue := range strings.Split(tag, "|") {
				switch {
					// обязательное поле
					case tagValue == "required":
						// проверка значения поля на ненулевое
						if structValue.Field(i).String() == "" {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must not be blank", field.Name))
						}

					// валидация email
					case tagValue == "email":
						// валидация средствами библиотеки
						errors.Append(validate.Validate(
							&validators.EmailIsPresent{
								Name: "Email",  // название поля
								Field: structValue.Field(i).String(),  // значение поля
								Message: "Email is not in the right format",  // сообщение при ошибке валидации
							},
						))

					// длина больше чем ... (пример, "min:8")
					case strings.HasPrefix(tagValue, "min"):
						// парсинг минимальной длины из тега
						minInt, _ := strconv.Atoi(strings.TrimPrefix(tagValue, "min:"))

						// проверка типа поля на string и его соответствие длины
						if field.Type == reflect.TypeOf("") && len(structValue.Field(i).String()) < minInt {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must contain at least %d symbols", field.Name, minInt))
						}

					// длина меньше чем ... (пример, "max:100")
					case strings.HasPrefix(tagValue, "max"):
						// парсинг максимальной длины из тега
						maxInt, _ := strconv.Atoi(strings.TrimPrefix(tagValue, "max:"))

						// проверка типа поля на string и его соответствие длины
						if field.Type == reflect.TypeOf("") && len(structValue.Field(i).String()) > maxInt {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must contain less than %d symbols", field.Name, maxInt))
						}
				}
			}
		}		
	}
}
