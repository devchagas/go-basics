package main

import "fmt"

type Piloto struct {
	nome        string
	equipe      string
	vitorias    int8
	nmrCorridas int8
}

func taxaVitoria(vitorias int8, nmrCorridas int8) float64 {
	taxa := 0.0
	taxa = (float64(vitorias) / float64(nmrCorridas)) * 100.00
	return taxa
}

func (piloto Piloto) String() string {
	return fmt.Sprintf("O piloto %s da equipe %s possui uma taxa de vitórias igual a: %.2f%%",
		piloto.nome, piloto.equipe, taxaVitoria(piloto.vitorias, piloto.nmrCorridas))
}

type Stringer interface {
	String() string
}

func show(stringer Stringer) {
	fmt.Println(stringer.String())
}

func main() {
	piloto1 := Piloto{
		nome:        "Yuki Tsunoda",
		equipe:      "AlphaTauri",
		vitorias:    7,
		nmrCorridas: 22,
	}

	piloto2 := Piloto{
		nome:        "Max Verstappen",
		equipe:      "Red Bull Racing",
		vitorias:    19,
		nmrCorridas: 22,
	}

	show(piloto1)
	show(piloto2)

}
