package site

// Translations เก็บข้อความแปลตาม key
type Translations map[string]string

// SupportedLangs — ภาษาที่รองรับ
var SupportedLangs = map[string]bool{"th": true, "en": true}

// DefaultLang — ภาษาเริ่มต้น
const DefaultLang = "th"

// AllTranslations — เก็บข้อความทุกภาษา
var AllTranslations = map[string]Translations{
	"th": {
		// ===== Navbar =====
		"nav_home":       "หน้าแรก",
		"nav_categories": "หมวดหมู่",
		"nav_articles":   "บทความ",
		"nav_community":  "ชุมชน",
		"nav_about":      "เกี่ยวกับเรา",
		"nav_search":     "ค้นหา",
		"sign_in":        "Sign in",
		"sign_up":        "Sign up",
		"my_profile":     "โปรไฟล์ของฉัน",
		"favorites":      "รายการโปรด",
		"logout":         "ออกจากระบบ",

		// ===== Home =====
		"home_subtitle":  "มังงะใหม่ มาแรง และเรตติ้งสูงสุด",
		"new_manga":      "มังงะใหม่",
		"trending_manga": "มังงะมาแรง",
		"top10_rating":   "Top 10 เรตติ้ง",
		"tab_monthly":    "ประจำเดือน",
		"tab_yearly":     "ประจำปี",
		"tab_alltime":    "ตลอดกาล",
		"scroll_left":    "เลื่อนซ้าย",
		"scroll_right":   "เลื่อนขวา",

		// ===== Info =====
		"synopsis":       "เรื่องย่อ",
		"reviews":        "รีวิว / คอมเมนต์",
		"login_to_review": "ต้องเข้าสู่ระบบเพื่อให้คะแนนและคอมเมนต์",
		"no_comments":    "ยังไม่มีคอมเมนต์",
		"manga_not_found": "ไม่พบมังงะ",
		"review_suffix":  "รีวิว",

		// ===== Search =====
		"search_placeholder": "ค้นหาชื่อมังงะ...",
		"search_btn":         "ค้นหา",
		"search_results_for": "ผลค้นหา",
		"no_results":         "ไม่พบผลลัพธ์",
		"search_hint":        "พิมพ์คำค้นด้านบนเพื่อค้นหามังงะ",

		// ===== Categories =====
		"categories_title":    "หมวดหมู่มังงะ",
		"categories_subtitle": "เลือกหมวดหมู่ที่ชอบ แล้วค้นพบมังงะเรื่องใหม่ๆ",
		"count_suffix":        "เรื่อง",

		// ===== Articles =====
		"articles_title":    "บทความ",
		"articles_subtitle": "ข่าวสาร รีวิว และเรื่องน่ารู้จากโลกมังงะ",
		"back_to_articles":  "กลับไปหน้าบทความ",
		"article_not_found": "ไม่พบบทความนี้",
		"article_mock_body": "นี่คือเนื้อหาตัวอย่างของบทความ เมื่อเชื่อม DB จริงจะแสดงเนื้อหาเต็ม",

		// ===== Community =====
		"community_title":    "กระทู้พูดคุย",
		"community_subtitle": "แลกเปลี่ยนความเห็น พูดคุยเรื่องมังงะกับเพื่อนๆ",
		"replies":            "ตอบกลับ",
		"back_to_community":  "กลับไปหน้าชุมชน",
		"thread_not_found":   "ไม่พบกระทู้นี้",
		"thread_mock_body":   "นี่คือเนื้อหาตัวอย่างของกระทู้ เมื่อเชื่อม DB จริงจะแสดงเนื้อหาเต็ม",
		"comments_heading":   "ความคิดเห็น",
		"no_comments_yet":    "ยังไม่มีความคิดเห็น — เชื่อม DB แล้วจะแสดงข้อมูลจริง",

		// ===== About =====
		"about_title":       "เกี่ยวกับเรา",
		"about_subtitle":    "แหล่งรวมรีวิวมังงะจากผู้อ่านจริง",
		"about_who":         "เราคือใคร?",
		"about_who_text":    "คือเว็บไซต์รีวิวมังงะที่สร้างขึ้นเพื่อชุมชนคนรักมังงะ เราเชื่อว่าทุกคนควรมีพื้นที่แบ่งปันความรู้สึกหลังอ่านมังงะจบ ไม่ว่าจะเป็นเรื่องที่ชอบหรือไม่ชอบ ทุกความเห็นมีค่า",
		"about_what":        "สิ่งที่เราทำ",
		"about_what_text":   "เรารวบรวมข้อมูลมังงะจากทั่วโลก ให้คุณค้นหา อ่านรีวิว ให้คะแนน และติดตามอันดับมังงะยอดนิยมที่อัปเดตทุกเดือน",
		"stat_manga":        "มังงะในระบบ",
		"stat_reviews":      "รีวิวจากผู้อ่าน",
		"stat_members":      "สมาชิก",
		"stat_categories":   "หมวดหมู่",
		"about_why":         "ทำไมต้อง",
		"feature_review":    "รีวิวจากผู้อ่านจริง",
		"feature_review_d":  "คะแนนและความเห็นจากคนอ่านจริงๆ ไม่มีบอท ไม่มีรีวิวปลอม",
		"feature_ranking":   "อันดับอัปเดตทุกเดือน",
		"feature_ranking_d": "Top 10 ประจำเดือน ประจำปี และตลอดกาล คำนวณจากรีวิวจริง",
		"feature_search":    "ค้นหาง่าย",
		"feature_search_d":  "ค้นหาจากชื่อ หมวดหมู่ หรือเรียงตามคะแนน หาเรื่องอ่านใหม่ได้เร็ว",
		"feature_community": "ชุมชนคนรักมังงะ",
		"feature_community_d": "พูดคุย แลกเปลี่ยน แนะนำมังงะกับเพื่อนๆ ในกระทู้",

		// ===== Auth =====
		"login_title":          "เข้าสู่ระบบ",
		"login_email_label":    "อีเมล หรือชื่อผู้ใช้",
		"login_password_label": "รหัสผ่าน",
		"login_btn":            "เข้าสู่ระบบ",
		"login_no_account":     "ยังไม่มีบัญชี?",
		"register_title":       "สมัครสมาชิก",
		"register_username":    "ชื่อผู้ใช้",
		"register_email":       "อีเมล",
		"register_password":    "รหัสผ่าน",
		"register_confirm":     "ยืนยันรหัสผ่าน",
		"register_btn":         "สมัคร",
		"register_has_account": "มีบัญชีอยู่แล้ว?",

		// ===== Footer =====
		"footer_desc":        "แหล่งรวมรีวิวมังงะที่ใหญ่ที่สุด ค้นหา ให้คะแนน และแชร์ความเห็นเกี่ยวกับมังงะเรื่องโปรดของคุณ อ่านรีวิวจากผู้อ่านจริง พร้อมอันดับมังงะยอดนิยมอัปเดตทุกเดือน",
		"footer_explore":     "สำรวจ",
		"footer_new_manga":   "มังงะใหม่ล่าสุด",
		"footer_hot_manga":   "มังงะมาแรงประจำสัปดาห์",
		"footer_top10":       "Top 10 มังงะเรตติ้งสูงสุด",
		"footer_search":      "ค้นหามังงะ",
		"footer_popular_cat": "หมวดหมู่ยอดนิยม",
		"footer_action":      "แอคชัน",
		"footer_romance":     "โรแมนซ์",
		"footer_fantasy":     "แฟนตาซี",
		"footer_comedy":      "คอมเมดี้",
		"footer_horror":      "สยองขวัญ",
		"footer_account":     "บัญชีผู้ใช้",
		"footer_seo":         "เว็บไซต์รีวิวมังงะออนไลน์ รวมข้อมูลมังงะยอดนิยมกว่าหลายพันเรื่อง ทั้ง One Piece, Demon Slayer, Jujutsu Kaisen, Chainsaw Man, Attack on Titan, Spy x Family และอีกมากมาย ดูคะแนนรีวิวจากผู้อ่านจริง เรื่องย่อ หมวดหมู่ และอันดับมังงะยอดนิยมประจำเดือน ประจำปี และตลอดกาล",
		"footer_copyright":   "สร้างด้วยความรักต่อมังงะ",

		// ===== Error =====
		"back_home":     "กลับหน้าแรก",
		"err_404":       "ไม่พบหน้าที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง",
		"err_manga":     "ไม่พบมังงะที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง",
		"err_article":   "ไม่พบบทความที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง",
		"err_thread":    "ไม่พบกระทู้ที่คุณต้องการ กรุณาตรวจสอบ URL อีกครั้ง",
	},

	"en": {
		// ===== Navbar =====
		"nav_home":       "Home",
		"nav_categories": "Categories",
		"nav_articles":   "Articles",
		"nav_community":  "Community",
		"nav_about":      "About Us",
		"nav_search":     "Search",
		"sign_in":        "Sign in",
		"sign_up":        "Sign up",
		"my_profile":     "My Profile",
		"favorites":      "Favorites",
		"logout":         "Logout",

		// ===== Home =====
		"home_subtitle":  "New, Trending & Top Rated Manga",
		"new_manga":      "New Manga",
		"trending_manga": "Trending Manga",
		"top10_rating":   "Top 10 Ratings",
		"tab_monthly":    "Monthly",
		"tab_yearly":     "Yearly",
		"tab_alltime":    "All Time",
		"scroll_left":    "Scroll left",
		"scroll_right":   "Scroll right",

		// ===== Info =====
		"synopsis":       "Synopsis",
		"reviews":        "Reviews / Comments",
		"login_to_review": "Log in to rate and comment",
		"no_comments":    "No comments yet",
		"manga_not_found": "Manga not found",
		"review_suffix":  "reviews",

		// ===== Search =====
		"search_placeholder": "Search manga by title...",
		"search_btn":         "Search",
		"search_results_for": "Results for",
		"no_results":         "No results found",
		"search_hint":        "Type a keyword above to search for manga",

		// ===== Categories =====
		"categories_title":    "Manga Categories",
		"categories_subtitle": "Pick a category and discover new manga",
		"count_suffix":        "titles",

		// ===== Articles =====
		"articles_title":    "Articles",
		"articles_subtitle": "News, reviews and interesting stories from the manga world",
		"back_to_articles":  "Back to articles",
		"article_not_found": "Article not found",
		"article_mock_body": "This is sample article content. Full content will appear once a database is connected.",

		// ===== Community =====
		"community_title":    "Discussion Threads",
		"community_subtitle": "Share opinions and talk about manga with fellow fans",
		"replies":            "replies",
		"back_to_community":  "Back to community",
		"thread_not_found":   "Thread not found",
		"thread_mock_body":   "This is sample thread content. Full content will appear once a database is connected.",
		"comments_heading":   "Comments",
		"no_comments_yet":    "No comments yet — will show real data once DB is connected",

		// ===== About =====
		"about_title":       "About Us",
		"about_subtitle":    "Real manga reviews from real readers",
		"about_who":         "Who are we?",
		"about_who_text":    "is a manga review website built for the manga-loving community. We believe everyone deserves a space to share how they feel after finishing a manga — whether they loved it or not, every opinion matters.",
		"about_what":        "What we do",
		"about_what_text":   "We gather manga data from around the world so you can search, read reviews, rate, and follow the most popular manga rankings updated every month.",
		"stat_manga":        "Manga titles",
		"stat_reviews":      "Reader reviews",
		"stat_members":      "Members",
		"stat_categories":   "Categories",
		"about_why":         "Why choose",
		"feature_review":    "Real reader reviews",
		"feature_review_d":  "Scores and comments from real readers. No bots, no fake reviews.",
		"feature_ranking":   "Rankings updated monthly",
		"feature_ranking_d": "Top 10 monthly, yearly, and all-time, calculated from real reviews.",
		"feature_search":    "Easy search",
		"feature_search_d":  "Search by name, category, or sort by rating. Find new reads quickly.",
		"feature_community": "Manga community",
		"feature_community_d": "Chat, share, and recommend manga with friends in threads.",

		// ===== Auth =====
		"login_title":          "Sign In",
		"login_email_label":    "Email or username",
		"login_password_label": "Password",
		"login_btn":            "Sign In",
		"login_no_account":     "Don't have an account?",
		"register_title":       "Register",
		"register_username":    "Username",
		"register_email":       "Email",
		"register_password":    "Password",
		"register_confirm":     "Confirm password",
		"register_btn":         "Register",
		"register_has_account": "Already have an account?",

		// ===== Footer =====
		"footer_desc":        "The largest manga review community. Search, rate, and share your thoughts about your favorite manga. Read real reviews and discover the most popular titles updated monthly.",
		"footer_explore":     "Explore",
		"footer_new_manga":   "Latest Manga",
		"footer_hot_manga":   "Trending This Week",
		"footer_top10":       "Top 10 Rated Manga",
		"footer_search":      "Search Manga",
		"footer_popular_cat": "Popular Categories",
		"footer_action":      "Action",
		"footer_romance":     "Romance",
		"footer_fantasy":     "Fantasy",
		"footer_comedy":      "Comedy",
		"footer_horror":      "Horror",
		"footer_account":     "Account",
		"footer_seo":         "Manga IMDB — Online manga review website. A database of thousands of popular manga including One Piece, Demon Slayer, Jujutsu Kaisen, Chainsaw Man, Attack on Titan, Spy x Family, and many more. Read real reader reviews, synopses, categories, and top manga rankings updated monthly, yearly, and all-time.",
		"footer_copyright":   "Built with love for manga",

		// ===== Error =====
		"back_home":     "Back to Home",
		"err_404":       "The page you're looking for doesn't exist. Please check the URL.",
		"err_manga":     "The manga you're looking for doesn't exist. Please check the URL.",
		"err_article":   "The article you're looking for doesn't exist. Please check the URL.",
		"err_thread":    "The thread you're looking for doesn't exist. Please check the URL.",
	},
}

// GetTranslations คืนค่า Translations ตามภาษา (fallback เป็น th)
func GetTranslations(lang string) Translations {
	if t, ok := AllTranslations[lang]; ok {
		return t
	}
	return AllTranslations[DefaultLang]
}

// ValidLang ตรวจสอบภาษาที่รองรับ
func ValidLang(lang string) bool {
	return SupportedLangs[lang]
}
