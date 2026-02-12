package site

import (
	"net/http"
)

func PageRoutes(mux *http.ServeMux) {
	// Home (GET / เป็น catch-all ใน Go, ต้องเช็ค path == "/" ด้วย)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			RenderError(w, 404, "ไม่พบหน้าที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง")
			return
		}
		Render(w, "pages/home/index", map[string]any{
			"Title":         "หน้าแรก",
			"NewManga":      MockNewManga(),
			"TrendingManga": MockTrendingManga(),
			"Top10Monthly":  MockTop10Monthly(),
			"Top10Yearly":   MockTop10Yearly(),
			"Top10AllTime":  MockTop10AllTime(),
		})
	})

	// Info
	mux.HandleFunc("GET /info/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		manga := MockMangaByID(id)
		if manga == nil {
			RenderError(w, 404, "ไม่พบมังงะที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง")
			return
		}
		comments := MockCommentsByMangaID(id)
		Render(w, "pages/info/info", map[string]any{
			"Title":    "รายละเอียดมังงะ",
			"MangaID":  id,
			"Manga":    manga,
			"Comments": comments,
		})
	})

	// Search
	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		Render(w, "pages/search/search", map[string]any{
			"Title":   "ค้นหามังงะ",
			"Query":   q,
			"Results": MockSearchManga(q),
		})
	})

	// Categories
	mux.HandleFunc("GET /categories", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "pages/categories/categories", map[string]any{
			"Title":      "หมวดหมู่มังงะ",
			"Categories": MockCategories(),
		})
	})

	// Articles
	mux.HandleFunc("GET /articles", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "pages/articles/articles", map[string]any{
			"Title":    "บทความ",
			"Articles": MockArticles(),
		})
	})

	// Article Detail
	mux.HandleFunc("GET /articles/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		article := MockArticleBySlug(slug)
		if article == nil {
			RenderError(w, 404, "ไม่พบบทความที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง")
			return
		}
		Render(w, "pages/articles/articles-info", map[string]any{
			"Title":   article.Title,
			"Article": article,
		})
	})

	// Community
	mux.HandleFunc("GET /community", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "pages/community/index", map[string]any{
			"Title":   "ชุมชน",
			"Threads": MockThreads(),
		})
	})

	// Community Post Detail
	mux.HandleFunc("GET /community/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		thread := MockThreadByID(id)
		if thread == nil {
			RenderError(w, 404, "ไม่พบกระทู้ที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง")
			return
		}
		Render(w, "pages/community/post-info", map[string]any{
			"Title":  "กระทู้",
			"Thread": thread,
		})
	})

	// About
	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "pages/about/about", map[string]any{"Title": "เกี่ยวกับเรา"})
	})

	// Auth (login + register อยู่ด้วยกัน)
	mux.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "pages/auth/login", map[string]any{"Title": "เข้าสู่ระบบ"})
	})
	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	mux.HandleFunc("GET /register", func(w http.ResponseWriter, r *http.Request) {
		Render(w, "pages/auth/register", map[string]any{"Title": "สมัครสมาชิก"})
	})
	mux.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
}
