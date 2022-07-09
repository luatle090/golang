package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NhanVien struct {
	id      int
	name    string
	manager bool
}

type Manager struct {
	nhanVien          NhanVien
	emloyeeManagement []NhanVien
}

type Company struct {
	dsNhanVien []NhanVien
}

func (c Company) ThemNhanVien() {
	for i := range c.dsNhanVien {
		c.dsNhanVien[i].id = i + 1
		c.dsNhanVien[i].name = "Nguyen Van A " + fmt.Sprintf("%v", i+1)
		c.dsNhanVien[i].manager = false
	}
}

func (c Company) XuatListNhanVien() {
	fmt.Println("----- Danh sach nhan vien ------")
	for _, nv := range c.dsNhanVien {
		fmt.Println(nv.id, nv.name, nv.manager)
	}
}

func (c Company) NominateManager(numberOfEmloyees int) int {
	fmt.Println("----- De cu manager -------")
	manager := rand.Intn(numberOfEmloyees)
	c.dsNhanVien[manager].manager = true
	return manager
}

func (c Company) EmloyeesUnderManager(managerIndex int) Manager {
	fmt.Println("----- Chon nhan vien cho Manager ------")
	numberOfEmloyees := len(c.dsNhanVien) / 2
	manager := Manager{c.dsNhanVien[managerIndex], make([]NhanVien, numberOfEmloyees)}

	for i := 0; i < numberOfEmloyees; i++ {
		var emloyeeUnderIndex int = rand.Intn(len(c.dsNhanVien))
		if c.dsNhanVien[emloyeeUnderIndex].id == c.dsNhanVien[managerIndex].id {
			i--
			continue
		}
		manager.emloyeeManagement[i] = c.dsNhanVien[emloyeeUnderIndex]
	}
	return manager
}

func (m Manager) XuatListNhanVienQuanLy() {
	fmt.Println("----- Danh sach nhan vien quan ly ------")
	for _, nv := range m.emloyeeManagement {
		fmt.Println(nv.id, nv.name, nv.manager)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numberOfEmloyees := rand.Intn(10) + 1
	company := Company{
		make([]NhanVien, numberOfEmloyees),
	}

	company.ThemNhanVien()
	company.XuatListNhanVien()
	managerIndex := company.NominateManager(numberOfEmloyees)
	company.XuatListNhanVien()
	manager := company.EmloyeesUnderManager(managerIndex)
	manager.XuatListNhanVienQuanLy()
}
