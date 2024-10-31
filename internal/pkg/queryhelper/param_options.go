package queryhelper

import "strconv"

type ParamOptions struct {
	Limit  int
	Offset int
}

func NewParamOptions() *ParamOptions {
	return &ParamOptions{
		Limit:  10,
		Offset: 0,
	}
}

func (p *ParamOptions) SetLimit(str string) {
	if str != "" {
		parseStr, err := strconv.Atoi(str)
		if err == nil && parseStr > 0 {
			p.Limit = parseStr
		}
	}
}

func (p *ParamOptions) SetOffset(str string) {
	if str != "" {
		parseStr, err := strconv.Atoi(str)
		if err == nil && parseStr > 0 {
			p.Offset = parseStr
		}
	}
}
