package main

import (
	"fmt"
)

type PhanSo struct {
	tu, mau int
}

func (ps PhanSo) cong(ps2 PhanSo) PhanSo {
	var kq PhanSo = PhanSo{0, 1}
	kq.tu = ps.tu*ps2.mau + ps.mau*ps2.tu
	kq.mau = ps.mau * ps2.mau
	return kq
}

func (ps PhanSo) nhan(ps2 PhanSo) PhanSo {
	var kq PhanSo = PhanSo{0, 1}
	kq.tu = ps.tu * ps2.tu
	kq.mau = ps.mau * ps2.mau
	return kq
}

func main() {
	ps1 := PhanSo{tu: 9, mau: 4}
	ps2 := PhanSo{tu: 3, mau: 5}

	ps3 := ps1.cong(ps2)
	ps4 := ps1.nhan(ps2)
	fmt.Println(ps3.tu, ps3.mau)
	fmt.Println(ps4.tu, ps4.mau)
}
