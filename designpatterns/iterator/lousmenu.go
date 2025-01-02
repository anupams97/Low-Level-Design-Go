package main

type lousMenu struct {
	menuList []*menuItem
}

func (m *lousMenu) addItem(name string, amount int, description string) {
	m.menuList = append(m.menuList, &menuItem{name, description, amount})
}

func (m *lousMenu) createIterator() iterator {
	return &lousMenuIterator{
		position:     0,
		lousMenuList: m.menuList,
	}
}
