# Căn bản
1. Khởi tạo biến
   - var width, height = 100, 50 //"int" được bỏ qua
   - block := b: Khởi tạo nhanh không cần dùng var
   - block := &Block{[]byte{}, []byte(data), prevHash}
     - block := &...: Khởi tạo con trỏ
     - []byte: Khởi tạo mảng byte
     - []byte{}: Khởi tạo mảng byte rỗng
     - []byte(data): Khởi tạo mảng byte với giá trị là data
   - blocks []*Block: Khởi tạo con trỏ mảng (dạng danh sách liên kết)


2. Run go system
   - go run main.go: Run single file <main>
   - go mod init github.com/haiphamthanh/golang-blockchain