package site

import "strings"

type MangaCard struct {
	ID          string
	Title       string
	CoverURL    string
	Score       string
	ReviewCount int
}

type MangaDetail struct {
	ID          string
	Title       string
	CoverURL    string
	Score       string
	ReviewCount int
	Synopsis    string
	Categories  []string
}

type Comment struct {
	Username string
	Score    int
	Comment  string
}

var mockMangaList = []MangaCard{
	{ID: "1", Title: "One Piece", CoverURL: "https://picsum.photos/seed/op/280/400", Score: "9.2", ReviewCount: 12450},
	{ID: "2", Title: "Demon Slayer", CoverURL: "https://picsum.photos/seed/ds/280/400", Score: "8.9", ReviewCount: 8920},
	{ID: "3", Title: "Jujutsu Kaisen", CoverURL: "https://picsum.photos/seed/jk/280/400", Score: "8.7", ReviewCount: 7650},
	{ID: "4", Title: "Chainsaw Man", CoverURL: "https://picsum.photos/seed/csm/280/400", Score: "8.8", ReviewCount: 5430},
	{ID: "5", Title: "Spy x Family", CoverURL: "https://picsum.photos/seed/sxf/280/400", Score: "8.6", ReviewCount: 4320},
	{ID: "6", Title: "My Hero Academia", CoverURL: "https://picsum.photos/seed/mha/280/400", Score: "8.4", ReviewCount: 9870},
	{ID: "7", Title: "Attack on Titan", CoverURL: "https://picsum.photos/seed/aot/280/400", Score: "9.0", ReviewCount: 11200},
	{ID: "8", Title: "Solo Leveling", CoverURL: "https://picsum.photos/seed/sl/280/400", Score: "8.5", ReviewCount: 6780},
	{ID: "9", Title: "Blue Lock", CoverURL: "https://picsum.photos/seed/bl/280/400", Score: "8.3", ReviewCount: 3210},
	{ID: "10", Title: "Oshi no Ko", CoverURL: "https://picsum.photos/seed/onk/280/400", Score: "8.7", ReviewCount: 4560},
	{ID: "11", Title: "Frieren", CoverURL: "https://picsum.photos/seed/fr/280/400", Score: "9.1", ReviewCount: 2890},
	{ID: "12", Title: "Dandadan", CoverURL: "https://picsum.photos/seed/ddd/280/400", Score: "8.6", ReviewCount: 2100},
	{ID: "13", Title: "Dandadan22222", CoverURL: "https://picsum.photos/seed/ddd/280/400", Score: "8.6", ReviewCount: 2100},
}

var mockSynopsis = map[string]string{
	"1":  "เรื่องราวของมังกี้ ดี. ลูฟี่ ที่ออกเดินทางเพื่อเป็นราชาโจรสลัด และตามหาสมบัติ One Piece ที่ปลายแกรนด์ไลน์",
	"2":  "ทานจิโร่ต้องต่อสู้กับอสูรและช่วยน้องสาวที่กลายเป็นอสูร ให้กลับมาเป็นมนุษย์อีกครั้ง",
	"3":  "ยูจิ ผู้ชายที่กลืนนิ้วของสุคุนะเข้าไป ต้องร่วมมือกับโรงเรียนเวทมนตร์เพื่อล่าอสูรและกินนิ้วที่เหลือ",
	"4":  "เด็นจิอยากใช้ชีวิตธรรมดา แต่กลับได้พลัง Chainsaw Devil และต้องทำงานให้หน่วยล่าอสูร",
	"5":  "สายลับต้องสร้างครอบครัวปลอมเพื่อปฏิบัติภารกิจ แต่ลูกสาวที่รับมาเป็นมือสังหาร ส่วนภรรยาเป็นมือปราบ",
	"6":  "เด็กชายที่เกิดมาไม่มีพลังในโลกที่คนมีพลังเหนือธรรมชาติ พยายามจะเป็นฮีโร่ให้ได้",
	"7":  "มนุษยชาติถูกปิดอยู่ในกำแพง ต้องต่อสู้กับไททันยักษ์ที่กินคน",
	"8":  "ฮันเตอร์ที่อ่อนแอที่สุดในโลก ได้ระบบที่ทำให้เขาสามารถเลเวลอัพและแข็งแกร่งขึ้นเรื่อยๆ",
	"9":  "โปรเจกต์ Blue Lock รวบรวมนักเตะฝีมือดีมาปรับ mindset ให้เป็นศูนย์กลางและทำประตู",
	"10": "idol ลึกลับกับหมอที่ถูกฆ่า เกิดใหม่เป็นลูกของ idol คนนั้น และตามหาเรื่องราวจริงของวงการ",
	"11": "แม่มดเอลฟ์ที่อยู่กับฮีโร่กลุ่มหนึ่งเดินทางไปส่งร่างเพื่อนที่จากไป ค่อยๆ เรียนรู้ความหมายของเวลาและความผูกพัน",
	"12": "เด็กสาวที่เชื่อเรื่องยูเอฟโอ กับเด็กชายที่เชื่อเรื่องผี ต้องช่วยกันตามหาอวัยวะที่ถูกยูเอฟโอขโมยไป",
}

var mockCategories = map[string][]string{
	"1":  {"แอคชัน", "ผจญภัย", "แฟนตาซี"},
	"2":  {"แอคชัน", "โชเน็น", "เหนือธรรมชาติ"},
	"3":  {"แอคชัน", "โชเน็น", "เหนือธรรมชาติ"},
	"4":  {"แอคชัน", "โชเน็น", "สยองขวัญ"},
	"5":  {"คอมเมดี้", "สายลับ", "ชีวิตประจำวัน"},
	"6":  {"แอคชัน", "โชเน็น", "ซูเปอร์ฮีโร่"},
	"7":  {"แอคชัน", "ดาร์กแฟนตาซี", "ทหาร"},
	"8":  {"แอคชัน", "แฟนตาซี", "เกม"},
	"9":  {"กีฬา", "โชเน็น", "ฟุตบอล"},
	"10": {"ดราม่า", "ลึกลับ", "ไอดอล"},
	"11": {"แฟนตาซี", "ผจญภัย", "ดราม่า"},
	"12": {"คอมเมดี้", "รอม com", "เหนือธรรมชาติ"},
}

var mockComments = map[string][]Comment{
	"1": {
		{Username: "LuffyFan", Score: 10, Comment: "เรื่องนี้คือชีวิต! ต้องอ่านให้จบ"},
		{Username: "PirateKing", Score: 9, Comment: "การผจญภัยและมิตรภาพสุดยอดมาก"},
	},
	"2": {{Username: "Tanjiro", Score: 9, Comment: "ภาพวาดสวยมาก เรื่องเศร้าแต่สวย"}},
	"3": {{Username: "Gojo", Score: 10, Comment: "Gojo sensei เจ๋งสุดๆ"}},
}

func MockNewManga() []MangaCard      { return mockMangaList[6:13] }
func MockTrendingManga() []MangaCard { return mockMangaList[0:6] }
func MockTop10Monthly() []MangaCard  { return mockMangaList[0:10] }
func MockTop10Yearly() []MangaCard   { return mockMangaList[0:10] }
func MockTop10AllTime() []MangaCard  { return mockMangaList[0:10] }

func MockMangaByID(id string) *MangaDetail {
	for _, m := range mockMangaList {
		if m.ID == id {
			synopsis := mockSynopsis[id]
			if synopsis == "" {
				synopsis = "เรื่องย่อจะเพิ่มในภายหลัง"
			}
			cats := mockCategories[id]
			if cats == nil {
				cats = []string{"มังงะ"}
			}
			return &MangaDetail{
				ID: m.ID, Title: m.Title, CoverURL: m.CoverURL, Score: m.Score,
				ReviewCount: m.ReviewCount, Synopsis: synopsis, Categories: cats,
			}
		}
	}
	return nil
}

func MockCommentsByMangaID(id string) []Comment { return mockComments[id] }

// ---------- Categories ----------

type Category struct {
	Name  string
	Icon  string
	Color string
	Count int
}

func MockCategories() []Category {
	return []Category{
		{Name: "แอคชัน", Icon: "⚔️", Color: "#ef4444", Count: 342},
		{Name: "โรแมนซ์", Icon: "💕", Color: "#ec4899", Count: 287},
		{Name: "แฟนตาซี", Icon: "🧙", Color: "#8b5cf6", Count: 256},
		{Name: "คอมเมดี้", Icon: "😂", Color: "#f59e0b", Count: 198},
		{Name: "สยองขวัญ", Icon: "👻", Color: "#6366f1", Count: 124},
		{Name: "กีฬา", Icon: "⚽", Color: "#22c55e", Count: 89},
		{Name: "ดราม่า", Icon: "🎭", Color: "#06b6d4", Count: 213},
		{Name: "โชเน็น", Icon: "🔥", Color: "#f97316", Count: 310},
		{Name: "ชีวิตประจำวัน", Icon: "☕", Color: "#a78bfa", Count: 145},
		{Name: "ผจญภัย", Icon: "🗺️", Color: "#14b8a6", Count: 178},
		{Name: "ลึกลับ", Icon: "🔎", Color: "#64748b", Count: 96},
		{Name: "Sci-Fi", Icon: "🚀", Color: "#3b82f6", Count: 73},
		{Name: "ดาร์กแฟนตาซี", Icon: "🌑", Color: "#71717a", Count: 62},
		{Name: "ไอดอล", Icon: "🎤", Color: "#e879f9", Count: 41},
		{Name: "ทหาร", Icon: "🎖️", Color: "#78716c", Count: 37},
	}
}

// ---------- Articles ----------

type Article struct {
	Slug     string
	Title    string
	Excerpt  string
	CoverURL string
	Tag      string
	Author   string
	Date     string
}

func MockArticles() []Article {
	return []Article{
		{Slug: "one-piece-gear-6", Title: "One Piece: ทฤษฎี Gear 6 จะมาในอาร์คสุดท้าย?", Excerpt: "แฟนๆ คาดเดากันว่า Gear 6 จะเป็นพลังสุดท้ายของลูฟี่ ก่อนจะเป็นราชาโจรสลัด", CoverURL: "https://picsum.photos/seed/art1/600/340", Tag: "ทฤษฎี", Author: "MangaGuru", Date: "10 ก.พ. 2026"},
		{Slug: "top-manga-2026", Title: "10 มังงะที่ต้องอ่านในปี 2026", Excerpt: "รวมมังงะใหม่และต่อเนื่องที่ไม่ควรพลาดในปีนี้ ครบทุกแนว ทุกรสชาติ", CoverURL: "https://picsum.photos/seed/art2/600/340", Tag: "แนะนำ", Author: "OtakuWriter", Date: "8 ก.พ. 2026"},
		{Slug: "jjk-ending-review", Title: "รีวิว Jujutsu Kaisen ตอนจบ: สมบูรณ์แบบหรือผิดหวัง?", Excerpt: "วิเคราะห์ตอนจบของ JJK ที่แฟนๆ มีความเห็นแตกต่างกันอย่างมาก", CoverURL: "https://picsum.photos/seed/art3/600/340", Tag: "รีวิว", Author: "MangaCritic", Date: "5 ก.พ. 2026"},
		{Slug: "manga-vs-anime", Title: "มังงะ vs อนิเมะ: อ่านหรือดูก่อนดี?", Excerpt: "เปรียบเทียบข้อดีข้อเสียของการอ่านมังงะต้นฉบับกับการดูอนิเมะดัดแปลง", CoverURL: "https://picsum.photos/seed/art4/600/340", Tag: "บทความ", Author: "AnimeNerd", Date: "2 ก.พ. 2026"},
		{Slug: "frieren-why-popular", Title: "ทำไม Frieren ถึงเป็นมังงะแห่งปี?", Excerpt: "วิเคราะห์ความสำเร็จของ Frieren ที่ทำลายกฎมังงะแบบเดิมๆ", CoverURL: "https://picsum.photos/seed/art5/600/340", Tag: "วิเคราะห์", Author: "MangaGuru", Date: "28 ม.ค. 2026"},
		{Slug: "beginner-manga-guide", Title: "เริ่มอ่านมังงะยังไง? คู่มือสำหรับมือใหม่", Excerpt: "ไกด์ครบจบในบทความเดียว ตั้งแต่เลือกแนว ไปจนถึงแหล่งอ่าน", CoverURL: "https://picsum.photos/seed/art6/600/340", Tag: "คู่มือ", Author: "OtakuWriter", Date: "25 ม.ค. 2026"},
	}
}

// ---------- Threads ----------

type Thread struct {
	ID        string
	Title     string
	Excerpt   string
	Author    string
	AvatarURL string
	Date      string
	Replies   int
	Tag       string
}

func MockThreads() []Thread {
	return []Thread{
		{ID: "1", Title: "One Piece อาร์คไหนดีที่สุด?", Excerpt: "ส่วนตัวชอบ Marineford มากสุด แต่เพื่อนว่า Wano ดีกว่า", Author: "LuffyFan", AvatarURL: "https://api.dicebear.com/9.x/thumbs/svg?seed=LuffyFan", Date: "12 ก.พ. 2026", Replies: 48, Tag: "พูดคุย"},
		{ID: "2", Title: "แนะนำมังงะแนว Isekai ที่ไม่ซ้ำใครหน่อย", Excerpt: "อ่าน isekai มาหลายเรื่อง หาเรื่องใหม่ๆ ที่ไม่เหมือนเรื่องอื่น", Author: "IsekaiHunter", AvatarURL: "https://api.dicebear.com/9.x/thumbs/svg?seed=IsekaiHunter", Date: "11 ก.พ. 2026", Replies: 32, Tag: "แนะนำ"},
		{ID: "3", Title: "Chainsaw Man Part 2 เป็นไงบ้าง?", Excerpt: "พาร์ท 2 สนุกเหมือนพาร์ทแรกไหม หรือลดลง?", Author: "DevilHunter", AvatarURL: "https://api.dicebear.com/9.x/thumbs/svg?seed=DevilHunter", Date: "10 ก.พ. 2026", Replies: 27, Tag: "รีวิว"},
		{ID: "4", Title: "มังงะเรื่องไหนทำให้ร้องไห้มากที่สุด?", Excerpt: "ส่วนตัวอ่าน Your Lie in April จบแล้วน้ำตาไหลเลย", Author: "CryBaby", AvatarURL: "https://api.dicebear.com/9.x/thumbs/svg?seed=CryBaby", Date: "9 ก.พ. 2026", Replies: 65, Tag: "พูดคุย"},
		{ID: "5", Title: "Blue Lock กับฟุตบอลจริง ต่างกันยังไง?", Excerpt: "คนเล่นบอลจริงมาแชร์ความเห็น เรื่องนี้สมจริงแค่ไหน?", Author: "FootballOtaku", AvatarURL: "https://api.dicebear.com/9.x/thumbs/svg?seed=FootballOtaku", Date: "8 ก.พ. 2026", Replies: 19, Tag: "วิเคราะห์"},
		{ID: "6", Title: "Solo Leveling จบดีไหม? (ระวังสปอยล์)", Excerpt: "อ่านจบแล้ว อยากคุยกับคนที่อ่านจบเหมือนกัน", Author: "ShadowMonarch", AvatarURL: "https://api.dicebear.com/9.x/thumbs/svg?seed=ShadowMonarch", Date: "7 ก.พ. 2026", Replies: 41, Tag: "สปอยล์"},
	}
}

func MockArticleBySlug(slug string) *Article {
	for _, a := range MockArticles() {
		if a.Slug == slug {
			return &a
		}
	}
	return nil
}

func MockThreadByID(id string) *Thread {
	for _, t := range MockThreads() {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

func MockSearchManga(q string) []MangaCard {
	if q == "" {
		return nil
	}
	q = strings.ToLower(strings.TrimSpace(q))
	var out []MangaCard
	for _, m := range mockMangaList {
		if strings.Contains(strings.ToLower(m.Title), q) {
			out = append(out, m)
		}
	}
	if len(out) == 0 {
		return mockMangaList
	}
	return out
}
