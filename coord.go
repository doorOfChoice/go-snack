package main

type coord struct {
	x, y int
}

func (this coord) Equal(c coord) bool {
	return this.x == c.x && this.y == c.y
}

func (this coord) EqualValue(x, y int) bool {
	return this.x == x && this.y == y
}