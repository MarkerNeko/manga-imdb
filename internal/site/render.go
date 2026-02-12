package site

import (
	"html/template"
	"manga-imdb/internal/utilities"
	"net/http"
	"path/filepath"
	"strconv"
)

// stars แปลงคะแนน 0–10 เป็นดาว 5 ดวง
// ★ เต็ม (สีทอง), ★ ครึ่ง (ซ้ายทอง ขวาเทา), ★ ว่าง (สีเทา)

func Render(w http.ResponseWriter, page string, data any) {
	t := template.New("root").Funcs(template.FuncMap{
		"add":      func(a, b int) int { return a + b },
		"sub":      func(a, b int) int { return a - b },
		"mul":      func(a, b int) int { return a * b },
		"div":      func(a, b int) int { return a / b },
		"stars":    utilities.Stars,
		"starsInt": func(i int) template.HTML { return utilities.Stars(strconv.Itoa(i)) },
	})
	template.Must(t.ParseGlob("web/html/layouts/*.html"))
	template.Must(t.ParseGlob("web/html/components/*.html"))
	// template.Must(t.ParseGlob("web/html/components/**/*.html"))

	// TODO: ดึง User จริงจาก session/cookie — ตอนนี้ mock ไว้เพื่อทดสอบ
	// ตั้ง mockLoggedIn = true เพื่อดู avatar, false เพื่อดูปุ่ม Sign in/Sign up
	mockLoggedIn := false
	var user map[string]any
	if mockLoggedIn {
		user = map[string]any{
			"Username":  "OzarkK",
			"AvatarURL": "https://api.dicebear.com/9.x/thumbs/svg?seed=OzarkK",
		}
	}

	merged := map[string]any{
		"Website": "Manga IMDB",
		"URL":     "imdb.com",
		"User":    user,
	}

	// 🧩 merge field จาก data เข้า merged โดยตรง (flatten)
	if base, ok := data.(map[string]any); ok {
		for k, v := range base {
			merged[k] = v
		}
	}

	// 2) เฉพาะเพจที่เรียกใช้ (เช่น "pages/index")
	file := filepath.FromSlash("web/html/" + page + ".html")
	template.Must(t.ParseFiles(file))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// layout จะเรียก block "title"/"content" ที่เพจ define
	_ = t.ExecuteTemplate(w, "layouts/default", merged)
}

// RenderError renders the error page with the given HTTP status code.
func RenderError(w http.ResponseWriter, status int, message string) {
	Render404(w, status, map[string]any{
		"Title":      http.StatusText(status),
		"StatusCode": status,
		"Message":    message,
	})
}

func Render404(w http.ResponseWriter, status int, data map[string]any) {
	t := template.New("root").Funcs(template.FuncMap{
		"add":      func(a, b int) int { return a + b },
		"sub":      func(a, b int) int { return a - b },
		"mul":      func(a, b int) int { return a * b },
		"div":      func(a, b int) int { return a / b },
		"stars":    utilities.Stars,
		"starsInt": func(i int) template.HTML { return utilities.Stars(strconv.Itoa(i)) },
	})
	template.Must(t.ParseGlob("web/html/layouts/*.html"))
	template.Must(t.ParseGlob("web/html/components/*.html"))

	file := filepath.FromSlash("web/html/pages/error/error.html")
	template.Must(t.ParseFiles(file))

	merged := map[string]any{
		"Website": "Manga IMDB",
		"URL":     "imdb.com",
	}
	for k, v := range data {
		merged[k] = v
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_ = t.ExecuteTemplate(w, "layouts/default", merged)
}
