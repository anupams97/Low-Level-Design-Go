package main

type iterator interface {
	next() *menuItem
	hasNext() bool
}
