# Design Patterns in Go

### 1. Định nghĩa bộ chuỗi công đoạn -> Builder.
### 2. Định nghĩa đầu ra sản phẩm -> Factory

### 3. Adapter sinh ra để đại diện cho thiết bị S không hợp chuẩn kết nối A, kết nối với kết nối A và convert tín hiệu sang chuẩn thiết bị S hỗ trợ. Adapter sẽ bọc thiết bị S bên trong khai báo (có vẻ giống như Proxy Design pattern.

### 4. Abtract Factory: 
Khi ta biết rõ sản phẩm, bộ sản phẩm ta muốn làm ra nhưng muốn tuy biến được Factory sản xuất. DP này sẽ trả về cho ta 1 factory cụ thể có thể thực hiện được sản phẩm/ bộ sản phẩm đó bằng cách truyền vào tên của nhà máy.

### 5. Singleon: 
Để an toàn và đúng đắn, khi viết Singleton cần dùng Lock và Double- Check
```go
func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

```


> Giả sử có hai goroutine (A và B) cùng gọi getInstance() tại cùng một thời điểm.
>
> Cả A và B đều vượt qua dòng đầu tiên if singleInstance == nil, vì tại thời điểm đó singleInstance chưa được tạo.
>
>Cả A và B đều đi tiếp vào phần lock.Lock().
>
>Goroutine A chiếm được lock trước.
>
>Goroutine B bị chặn và phải chờ lock được thả.
>
>Goroutine A tạo singleInstance và gán nó (singleInstance = &Single{}), sau đó unlock.
>
>Lúc này, goroutine B tiếp tục thực hiện từ phần sau lock.Lock().
>
>Nếu không có if singleInstance == nil lần 2, thì goroutine B cũng sẽ tạo thêm một instance khác, phá vỡ nguyên tắc singleton.
> 
> 
