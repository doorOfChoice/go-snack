package main


const (
	LEFT = iota
	UP
	RIGHT
	DOWN
)

type SnakeBody struct {
	coord
	shape string
}

type Snake struct {
	bodys Vector
	direction int
}

func NewSnake(x, y int) *Snake{
	return &Snake{
		bodys : NewVector(),
		direction : UP,
	}
}

func (this *Snake) Grow(x, y int, shape string) {
	this.bodys.Add(&SnakeBody {
		coord : coord {x, y},
		shape : shape,
	})
}


func (this *Snake) ChangeDirection(direction int) {
	this.direction = direction
}

func (this *Snake) Move() (bool){
	size := len(this.bodys)
	for i := size - 1; i > 0; i-- {
		next := this.bodys[i - 1].(*SnakeBody)
		current := this.bodys[i].(*SnakeBody)
		current.coord.x = next.coord.x
		current.coord.y = next.coord.y
	}
	
	head := this.GetHead()
	
	switch this.direction {
	case UP: head.coord.y -= 1
	case DOWN : head.coord.y += 1
	case LEFT : head.coord.x -= 1
	case RIGHT : head.coord.x += 1
	}

	return !this.isBited()
}


func (this *Snake) isBited() bool{
	head := this.bodys[0].(*SnakeBody)

	for _, v := range this.bodys[1:] {
		body := v.(*SnakeBody)
		if head.coord.Equal(body.coord) {
			return true
		}
	}

	return false
}

func (this *Snake) GetHead() *SnakeBody {
	return this.bodys[0].(*SnakeBody)
}