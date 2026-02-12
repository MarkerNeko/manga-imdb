package utilities

import (
	"html/template"
	"strconv"
	"strings"
)

func Stars(score string) template.HTML {
	f, _ := strconv.ParseFloat(strings.TrimSpace(score), 64)
	if f < 0 {
		f = 0
	}
	if f > 10 {
		f = 10
	}
	rating := f / 2 // 0–5
	full := int(rating)
	remainder := rating - float64(full)
	half := remainder >= 0.25
	empty := 5 - full
	if half {
		empty--
	}
	var sb strings.Builder
	for i := 0; i < full; i++ {
		sb.WriteString(`<span class="star-full">★</span>`)
	}
	if half {
		sb.WriteString(`<span class="star-half">★</span>`)
	}
	for i := 0; i < empty; i++ {
		sb.WriteString(`<span class="star-empty">★</span>`)
	}
	return template.HTML(sb.String())
}
