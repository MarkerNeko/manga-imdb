# Manga IMDB — Development Log & Checklist

> สำหรับ AI (Claude / Cursor) ที่จะมาอ่านและทำงานต่อ
> อัปเดตล่าสุด: 12 ก.พ. 2026

---

## สถาปัตยกรรมภาพรวม

- **ภาษา/เฟรมเวิร์ก**: Go 1.25, `net/http` (stdlib only, ไม่มี framework)
- **Template engine**: `html/template` — แบ่งเป็น layouts / components / pages
- **Live reload**: [air](https://github.com/air-verse/air) (`.air.toml`)
- **Build**: `go build -o ./tmp/main .` / `make dev` / `make run`
- **Port**: `:8080`

---

## โครงสร้างโฟลเดอร์

```
manga-imdb/
├── docs/
│   └── dev-log.md              ← ไฟล์นี้
├── internal/
│   └── site/
│       ├── render.go           ← Render() + template funcs (stars, starsInt, add, sub, mul, div)
│       ├── routes.go           ← PageRoutes() ลงทะเบียนทุก route
│       └── mockdata.go         ← Mock data มังงะ 12 เรื่อง + เรื่องย่อ + หมวดหมู่ + คอมเมนต์
├── web/
│   └── html/
│       ├── components/
│       │   ├── navbar.html     ← navbar fixed top + mobile menu
│       │   └── manga-card.html ← การ์ดมังงะ (รูป, badge คะแนน, ดาว, ชื่อ, จำนวนรีวิว)
│       ├── layouts/
│       │   └── default.html    ← layout หลัก + CSS ทั้งหมด (global styles)
│       └── pages/
│           ├── index.html      ← Home: hero, มังงะใหม่, มังงะมาแรง, Top 10 (แท็บ)
│           ├── info.html       ← รายละเอียดมังงะ: รูป, เรื่องย่อ, หมวดหมู่, คอมเมนต์
│           ├── search.html     ← ค้นหา: ช่องค้นหา + ผลลัพธ์ grid
│           ├── login.html      ← ฟอร์มเข้าสู่ระบบ
│           └── register.html   ← ฟอร์มสมัครสมาชิก
├── main.go                     ← entrypoint: สร้าง mux → PageRoutes → ListenAndServe
├── Makefile                    ← make dev / make run / make build
└── .air.toml                   ← config สำหรับ air (live reload)
```

---

## สิ่งที่ทำไปแล้ว (Completed)

### ระบบเทมเพลต
- [x] Layout หลัก (`default.html`) — ใช้ `block "title"` / `block "content"` / `block "meta"`
- [x] ฟังก์ชัน `Render()` — parse layout + components + page → merge data → execute
- [x] Template funcs: `add`, `sub`, `mul`, `div`, `stars`, `starsInt`
- [x] ใช้ `html/template` (ไม่ใช่ `text/template`) เพื่อรองรับ `template.HTML` สำหรับดาว

### หน้าเว็บ
- [x] **Home** (`/`) — 3 ส่วน: มังงะใหม่, มังงะมาแรง, Top 10 เรตติ้ง
- [x] **Top 10 แท็บ** — ประจำเดือน / ประจำปี / ตลอดกาล — JavaScript สลับ tab-panel
- [x] **Info** (`/info/{id}`) — รูปปก, ชื่อ, คะแนน+ดาว, หมวดหมู่ (tag), เรื่องย่อ, คอมเมนต์
- [x] **Search** (`/search?q=...`) — ช่องค้นหา + ผลลัพธ์เป็น manga-card grid
- [x] **Login** (`/login`) — ฟอร์ม username + password + ลิงก์ไปสมัคร
- [x] **Register** (`/register`) — ฟอร์ม username + email + password + confirm

### Routes
- [x] `GET /` → home (ส่ง mock data)
- [x] `GET /info/{id}` → info (ดึง mock manga + comments ตาม ID)
- [x] `GET /search` → search (ค้นจากชื่อ, ถ้าไม่เจอแสดงทั้งหมด)
- [x] `GET /login`, `POST /login` → login (POST redirect ไป `/`)
- [x] `GET /register`, `POST /register` → register (POST redirect ไป `/login`)

### Mock Data
- [x] มังงะ 12 เรื่อง: One Piece, Demon Slayer, Jujutsu Kaisen, Chainsaw Man, Spy x Family, My Hero Academia, Attack on Titan, Solo Leveling, Blue Lock, Oshi no Ko, Frieren, Dandadan
- [x] เรื่องย่อ (synopsis) ภาษาไทย ทั้ง 12 เรื่อง
- [x] หมวดหมู่ (categories) ทั้ง 12 เรื่อง
- [x] คอมเมนต์ตัวอย่าง สำหรับ ID 1, 2, 3
- [x] ฟังก์ชัน: MockNewManga, MockTrendingManga, MockTop10*, MockMangaByID, MockCommentsByMangaID, MockSearchManga

### UI / CSS
- [x] ธีมเข้ม: พื้น `#1a0f14`, การ์ดโทนม่วง, ตัวอักษร IBM Plex Sans Thai
- [x] Navbar fixed top + mobile responsive hamburger menu
- [x] `body { padding-top: 72px }` เพื่อไม่ให้ navbar ทับเนื้อหา
- [x] การ์ดมังงะ: รูป 3:4, badge คะแนนมุมขวาบน, hover ยกขึ้น + เงา glow
- [x] ดาวรีวิว 3 สถานะ: ★ เต็ม (สีทอง), ★ ครึ่ง (ซ้ายทอง ขวาเทา via CSS ::before), ★ ว่าง (สีเทา)
- [x] หน้า Auth (login/register): การ์ดกลางจอ, input มี focus glow ม่วง
- [x] หน้า Info: grid 220px + 1fr, แท็กหมวดหมู่, คอมเมนต์ list
- [x] หน้า Search: ฟอร์มค้นหา + grid ผลลัพธ์

---

## สิ่งที่ยังไม่ได้ทำ (TODO)

### ฐานข้อมูล
- [ ] เลือก DB (SQLite / PostgreSQL / ...)
- [ ] สร้างตาราง: `manga`, `users`, `reviews`
- [ ] Migration script
- [ ] เปลี่ยน mock data → query จาก DB จริง

### Authentication
- [ ] Hash password (bcrypt)
- [ ] Session / cookie-based auth
- [ ] Middleware ตรวจสอบ login
- [ ] แสดง username ใน navbar เมื่อ login แล้ว
- [ ] ซ่อนปุ่ม เข้าสู่ระบบ/สมัครสมาชิก เมื่อ login แล้ว

### รีวิว / คอมเมนต์
- [ ] ฟอร์มให้คะแนน (1–10) + เขียนคอมเมนต์ (ต้อง login)
- [ ] POST `/info/{id}/review` → บันทึกลง DB
- [ ] ป้องกันรีวิวซ้ำ (1 user : 1 review ต่อมังงะ)
- [ ] แสดงรีวิวจาก DB แทน mock

### Ranking / Top 10
- [ ] คำนวณ Top 10 จาก reviews จริง (GROUP BY + AVG)
- [ ] แยก: ประจำเดือน, ประจำปี, ตลอดกาล (WHERE created_at)
- [ ] มังงะมาแรง: จำนวนรีวิวใหม่ในช่วง 7 วัน

### ปรับปรุง UI
- [ ] Pagination สำหรับผลค้นหา
- [ ] Loading state / skeleton
- [ ] รูปปกมังงะจริง (ไม่ใช่ picsum.photos)
- [ ] Footer
- [ ] SEO meta tags

### อื่นๆ
- [ ] Static file serving (CSS/JS แยกไฟล์)
- [ ] Error handling (404, 500 pages)
- [ ] Input validation (XSS, CSRF)
- [ ] Deploy (Docker / fly.io / ...)

---

## หมายเหตุสำหรับ AI

- **Template**: ใช้ `html/template` (ไม่ใช่ `text/template`) — สำคัญเพราะฟังก์ชัน `stars` คืน `template.HTML`
- **Render flow**: `render.go` → parse layout glob → parse components glob → parse page file → merge data → execute `layouts/default`
- **CSS**: สไตล์ทั้งหมดอยู่ใน `layouts/default.html` + `components/navbar.html` (inline `<style>`)
- **Navbar padding-top**: มี 2 จุดที่ set `padding-top` ให้ `body` — ใน `default.html` (72px) และ `navbar.html` (70px) — cần ให้ sync กัน ถ้าเปลี่ยนความสูง navbar
- **Mock data**: อยู่ใน `mockdata.go` — เมื่อเชื่อม DB จริงให้แทนที่ฟังก์ชัน `Mock*` ทั้งหมด
- **ไม่มี external dependency** นอกจาก Go stdlib + air (dev tool) + Google Fonts (CDN)
