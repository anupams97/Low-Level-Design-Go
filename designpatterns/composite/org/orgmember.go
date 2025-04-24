package org

import "fmt"

type orgMember interface {
	Name() string
	Role() string
	Print(indent string)
}

type Manager struct {
	ManagerName string
	ManagerRole string
	Reports     []orgMember
}

func (m *Manager) Name() string {
	return m.ManagerName
}

func (m *Manager) Role() string {
	return m.ManagerRole
}

func (m *Manager) Print(indent string) {
	fmt.Println(m.Name() + " " + m.Role())
	for _, child := range m.Reports {
		child.Print(indent + " ")
	}
}
func (m *Manager) Add(r orgMember) {
	m.Reports = append(m.Reports, r)
}

type Employee struct {
	EmpName string
	EmpRole string
}

func (e *Employee) Name() string {
	return e.EmpName
}

func (e *Employee) Role() string {
	return e.EmpRole
}

func (e *Employee) Print(indent string) {
	fmt.Println(indent + e.Name() + " " + e.Role())
}
