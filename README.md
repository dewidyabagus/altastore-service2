# Alta Store Back-End Microservices Service 2

Alta Store Back-End (AS BE) adalah layanan web service (RESTful API) untuk kebutuhan pemrosesan data layanan Toko Online dimana fitur yang diberikan mulai dari registrasi, login pelanggan dan admin, master produk dan kategori produk, keranjang belanja dan metode pembayaran yang menggunakan payment gateway. Dengan sistem autentikasi JWT.

## Fitur Pada Service 2

- Penambahan produk kategori dan master produk barang
- Pembelian barang untuk kebutuhan stok tok

## Teknologi dan Arsitektur

`Teknologi` dan `Arsitektur` pengembangan sistem back-end mulai dari bahasa pemrograma, database server hingga infrastruktur yang digunakan.

- **_Development Tools_**

  - Golang (Program Language)
  - Echo Framework (Web Framework)
  - PostgreSQL (Database Server)
  - MongoDB (Database Logging)

- **_API Doc & Tester_**

  - Swagger
  - Postman

- **_Container_**

  - Docker

- **_Arsitektur_**
  - Heksagonal

## Memulai Web Service API

- Menjalankan **_Web Service API_** pada lingkungan pengembangan
  Sebelum menjalankan **_Web Service API_** pastikan anda sudah menduplikasi file `.env_example` dan menggantinya menjadi `.env`.
  File tersebut digunakan untuk menyimpan konfigurasi aplikasi, alamat database dan secret JWT.

  Untuk menjalankan **_Web Service API_** ketikan perintah ini di terminal anda:

  ```console
  go run .
  ```

  Anda dapat melihatnya di `http://localhost:8000` secara default sesuai dengan konfigurasi `.env`

- Menjalankan **_Web Service API_** dengan **_Docker_** <br>
  Kemudahan menjalankan **_Web Service API_** dengan **_Docker_**. Anda akan mendapatkan kemudahan dalam menyiapkan **_Web Service API_**
  dengan menggunakan docker tanpa harus melakukan instalasi `go language`, `postgresql` dan `mongodb`, cukup melakukan
  [install `docker`](https://docs.docker.com/engine/install/) dan jalankan perintah dibawah ini. Docker akan menyiapkan
  semua keperluan yang dibutuhkan.

  - Menggunaan Sistem Operasi Windows

    ```console
    windows-docker-compose.bat
    ```

  - Menggunaan Sistem Operasi Linux <br>
    Pastikan Anda sudah [menginstall docker compose](https://docs.docker.com/compose/install/)

    ```console
    /bin/bash linux-docker-compose.sh
    ```

  **_Docker_** akan melakukan build image web service sesuai dengan file konfigurasi `.env`,
  pastikan nama host `postgres` dan `mongodb` sesuai dengan nama container **_"secara default sudah sama"_**.<br>
  **_Docker-compose_** akan membuatkan container sesuai dengan setup lingkuan variabel dan pastikan kembali
  konfigurasi lingkungan variabel sudah sesuai **_"secara default sudah sama"_**

  Anda dapat melihatnya di `http://localhost:8000` secara default sesuai dengan konfigurasi `.env`

## Melakukan Request Web Service API

Kami menyiapkan yang terbaik untuk Anda. Untuk melihat dokumentasi **_Web Service API_** sudah kami siapkan di file `AltaStore.yaml`.
Kami juga menyiapkan collection **_Postman_** untuk pengujian dan permintaan ke **_Web Service API_** di file `BE Alta Store Service 2.json`
jika belum memiliki **_Posman_** Anda dapat [download dan install](https://www.postman.com/downloads/) Postman terlebih dahulu.

## Hasil Unit Testing

## Repository Service Lainnya

- [Alta Store Back-End Microservices Service 1](https://github.com/Yap0894/altastore-service1)
- [Alta Store Back-End Microservices Service 3](https://github.com/dewidyabagus/altastore-service3)

## Kontak

Kami sangat terbuka untuk berdiskusi, menerima kritik dan saran. Anda dapat menghubungi kami:

- Widya Ade Bagus - https://www.linkedin.com/in/widya-ade-bagus-3a660716b/
- Alexander Yap - https://www.linkedin.com/in/alexander-yap-a14015169/
