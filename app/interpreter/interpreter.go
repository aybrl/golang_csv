package interpreter

func Interpreter(filepath string, withHeader bool, seperator string, entireFile bool, fromLine int, toLine int, operations []string) (map[string][]int, []string, error) {

	//Parsing
	valuesMatrix, header, err := ParseCSV(filepath, withHeader, seperator)

	if err != nil {
		return nil, nil, err
	}

	results := make(map[string][]int)

	//Evaluating
	for _, operation := range operations {
		results[operation], err = EvaluateCSV(valuesMatrix, operation, entireFile, fromLine, toLine)
		if err != nil {
			return nil, nil, err
		}
	}

	//Results
	return results, header, nil
}
