package utils

import (
	// "encoding/json"
	// "github.com/valyala/fasthttp"
	"fmt"
	// "encoding/binary"
	// "github.com/asaskevich/govalidator"
)

type RuleItem struct {
	Type     string
	Required bool
}

type Validator struct {
	Rules map[string]RuleItem
}

type Peekrable interface {
	Peek(string) []byte
}

func (v *Validator) Validate(peekrable Peekrable) (map[string]interface{}, []string) {
	errors := []string{}
	data := map[string]interface{}{}
	// map validator params type Mapper
	for param, ruleItem := range v.Rules {

		p := peekrable.Peek(param)
		// fmt.Print(p)

		if ruleItem.Required == true && p == nil {
			errors = append(errors, param+" is required")
			continue
		}

		if p != nil {
			switch ruleItem.Type {
			case "numericId":
				if IsNumericId(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param+" is invalid")
				}
			case "binary":
				if IsBinary(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param+" is invalid")
				}
			case "numeric":
				if IsNumeric(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param+" is invalid")
				}
			case "phoneNum":
				if IsPhoneNum(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param+" is invalid")
				}
			default:
				data[param] = string(peekrable.Peek(param))
			}
		}
	}

	fmt.Println(data, errors)
	return data, errors
}

func IsNumericId(p string) bool {
	return true
}

func IsNumeric(p string) bool {
	return true
}

// func IsPageNum(p interface{}) bool {}

// func IsPerpage(p interface{}) bool {}

func IsBinary(p interface{}) bool {
	return true
}

// func IsNumericId(p interface{}) bool {}

func IsPhoneNum(p interface{}) bool {
	return true
}
