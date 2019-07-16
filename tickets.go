package tickets

import "fmt"

type Personas interface {
	cost(price int) int
	quantity() int
	str() string
}

type Invitado struct {
	name     string
	cantidad int
}

type Jubilado struct {
	name     string
	cantidad int
}

type Normal struct {
	name     string
	cantidad int
}

func (p Invitado) cost(price int) int { return 0 }
func (p Invitado) quantity() int      { return p.cantidad }
func (p Invitado) str() string        { return p.name }

func (p Jubilado) cost(price int) int { return price / 2 }
func (p Jubilado) quantity() int      { return p.cantidad }
func (p Jubilado) str() string        { return p.name }

func (p Normal) cost(price int) int { return price }
func (p Normal) quantity() int      { return p.cantidad }
func (p Normal) str() string        { return p.name }

func GetTotal(price, normales, jubilados, invitados int) int {
	i := Invitado{name: "Invitado", cantidad: invitados}
	j := Jubilado{name: "Jubilado", cantidad: jubilados}
	n := Normal{name: "Normal", cantidad: normales}

	ps := []Personas{}
	ps = append(ps, i)
	ps = append(ps, j)
	ps = append(ps, n)

	total := 0

	for _, v := range ps {
		total += v.quantity() * v.cost(price)
		fmt.Printf("Cantidad de tipo <%s> %d, total: %d.\n", v.str(), v.quantity(), v.quantity()*v.cost(price))
	}

	return total
}
