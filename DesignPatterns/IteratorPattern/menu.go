package main

type menu interface {
	createIterator() iterator
}
