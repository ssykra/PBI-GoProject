# Rakamin X Evermos Virtual Internship - Backend Developer Project

Project ini merupakan project berbasis internship dengan posisi **Backend Developer** dari **Rakamin x Evermos**.

Aplikasi ini dibangun menggunakan **Golang**, **GORM** dan **GIN FRAMEWORK** dengan prinsip **Clean Architecture** dan JWT untuk autentikasi.

---

## Fitur

1. **User Service**
- Registrasi dan Login menggunakan JWT
- User hanya dapat mengelola akun sendiri
2. **Toko Service**
- Toko otomatis dibuat saat user mendaftar akun
- User hanya dapat mengelola tokonya sendiri
3. **Alamat Service**
- User dapat membuat, mengupdate, melihat, dan menghapus alamat sendiri
- Alamat digunakan untuk pengiriman transaksi
- User tidak dapat mengelola alamat milik user lain
4. **Kategori Service**
- Hanya **Admin** yang dapat mengelola kategori (CRUD)
- Untuk membuat admin, ubah langsung status `is_admin` di database
5. **Produk Service**
- User dapat membuat, mengupdate, melihat, dan menghapus produk miliknya
- User tidak dapat mengelola produk milik user lain
- Mendukung upload file untuk foto produk
- Terdapat pagination dan Filtering
6. **Transaksi Service**
- User dapat membuat transaksi sendiri
- Detail transaksi otomatis tersimpan di tabel `log_produk`
- `log_produk` menyimpan snapshot produk saat transaksi dilakukan
- User tidak dapar mengelola transaksi milik user lain
7. **Keamanan & Validasi**
- JWT untuk autentikasi
- Validasi input, email dan nomor telepon unik
- Pagination dan filtering data
- Mengikuti clean architecture
