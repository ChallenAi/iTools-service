package utils

import (
	// "encoding/json"
	// "github.com/valyala/fasthttp"
	"fmt"
	// "encoding/binary"
	// "github.com/asaskevich/govalidator"
)

type RuleItem struct {
	Type string
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
			errors = append(errors, param + " is required")
			continue
		}

		if p != nil {
			switch ruleItem.Type {
			case "string":
				if IsString(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param + " is invalid")
				}
			case "bool":
				if IsBool(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param + " is invalid")
				}
			case "int":
				if IsInt(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param + " is invalid")
				}
			case "phoneNum":
				if IsString(ruleItem.Type) {
					data[param] = string(peekrable.Peek(param))
				} else {
					errors = append(errors, param + " is invalid")
				}
			default:
				data[param] = string(peekrable.Peek(param))
			}
		}
	}

	fmt.Println(data, errors)
	return data, errors
}



func IsString(p string) bool {
	return true
}

func IsInt(p string) bool {
	return true
}

// func IsPageNum(p interface{}) bool {}

// func IsPerpage(p interface{}) bool {}

// func IsBinary(p interface{}) bool {}

func IsBool(p interface{}) bool {
	return true
}

// func IsIntId(p interface{}) bool {}

func IsPhoneNum(p interface{}) bool {
	return true
}


