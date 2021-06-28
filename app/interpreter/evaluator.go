package interpreter

import (
	"errors"
	"math"
)

type fn func([][]int, int, int) ([]int, error)

func EvaluateCSV(matrix [][]int, operation string, entireFile bool, fromLine int, toLine int) ([]int, error) {

	//Si calcul sur certains lignes
	var from, to int = 0, len(matrix)

	if !entireFile {
		if fromLine >= from && fromLine < to {
			from = fromLine
		}
		if toLine <= to && toLine > from {
			to = toLine
		}
	}

	functionHandler := map[string]fn{
		"somme":    somme,
		"moyenne":  moyenne,
		"mediane":  mediane,
		"maxValue": maxValue,
	}

	return functionHandler[operation](matrix, from, to)

}

//Operations functions
func somme(matrix [][]int, from int, to int) ([]int, error) {

	somme := make([]int, 0)

	for i := 0; i < len(matrix[0]); i++ {
		var sommeColumn int = 0
		for j := from; j < to; j++ {
			sommeColumn = sommeColumn + matrix[j][i]
		}
		somme = append(somme, sommeColumn)
	}
	return somme, nil
}

func moyenne(matrix [][]int, from int, to int) ([]int, error) {

	somme, _ := somme(matrix, from, to)

	if len(matrix)*len(matrix[0]) == 0 {
		return nil, errors.New("aucun champs numérique n'a été trouvé. Veuillez vérifier les paramètres entrées")
	}

	moyenne := make([]int, 0)

	for i := 0; i < len(somme); i++ {
		moyenne = append(moyenne, somme[i]/(len(matrix)))
	}

	return moyenne, nil
}

func mediane(matrix [][]int, from int, to int) ([]int, error) {
	return nil, nil
}

func maxValue(matrix [][]int, from int, to int) ([]int, error) {

	maxValue := make([]int, 0)

	for i := 0; i < len(matrix[0]); i++ {
		var maxColumn int = int(math.Inf(-1))
		for j := from; j < to; j++ {
			if maxColumn < matrix[j][i] {
				maxColumn = matrix[j][i]
			}
		}
		maxValue = append(maxValue, maxColumn)
	}

	return maxValue, nil
}
