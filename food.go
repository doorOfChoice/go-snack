package main

type Food struct {
	coord
	Score int
	Shape string
}

func NewFood(x, y, score int, shape string) *Food{
	return &Food {
		coord : coord {
			x : x,
			y : y,
		},
		Score : score,
		Shape : shape,
	}
}

