# Go-Api
Repository untuk belajar membuat API dengan golang dan library Gin dan Gorm
#

Cara Setup dan Menjalankan
1. Clone/Download Repository
```
$ git clone https://github.com/habibmahaasin/Go-Api.git
```
2. Import SQL file yang ada pada direktori :
```
./sql/mysql.sql 
```
 atau 
```
./sql/postgres.sql 
```
sesuaikan penggunaan dengan database yang ingin digunakan

3. Sesuaikan penggunaan gorm mysql/postgresql pada halaman `app/database/database.go`
4. Sesuaikan isi file `.env` dengan settingan local yang ada
5. Sesuaikan penggunaan query pada halaman `modules/user/repository/repository.go` sesuai database yang digunakan

6. Run and Setup Golang
```
go run main.go
```
7. Endpoint siap digunakan.
8. Untuk dapat mengakses endpoint perlu set header `access-token` : `jwt token`
9. `access-token` dapat didapatkan setelah melakukan login dengan akun yang tersedia/mendaftarkan akun baru.

<br></br>
## List Endpoint yang dapat diakses saat ini :
- Login `/api/v1/user/login` 
  - Method `POST`
- Create `/api/v1/user/register` 
  - Method `POST` 
- List `/api/v1/user` 
  - Method `GET`
  - Set Header `access-token`
- Detail `/api/v1/user/:id` 
  - Method `GET` 
  - Set Header `access-token`
- Edit `/api/v1/user/:id` 
  - Method `PUT` 
  - Set Header `access-token`
- Delete `/api/v1/user/:id` 
  - Method `Delete` 
  - Set Header `access-token`
