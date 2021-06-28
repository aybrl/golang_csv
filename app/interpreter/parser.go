package interpreter

/**
	Cette classe et ses fonctions permettent de scanner le fichier
	et retourner une matrice contenant l'ensemble des valeurs y présentes
**/

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type fileHeader struct {
	nColumns int
	header   []string
}

var matrix [][]int

func ParseCSV(file string, withHeader bool, seperator string) ([][]int, []string, error) {
	//Open file
	fileContent, err := os.Open(file)

	//Open Error
	if err != nil {
		return nil, nil, errors.New((returnError("lecture du fichier impossible", err)))
	}

	fileScanner := bufio.NewScanner(fileContent)

	fileScanner.Split(bufio.ScanLines)

	//Set the Header
	var header fileHeader
	if withHeader {
		read := fileScanner.Scan()
		if read {
			header = *getHeader(fileScanner.Text(), seperator)
		}
	}

	//Infer Columns Types
	var colmunsType []reflect.Type
	read := fileScanner.Scan()

	if read {
		colmunsType, err = inferColumnsType(fileScanner.Text(), header, seperator)
		if err != nil {
			return nil, nil, err
		}
	}

	//Construct Values Matrix
	matrix = nil
	matrix, err := constructMatrix(fileScanner, &header, &colmunsType, seperator)

	if err != nil {
		return nil, nil, err
	}

	fileContent.Close()

	return matrix, header.header, nil
}

func getHeader(line string, seperator string) *fileHeader {

	headerSlice := parseLine(line, seperator)

	var h_slice []string

	//Delete " surrouding titles
	for i := 0; i < len(headerSlice); i++ {
		var s = headerSlice[i]
		s = strings.TrimPrefix(s, "\"")
		s = strings.TrimSuffix(s, "\"")
		h_slice = append(h_slice, s)
	}

	header := fileHeader{
		nColumns: len(headerSlice),
		header:   h_slice,
	}

	return &header
}

func inferColumnsType(line string, header fileHeader, seperator string) ([]reflect.Type, error) {

	row := parseLine(line, seperator)

	if len(row) != header.nColumns {
		return nil, errors.New(returnError("ligne 1 erronée : nombre des columns incohérent"))
	}

	var temp_s []int
	for i := 0; i < len(row); i++ {
		value, _ := strconv.Atoi(row[i])
		temp_s = append(temp_s, value)
	}
	matrix = append(matrix, temp_s)

	return getColumnTypes(row), nil
}

func constructMatrix(fileScanner *bufio.Scanner, header *fileHeader, columnTypes *[]reflect.Type, seperator string) ([][]int, error) {

	pointer := 2

	//Reading and Constructing
	for fileScanner.Scan() {
		row := parseLine(fileScanner.Text(), seperator)

		//Check that row length match the header length
		if len(row) != header.nColumns {
			return nil, errors.New(returnError("ligne " + strconv.Itoa(pointer) + " erronée : nombre des columns incohérent"))
		}

		//Check that row columns types match the header types
		rowTypes := getColumnTypes(row)
		comparison := compareColumnTypes(rowTypes, *columnTypes)

		if !comparison {
			return nil, errors.New(returnError("ligne " + strconv.Itoa(pointer) + " erronée : type des columns incohérent"))
		}

		var temp_s []int
		for i := 0; i < len(row); i++ {
			if rowTypes[i] == reflect.TypeOf(i) {
				value, _ := strconv.Atoi(row[i])
				temp_s = append(temp_s, value)
			}
		}
		matrix = append(matrix, temp_s)
		pointer++
	}

	//Return
	return matrix, nil
}

func parseLine(line string, seperator string) []string {
	return strings.Split(line, seperator)
}

func getColumnTypes(row []string) []reflect.Type {

	var types_slice []reflect.Type

	for i := 0; i < len(row); i++ {
		num, err := strconv.Atoi(row[i])
		if err == nil {
			types_slice = append(types_slice, reflect.TypeOf(num))
		} else {
			types_slice = append(types_slice, reflect.TypeOf(row[i]))
		}
	}

	return types_slice
}

func compareColumnTypes(typeOfRow []reflect.Type, types []reflect.Type) bool {
	for i := 0; i < len(typeOfRow); i++ {
		if typeOfRow[i] != types[i] {
			return false
		}
	}
	return true
}

func returnError(message string, err ...error) string {
	errorMessage := "Erreur de parsing du fichier : " + message + "!"

	if len(err) > 0 {
		errorMessage += " Erreur detaillée : " + err[0].Error()
	}

	return errorMessage
}
