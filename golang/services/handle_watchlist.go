package services

import (
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"strconv"
)

func AddContractWatchlist() (bool, error) {
	file, err := os.OpenFile("watchlist.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	// loop for entering pair contract info
	var pass string
	inputs := make([]string, 2)
	exchanges := [3]string{"HERMES", "TETHYS", "NETSWAP"}
	for {
		fmt.Println("\n\nLiquidity Pool's Exchange")
		fmt.Println("1 - HERMES\n2 - TETHYS\n3 - NETSWAP")
		_, err = fmt.Scanln(&inputs[0])
		if err != nil {
			return false, err
		}
		switch inputs[0] {
		case "1", "2", "3":
			ind, _ := strconv.Atoi(inputs[0])
			inputs[0] = exchanges[ind-1]
		}
		fmt.Println("Liquidity Pool's Address")
		_, err = fmt.Scanln(&inputs[1])
		if err != nil {
			return false, err
		}
		fmt.Println("\n\tVerify Inputs\nExchange: ", inputs[0], "\nAddress: ", inputs[1])
		fmt.Println("1 - Valid inputs\n2 - Quit")
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
		case "2":
			return false, nil
		}
	}
}

func RemoveContractWatchlist() (bool, error) {
	var userInput string
	i := 0
	file, err := os.Open("watchlist.csv")
	if err != nil {
		return false, err
	}
	// read contents in prior to attempt delete contract, close file instantly after read
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	err = file.Close()
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	rowsToWrite := make([][]string, len(csvData)-1)
	for loop := true; loop; {
		fmt.Println("\n\nRemove lp from watchlist by contract address: ")
		fmt.Println("1 - Quit")
		_, err = fmt.Scanln(&userInput)
		if err != nil {
			return false, err
		}
		if userInput == "1" {
			return false, nil
		}
		if common.IsHexAddress(userInput) {
			for _, data := range csvData {
				if data[1] == userInput {
					loop = false
				} else {
					// error if csv rows <= 1
					rowsToWrite[i] = data
					i++
				}
			}
		}
	}
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
