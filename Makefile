#วิธีการสร้างและเรียกใช้งานโปรแกรม
#1.สร้างฐานข้อมูลชื่อ image_db นำเข้า ไฟล์ golang_upload_file/database/image_db.sql เพื่อเก็บข้อมูลรูปภาพ
#2.แก้ไข Database config ในไฟล์ .env ใส่ค่า config ของ Database ให้ถูกต้อง
#3.แก้ไข Token ในไฟล์ .env เป็นค่าอะไรก็ได้ TOKEN นี้จะไปเก็บไว้ใน hidden input
#4.เมื่อโปรแกรมทำงานจะเปิดเว็บบราวเซอร์ขึ้นมาให้ทดสอบ
#5.เลือกไฟล์ที่ต้องการทดสอบแล้วคลิกปุ่ม upload รอผลลัพธ์ ผลลัพธ์จะแสดงตรง status และ message

#การใช้งาน makefile
#1.make tool เรียกใช้งาน Go Modules และ Go package
#2.make run เพื่อ run server http://127.0.0.1:8080/
#3.make compile เพื่อ build ไฟล์สำหรับ linux
#4.make test เพื่อใช้งาน Unit test

tool:
	go mod init
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/joho/godotenv

run:
	go run main.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go

test:
	go test -v