package main

import(
	"tool"
	"math/rand"
	"math"
	"fmt"
	"time"
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

const(
	RED = termbox.ColorRed
	BLACK = termbox.ColorBlack
	WHITE = termbox.ColorWhite
	GREEN = termbox.ColorGreen
	YELLOW = termbox.ColorYellow

	S_HEAD = "O"
	S_BODY = "+"
	G_FOOD = "@"
	G_WALL = "#"
)



type Game struct {
	width, height int
	snake *Snake
	food *Food
	keyChan chan Command
	score int
}

func NewGame(width, height int) *Game{
	snake := NewSnake(10, 10)
	snake.Grow(10, 10, S_HEAD)
	game := &Game{width, height, snake, nil, make(chan Command), 0}
	game.food, _ = game.randomFood()

	return game
}

func (this *Game) StartGame() {
	if err := termbox.Init(); err != nil {
		fmt.Println(err.Error())
		return
	}
	
	var endMessage string

	defer func() {
		termbox.Close()
		fmt.Println(endMessage)
	}()

	go ListenKeyEvent(this.keyChan)

	loop:
	for {
		select{
		case v := <- this.keyChan :
			switch(v.command) {
			case MOVE :
				way := parseDirection(v.key)
				
				if len(this.snake.bodys) == 1 {
					this.snake.ChangeDirection(way)
				}else {
					if math.Abs(float64(way - this.snake.direction)) != 2 {
						this.snake.ChangeDirection(way)
					}
				}
			case END:
				endMessage = "Game ended by your self"
				break loop
			}
		case <-time.After(8e7):
			if !this.snake.Move() || this.collision() {
				endMessage = fmt.Sprintf("Oh, you bite your self! your score is %d", this.score)
				break loop
			}

			this.eatFood()
			this.render()
		}
	}
}

func parseDirection(key termbox.Key) int{
	switch key {
		case termbox.KeyArrowUp: return UP
		case termbox.KeyArrowDown: return DOWN
		case termbox.KeyArrowLeft : return LEFT
		case termbox.KeyArrowRight: return RIGHT
		default: return -1
	}
}

func (this *Game) getEmpty() (tool.Vector, bool) {
	empty := tool.NewVector()

	for i := 1; i < this.height - 1; i++ {
		for j := 1; j < this.width - 1; j++ {
			is := true
			for _, v := range this.snake.bodys {
				if v.(*SnakeBody).EqualValue(j, i) {
					is = false
					break;
				}
			}

			if is {
				empty.Add(coord{j, i})
			}
		}
	}

	return empty, len(empty) == 0
}

func (this *Game) randomFood() (*Food, error){
	blocks, isEmpty:= this.getEmpty()

	if isEmpty {
		return nil, fmt.Errorf("空间已经满了")
	}

	rand.NewSource(int64(time.Now().Nanosecond()))
	index := rand.Intn(len(blocks))
	randBlock := blocks[index].(coord)


	return NewFood(randBlock.x, randBlock.y, rand.Intn(20), G_FOOD), nil
}

func (this *Game) eatFood() {
	head := this.snake.GetHead()
	
	if head.coord.Equal(this.food.coord) {
		this.snake.Grow(-1, -1, S_BODY)
		this.score += this.food.Score
		this.food, _= this.randomFood()	
	}
}

func (this *Game) collision() bool{
	head := this.snake.GetHead()

	return head.x <= 0          ||
		   head.x >= this.width ||
		   head.y <= 0          ||
		   head.y >= this.height
}

//全体渲染
func (this *Game) render() error{
	termbox.Clear(BLACK, BLACK)

	this.drawMap()
	this.drawSnake()
	this.drawFood()
	tprint(0, this.height + 2, fmt.Sprintf("Score:%d", this.score), WHITE, RED)
	return termbox.Flush()
}

//渲染蛇
func (this *Game) drawSnake() {
	
	for _, v := range this.snake.bodys {
		body := v.(*SnakeBody)
		
		tprint(body.x, body.y, body.shape, YELLOW, BLACK)
	}

}

//渲染地图
func (this *Game) drawMap() {
	for i := 0; i < this.height; i++ {
		for j := 0; j < this.width; j++ {
			if(i == 0 || i == this.height - 1 || j == 0 || j == this.width - 1) {
				tprint(j, i, G_WALL, WHITE, BLACK)
			}
		}
	}

}

//渲染食物
func (this *Game) drawFood() {
	tprint(this.food.x, this.food.y, this.food.Shape, WHITE, BLACK)
}


//渲染字符串
func tprint(x, y int, message string, fg, bg termbox.Attribute) {
	for _, v := range message {
		termbox.SetCell(x, y, v, fg, bg)
		x += runewidth.RuneWidth(v) 
	}
}