# Go-Api
Repository untuk belajar membuat API dengan golang dan library Gin dan Gorm
#

Cara Setup dan Menjalankan
1. Clone/Download Repository
2. Import SQL file yang ada pada direktori :
```
./sql/gop-api.sql 
```
3. Run and Setup Golang
```
go run main.go
```
4. Endpoint siap digunakan.

# List Endpoint
Berikut ini list endpoint yang dapat digunakan
## User
- Login `/api/user/v1/login` - Method `POST`
- Create `/api/user/v1/create` - Method `POST`
- List `/api/user/v1/list` - Method `GET`
- Detail `/api/user/v1/detail/:id` - Method `GET`
- Edit `/api/user/v1/edit/:id` - Method `PUT`
- Delete `/api/user/v1/delete/:id` - Method `Delete`
