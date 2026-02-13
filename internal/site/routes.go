package site

import (
	"net/http"
)

func PageRoutes(mux *http.ServeMux) {

	// ===== Root redirect =====
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/th/", http.StatusFound)
			return
		}
		// URL ที่ไม่ match route ใดเลย
		RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
	})

	// ===== Home =====
	mux.HandleFunc("GET /{lang}/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		// catch-all: /{lang}/ เป็น subtree — เช็คว่าเป็น home จริงๆ
		if r.URL.Path != "/"+lang+"/" {
			RenderError(w, lang, 404, GetTranslations(lang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		Render(w, lang, "pages/home/index", map[string]any{
			"Title":         t["nav_home"],
			"NewManga":      MockNewManga(),
			"TrendingManga": MockTrendingManga(),
			"Top10Monthly":  MockTop10Monthly(),
			"Top10Yearly":   MockTop10Yearly(),
			"Top10AllTime":  MockTop10AllTime(),
		})
	})

	// ===== Info =====
	mux.HandleFunc("GET /{lang}/info/{id}", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		id := r.PathValue("id")
		manga := MockMangaByID(id)
		if manga == nil {
			RenderError(w, lang, 404, t["err_manga"])
			return
		}
		comments := MockCommentsByMangaID(id)
		Render(w, lang, "pages/info/info", map[string]any{
			"Title":    manga.Title,
			"MangaID":  id,
			"Manga":    manga,
			"Comments": comments,
		})
	})

	// ===== Search =====
	mux.HandleFunc("GET /{lang}/search", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		q := r.URL.Query().Get("q")
		Render(w, lang, "pages/search/search", map[string]any{
			"Title":   t["nav_search"],
			"Query":   q,
			"Results": MockSearchManga(q),
		})
	})

	// ===== Categories =====
	mux.HandleFunc("GET /{lang}/categories", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		Render(w, lang, "pages/categories/categories", map[string]any{
			"Title":      t["categories_title"],
			"Categories": MockCategories(),
		})
	})

	// ===== Articles =====
	mux.HandleFunc("GET /{lang}/articles", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		Render(w, lang, "pages/articles/articles", map[string]any{
			"Title":    t["articles_title"],
			"Articles": MockArticles(),
		})
	})

	// ===== Article Detail =====
	mux.HandleFunc("GET /{lang}/articles/{slug}", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		slug := r.PathValue("slug")
		article := MockArticleBySlug(slug)
		if article == nil {
			RenderError(w, lang, 404, t["err_article"])
			return
		}
		Render(w, lang, "pages/articles/articles-info", map[string]any{
			"Title":   article.Title,
			"Article": article,
		})
	})

	// ===== Community =====
	mux.HandleFunc("GET /{lang}/community", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		Render(w, lang, "pages/community/index", map[string]any{
			"Title":   t["nav_community"],
			"Threads": MockThreads(),
		})
	})

	// ===== Community Post Detail =====
	mux.HandleFunc("GET /{lang}/community/{id}", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		id := r.PathValue("id")
		thread := MockThreadByID(id)
		if thread == nil {
			RenderError(w, lang, 404, t["err_thread"])
			return
		}
		Render(w, lang, "pages/community/post-info", map[string]any{
			"Title":  thread.Title,
			"Thread": thread,
		})
	})

	// ===== About =====
	mux.HandleFunc("GET /{lang}/about", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		Render(w, lang, "pages/about/about", map[string]any{"Title": t["about_title"]})
	})

	// ===== Auth =====
	mux.HandleFunc("GET /{lang}/login", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		Render(w, lang, "pages/auth/login", map[string]any{"Title": t["login_title"]})
	})
	mux.HandleFunc("POST /{lang}/login", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			lang = DefaultLang
		}
		http.Redirect(w, r, "/"+lang+"/", http.StatusSeeOther)
	})
	mux.HandleFunc("GET /{lang}/register", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			RenderError(w, DefaultLang, 404, GetTranslations(DefaultLang)["err_404"])
			return
		}
		t := GetTranslations(lang)
		Render(w, lang, "pages/auth/register", map[string]any{"Title": t["register_title"]})
	})
	mux.HandleFunc("POST /{lang}/register", func(w http.ResponseWriter, r *http.Request) {
		lang := r.PathValue("lang")
		if !ValidLang(lang) {
			lang = DefaultLang
		}
		http.Redirect(w, r, "/"+lang+"/login", http.StatusSeeOther)
	})
}
