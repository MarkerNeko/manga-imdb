package site

import (
	"html/template"
	"manga-imdb/internal/utilities"
	"net/http"
	"path/filepath"
	"strconv"
)

func Render(w http.ResponseWriter, lang string, page string, data any) {
	trans := GetTranslations(lang)

	t := template.New("root").Funcs(template.FuncMap{
		"add":        func(a, b int) int { return a + b },
		"sub":        func(a, b int) int { return a - b },
		"mul":        func(a, b int) int { return a * b },
		"div":        func(a, b int) int { return a / b },
		"stars":      utilities.Stars,
		"starsInt":   func(i int) template.HTML { return utilities.Stars(strconv.Itoa(i)) },
		"langPrefix": func() string { return "/" + lang },
		"t":          func(key string) string { return trans[key] },
	})
	template.Must(t.ParseGlob("web/html/layouts/*.html"))
	template.Must(t.ParseGlob("web/html/components/*.html"))

	// TODO: ดึง User จริงจาก session/cookie — ตอนนี้ mock ไว้เพื่อทดสอบ
	mockLoggedIn := false
	var user map[string]any
	if mockLoggedIn {
		user = map[string]any{
			"Username":  "OzarkK",
			"AvatarURL": "https://api.dicebear.com/9.x/thumbs/svg?seed=OzarkK",
		}
	}

	merged := map[string]any{
		"Website":    "Manga IMDB",
		"URL":        "imdb.com",
		"User":       user,
		"Lang":       lang,
		"LangPrefix": "/" + lang,
		"T":          trans,
	}

	if base, ok := data.(map[string]any); ok {
		for k, v := range base {
			merged[k] = v
		}
	}

	file := filepath.FromSlash("web/html/" + page + ".html")
	template.Must(t.ParseFiles(file))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = t.ExecuteTemplate(w, "layouts/default", merged)
}

// RenderError renders the error page with the given HTTP status code.
func RenderError(w http.ResponseWriter, lang string, status int, message string) {
	trans := GetTranslations(lang)

	t := template.New("root").Funcs(template.FuncMap{
		"add":        func(a, b int) int { return a + b },
		"sub":        func(a, b int) int { return a - b },
		"mul":        func(a, b int) int { return a * b },
		"div":        func(a, b int) int { return a / b },
		"stars":      utilities.Stars,
		"starsInt":   func(i int) template.HTML { return utilities.Stars(strconv.Itoa(i)) },
		"langPrefix": func() string { return "/" + lang },
		"t":          func(key string) string { return trans[key] },
	})
	template.Must(t.ParseGlob("web/html/layouts/*.html"))
	template.Must(t.ParseGlob("web/html/components/*.html"))

	file := filepath.FromSlash("web/html/pages/error/error.html")
	template.Must(t.ParseFiles(file))

	merged := map[string]any{
		"Website":    "Manga IMDB",
		"URL":        "imdb.com",
		"Lang":       lang,
		"LangPrefix": "/" + lang,
		"T":          trans,
		"Title":      http.StatusText(status),
		"StatusCode": status,
		"Message":    message,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_ = t.ExecuteTemplate(w, "layouts/default", merged)
}
