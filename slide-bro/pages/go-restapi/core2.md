
---

## 1. รูปแบบพื้นฐาน

```go
var arr [3]int  // array ความยาว 3 เก็บ int ค่า default = 0
```

## 2. กำหนดค่าเริ่มต้น

```go
var nums = [3]int{1, 2, 3}
```

หรือให้คอมไพเลอร์นับเอง:

```go
nums := [...]int{1, 2, 3, 4, 5}  // ความยาว = 5
```

---

## 3. เข้าถึงสมาชิก

```go
nums := [3]int{10, 20, 30}
fmt.Println(nums[0])  // 10
nums[1] = 50
fmt.Println(nums)     // [10 50 30]
```

## 4. Array หลายมิติ

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
fmt.Println(matrix)  // [[1 2 3] [4 5 6]]
```

---

## ✅ สรุป

* `[n]T` = array ขนาดคงที่ (length = n, type = T)
* `[...]T{...}` = ให้ Go infer ความยาวจากจำนวน element
* Array มี **ขนาดตายตัว** → ถ้าอยาก dynamic/resizable ใช้ `slice` (`[]T`) แทน

---
