// controw flow

// if else does'nt have ()
if name != "" {
	fmt.Printf("Hello %s\n", name)
} else {
	fmt.Println("Hello friend")
}

// ใส่ statement เข้ามาไว้ใน if ก่อน condition แล้วใช้ ; คั่น
if n, err := strconv.Atoi("5s"); err != null {
	fmt.Println("NaN:", err)
} else {
	fmt.Println(n)
}
// from
n, err := strconv.Atoi("5s")
if err != null {
	fmt.Println("NaN:", err)
} else {
	fmt.Println(n)
}