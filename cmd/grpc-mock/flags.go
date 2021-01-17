package main

import "strings"

type StringsValue struct {
	Delimiter string
	Elements  []string
}

func (v *StringsValue) String() string {
	return strings.Join(v.Elements, v.Delimiter)
}

func (v *StringsValue) Set(s string) error {
	parts := strings.Split(s, v.Delimiter)
	for _, p := range parts {
		v.Elements = append(v.Elements, strings.TrimSpace(p))
	}
	return nil
}
