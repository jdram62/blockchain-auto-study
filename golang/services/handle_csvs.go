package services

import (
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"os"
)

func GetEndpoints() [][]string {
	file, err := os.Open("endpoints.csv")
	if err != nil {
		log.Fatalln("getWatchList - Open:", err)
	}
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("getWatchList - ReadAll:", err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln("getWatchList - Close:", err)
	}
	return csvData
}

func GetWatchList() [][]string {
	file, err := os.Open("watchlist.csv")
	if err != nil {
		log.Fatalln("getWatchList - Open:", err)
	}
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("getWatchList - ReadAll:", err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln("getWatchList - Close:", err)
	}
	return csvData
}

func AddContractWatchlist() (bool, error) {
	file, err := os.OpenFile("watchlist.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	var pass string
	inputs := make([]string, 3)
	fmt.Println("\n\nAdd New LP Contract to watchlist:")
	fmt.Println("Pair Address")
	_, err = fmt.Scanln(&inputs[0])
	if err != nil {
		return false, err
	}
	fmt.Println("Reserve0 Address")
	_, err = fmt.Scanln(&inputs[1])
	if err != nil {
		return false, err
	}
	fmt.Println("Reserve1 Address")
	_, err = fmt.Scanln(&inputs[2])
	if err != nil {
		return false, err
	}
	fmt.Println("\n\tVerify Inputs\nPair Address: ", inputs[0], "\nReserve0: ", inputs[1], "\nReserve1: ", inputs[2])
	fmt.Println("1 - Valid inputs\nELSE - Quit")
	_, err = fmt.Scanln(&pass)
	if err != nil {
		return false, err
	}
	switch pass {
	case "1":
		err = csvWriter.Write(inputs)
		if err != nil {
			return false, err
		}
		csvWriter.Flush()
		fmt.Println("successful contract add")
		return true, nil
	default:
		return false, nil
	}
}

func RemoveContractWatchlist() (bool, error) {
	var userInput string
	i := 0
	file, err := os.Open("watchlist.csv")
	if err != nil {
		return false, err
	}
	// read contents in prior to attempt to delete contract, close file instantly after read
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	err = file.Close()
	if err != nil {
		return false, err
	}
	rowsToWrite := make([][]string, len(csvData)-1)
	fmt.Println("\n\nRemove lp Contract from watchlist:\nELSE - Quit")
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		return false, err
	}
	// return if input is not address
	if !common.IsHexAddress(userInput) {
		return false, err
	}
	found := false
	for _, data := range csvData {
		if data[0] == userInput {
			found = true
		} else {
			// error if csv rows <= 1
			rowsToWrite[i] = data
			i++
		}
	}
	if found {
		// address is confirmed in csv, empty contents
		if err = os.Truncate("watchlist.csv", 0); err != nil {
			return false, err
		}
		// write contents to csv, deleted entry is not included
		file, err = os.OpenFile("watchlist.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		defer file.Close()
		csvWriter := csv.NewWriter(file)
		csvWriter.Flush()
		err = csvWriter.WriteAll(rowsToWrite)
		fmt.Println("successful contract removal")
		return true, err
	}
	return false, nil
}
