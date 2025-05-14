# Design Patterns in Go

### 1. Định nghĩa bộ chuỗi công đoạn -> Builder.
### 2. Định nghĩa đầu ra sản phẩm -> Factory

### 3. Adapter sinh ra để đại diện cho thiết bị S không hợp chuẩn kết nối A, kết nối với kết nối A và convert tín hiệu sang chuẩn thiết bị S hỗ trợ. Adapter sẽ bọc thiết bị S bên trong khai báo (có vẻ giống như Proxy Design pattern.

### 4. Abtract Factory: Khi ta biết rõ sản phẩm, bộ sản phẩm ta muốn làm ra nhưng muốn tuy biến được Factory sản xuất. DP này sẽ trả về cho ta 1 factory cụ thể có thể thực hiện được sản phẩm/ bộ sản phẩm đó bằng cách truyền vào tên của nhà máy.