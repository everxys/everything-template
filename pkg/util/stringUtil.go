package util

import "strings"

type FluentStringBuilder struct {
	Builder *strings.Builder
}

func (sb FluentStringBuilder) WriteString(str string) *FluentStringBuilder {
	sb.Builder.WriteString(str)
	return &sb
}

func (sb FluentStringBuilder) String() string {
	return sb.Builder.String()
}

func NewFluentBuilder() *FluentStringBuilder {
	var builder strings.Builder
	return &FluentStringBuilder{Builder: &builder}
}
