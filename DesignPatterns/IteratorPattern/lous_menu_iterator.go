package main

type lousMenuIterator struct {
	position     int
	lousMenuList []*menuItem
}

func (l *lousMenuIterator) next() *menuItem {
	defer func() {
		l.position++
	}()

	if l.hasNext() {
		return l.lousMenuList[l.position]
	}
	return nil

}
func (l *lousMenuIterator) hasNext() bool {
	if l.position < len(l.lousMenuList) {
		return true
	}
	return false
}
