package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	Date        string  `json:"firstname"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
}

type TransactionList struct {
	transactions []Transaction
}

func NewTransactionList(transactions []Transaction) *TransactionList {
	return &TransactionList{
		transactions: transactions,
	}
}

func returnFloat(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}

func Filter(t Transaction, filterFunc func(t Transaction) bool) bool {
	if filterFunc(t) {
		return true
	}
	return false
}

func (tl *TransactionList) sum() float64 {
	sum := 0.0
	for _, transaction := range tl.transactions {
		sum += transaction.Amount
	}
	return sum
}

func foodFilter(t Transaction) bool {
	fmt.Println("within foodFilter")
	fmt.Println(len(strings.TrimSpace(t.Category)))
	fmt.Println(len("Food & Drink"))
	return t.Category == "Food & Drink"
}
func main() {
	fmt.Println("Hello world")

	f, _ := os.Open("sheets/Chase0588_Activity20190319.CSV")
	// reader := csv.NewReader(bufio.NewReader(f))
	var transactions []Transaction

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := Transaction{
			Date:        line[0],
			Description: line[2],
			Category:    line[3],
			Type:        line[4],
			Amount:      returnFloat(line[5]),
		}
		fmt.Println(data.Category)

		if Filter(data, foodFilter) {
			transactions = append(transactions, data)
		}
		out, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
	fmt.Println(len(transactions))
	tl := NewTransactionList(transactions)
	sum := tl.sum()
	fmt.Println(sum)
}
