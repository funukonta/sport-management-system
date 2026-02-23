## Soal 2

### 1. Website ayo.co.id

**a. Generate Username Tanpa Jeda**
Saat di form pendaftaran, setiap ketik nama langsung trigger request ke `generate-username`. Sepertinya tidak pakai debounce. Seharusnya diberi jeda sekitar 500-800ms supaya server tidak terbebani request berulang.

**b. Verifikasi OTP Kurang Efisien**
Pada proses pendaftaran, setelah verifikasi via WA/SMS, user masih diminta verifikasi email lagi. Menurut saya verifikasi email sebaiknya dijadikan opsional atau baru diwajibkan saat user mau booking. Ini bisa meningkatkan conversion rate karena mengurangi friction di awal.

**c. Bug Email "Hangus"**
Ada bug yang cukup mengganggu:
1. User isi form pendaftaran
2. Masuk halaman OTP
3. Kembali ke form (back)
4. Klik next lagi, muncul error "Isian email sudah digunakan sebelumnya."

Padahal OTP belum diterima dan akun belum jadi. Emailnya langsung tidak bisa dipakai lagi. Solusinya mungkin pakai status `pending_verification` atau kasih expiry time untuk registrasi yang belum selesai.

---

### 2. Aplikasi AYO

**a. Notifikasi WhatsApp**
Bakal lebih baik kalau notifikasi booking (jadwal, lokasi, waktu bayar) juga dikirim ke WhatsApp, tidak hanya push notification. Lebih mudah dilacak dan praktis.

**b. Mini App Score Counter**
Kalau bisa ditambahkan fitur untuk mencatat skor pertandingan dan timeline match. Jadi aplikasinya tidak hanya untuk booking, tapi juga dipakai saat sedang bermain. Ini bisa meningkatkan user engagement.


## Soal 3

### 1. Kenapa Saya Cocok untuk Posisi Ini

Saya punya background sebagai Backend Engineer dan Technical Trainer. Jadi selain bisa membangun sistem yang scalable, saya juga bisa menjelaskan hal teknis ke tim dengan cara yang mudah dipahami. Dari sisi teknis, saya terbiasa bekerja dengan Golang, PostgreSQL, API design, dan optimasi database. Dari sisi trainer, saya terbiasa membuat dokumentasi yang rapi dan melakukan knowledge sharing dengan tim.

---

### 2. Tiga Keunggulan Utama

**1. Backend & Problem-Solving**
Saya sudah terbiasa menangani API development, optimasi query, dan debugging performa sistem. Nyaman membaca execution plan dan mengidentifikasi bottleneck.

**2. Curious & Continuous Learner**
Saya selalu berusaha update dengan teknologi terbaru, best practices, dan tools yang bisa meningkatkan produktivitas. Saya percaya learning is a lifelong journey.

**3. Communication**
Pengalaman sebagai trainer membuat saya bisa menjelaskan hal kompleks menjadi sederhana. Sangat berguna saat kolaborasi, code review, atau diskusi teknis.

---

### 3. Pencapaian Relevan

Beberapa hal yang pernah saya kerjakan:
- Membangun API backend untuk aplikasi dengan berbagai platform (web, mobile, desktop) yang melayani ribuan pengguna
- Membuat proses otomasi untuk ticketing system dengan Whatsapp API
- Membantu puluhan peserta bootcamp sehingga mampu membangun project production-ready

---

### 4. Motivasi Melamar

Saya ingin kembali lebih fokus ke engineering role dan terlibat langsung dalam pengembangan sistem yang berdampak. Saya ingin terus berkembang secara teknis, menghadapi problem yang lebih kompleks, dan berkontribusi bersama tim yang profesional.

---

### 5. Rencana 3 Bulan Pertama

**Bulan 1**  
Memahami business domain, product flow, dan arsitektur sistem. Membaca codebase dan mengambil task kecil-menengah untuk memahami workflow tim.

**Bulan 2**  
Mulai mengambil ownership pada fitur/improvement tertentu. Memberikan masukan teknis yang konstruktif.

**Bulan 3**  
Mulai terlibat dalam perencanaan jangka panjang, optimasi sistem, dan mentoring anggota tim yang lebih junior.
