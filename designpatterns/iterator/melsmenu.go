package main

type melsMenu struct {
	menuList []menuItem
}

func (m *melsMenu) addItem(name string, amount int, description string) {
	m.menuList = append(m.menuList, menuItem{name, description, amount})
}
