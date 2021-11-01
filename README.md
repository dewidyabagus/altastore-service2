# Alta Store Back-End Microservices Service 2

Alta Store Back-End (AS BE) adalah layanan web service (RESTful API) untuk kebutuhan pemrosesan data layanan Toko Online dimana fitur yang diberikan mulai dari registrasi, login pelanggan dan admin, master produk dan kategori produk, keranjang belanja dan metode pembayaran yang menggunakan payment gateway.

## Teknologi dan Arsitektur
```Teknologi``` dan ```Arsitektur``` pengembangan sistem back-end mulai dari bahasa pemrograma, database server hingga infrastruktur yang digunakan.

- ***Development Tools***
  - Golang (Program Language)
  - Echo Framework (Web Framework)
  - PostgreSQL (Database Server)
  - MongoDB (Database Logging)

- ***API Doc & Tester***
  - Swagger
  - Postman

- ***Container***
  - Docker

- ***Arsitektur***
  - Heksagonal

## Langkah Awal Memulai Web Service API
- Menjalankan ***Web Service API*** pada lingkungan pengembangan
  Sebelum menjalankan ***Web Service API*** pastikan anda sudah menduplikasi file ```.env_example``` dan menggantinya menjadi ```.env```.
  File tersebut digunakan untuk menyimpan konfigurasi aplikasi, alamat database dan secret JWT.
  
  Untuk menjalankan ***Web Service API*** ketikan perintah ini di terminal anda:
  ```console
  go run .
  ```

- Menjalankan ***Web Service API*** dengan ```docker-compose```
  - Langkah pertama build terlebih dahulu image docker ***Web Service API***
  
## Hasil Unit Testing

## Kontak
