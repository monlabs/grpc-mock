package stub

import (
	"errors"
	"reflect"
	"regexp"

	log "github.com/golang/glog"
)

type Stub struct {
	Service string  `json:"service"`
	Method  string  `json:"method"`
	In      *Input  `json:"in"`
	Out     *Output `json:"out"`
}

type Input struct {
	Equals   map[string]interface{} `json:"equals"`
	Contains map[string]interface{} `json:"contains"`
	Matches  map[string]interface{} `json:"matches"`
}

type Output struct {
	Data    map[string]interface{} `json:"data"`
	Code    int32                  `json:"code"`
	Message string                 `json:"message"`
}

func (s *Stub) Validate() error {
	if s.Service == "" {
		return errors.New("missing service")
	}
	if s.Method == "" {
		return errors.New("missing method")
	}
	if s.In != nil && len(s.In.Equals) == 0 && len(s.In.Contains) == 0 && len(s.In.Matches) == 0 {
		return errors.New("require at least one of equals, contains or matches")
	}
	if s.Out == nil {
		return errors.New("missing output")
	}
	return nil
}

func (s *Stub) Match(in map[string]interface{}) bool {
	if s.In == nil {
		return true
	}

	if s.In.Equals != nil {
		return equals(s.In.Equals, in)
	}

	if s.In.Contains != nil {
		return contains(s.In.Contains, in)
	}

	if s.In.Matches != nil {
		return matches(s.In.Matches, in)
	}

	return false
}

func equals(pattern, in map[string]interface{}) bool {
	return reflect.DeepEqual(pattern, in)
}

func contains(pattern, in map[string]interface{}) bool {
	for k, v := range in {
		p := pattern[k]
		if p == nil || !reflect.DeepEqual(p, v) {
			return false
		}
	}
	return true
}

func matches(pattern, in map[string]interface{}) bool {
	for k, v := range in {
		valStr, ok := v.(string)
		if !ok {
			return false
		}

		pStr, ok := pattern[k].(string)
		if !ok {
			return false
		}

		match, err := regexp.Match(pStr, []byte(valStr))
		if err != nil {
			log.Errorf("match regexp '%s' with '%s' failed: %v", pStr, valStr, err)
		}

		if !match {
			return false
		}
	}
	return true
}
