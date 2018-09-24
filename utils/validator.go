package utils

import (
	// "encoding/json"
	// "github.com/valyala/fasthttp"
	"fmt"
	"strconv"
	// "encoding/binary"
	"github.com/asaskevich/govalidator"
)

type RuleItem struct {
	Type     string
	Required bool
	Alias    string
}

type Validator struct {
	Rules map[string]RuleItem
}

type Peekrable interface {
	Peek(string) []byte
}

type ServiceParams struct {
	CommonParams map[string]interface{}
	LikeQuery    string
	Limit        int
	Offset       int
}

func (v *Validator) Validate(peekrable Peekrable) (*ServiceParams, []string) {
	var (
		pageNumber int
		pageSize   int
		errors     []string
		p          string
	)

	data := &ServiceParams{
		CommonParams: map[string]interface{}{},
		Limit:        10,
		Offset:       0,
	}
	// map validator params type Mapper
	for param, ruleItem := range v.Rules {

		p = string(peekrable.Peek(param))
		// fmt.Println(data)
		// fmt.Println(param, string(peekrable.Peek(param)))

		if ruleItem.Required == true && p == "" {
			errors = append(errors, param+" is required")
			continue
		}

		if p != "" {
			switch ruleItem.Type {
			case "pageNumber":
				if IsPageNumber(p) {
					pageNumber, _ = strconv.Atoi(p)
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "pageSize":
				if IsPageSize(p) {
					pageSize, _ = strconv.Atoi(string(p))
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "likeQuery":
				data.LikeQuery = string(p)
				break
			case "binary":
				if IsBinary(p) {
					if p == "1" {
						if alias := ruleItem.Alias; alias != "" {
							data.CommonParams[alias] = true
						} else {
							data.CommonParams[param] = true
						}
					} else {
						if alias := ruleItem.Alias; alias != "" {
							data.CommonParams[alias] = false
						} else {
							data.CommonParams[param] = false
						}
					}
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "number":
				if IsNumber(p) {
					if alias := ruleItem.Alias; alias != "" {
						data.CommonParams[alias] = p
					} else {
						data.CommonParams[param] = p
					}
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "phoneNum":
				if IsPhoneNum(p) {
					if alias := ruleItem.Alias; alias != "" {
						data.CommonParams[alias] = p
					} else {
						data.CommonParams[param] = p
					}
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			}
		}
	}

	if pageSize != 0 {
		data.Limit = pageSize
	} else {
		data.Limit = 10
	}

	if pageNumber != 0 {
		data.Offset = data.Limit * (pageNumber - 1)
	} else {
		data.Offset = 0
	}

	// fmt.Println(data, errors)
	fmt.Println("")
	return data, errors
}

func IsNumber(p string) bool {
	return govalidator.IsInt(p)
}

func IsPageNumber(p string) bool {
	return govalidator.IsInt(p)
}

func IsPageSize(p string) bool {
	return govalidator.IsInt(p)
}

func IsBinary(p string) bool {
	return govalidator.IsInt(p)
}

func IsPhoneNum(p string) bool {
	return govalidator.IsNumeric(p)
}
