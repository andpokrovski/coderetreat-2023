package main

type Field [][]bool
type Preset [][2]int

type Game struct {
	field Field
}

func New(preset Preset, fieldSize int) *Game {
	var f Field
	f = CreateField(fieldSize)

	for _, elem := range preset {
		f[elem[0]][elem[1]] = true
	}

	return &Game{
		field: f,
	}
}

func CreateField(fieldSize int) Field {
	field := make(Field, fieldSize)
	for i := 0; i < fieldSize; i++ {
		field[i] = make([]bool, fieldSize)
	}

	return field
}

func (g *Game) State() Field {
	return g.field
}

func (g *Game) CalculateState() {
	g.field[0][0] = false
}
