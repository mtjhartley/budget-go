package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CategoryFilter string

const (
	Bills         CategoryFilter = "Bills & Utilities"
	Entertainment CategoryFilter = "Entertainment"
	Food          CategoryFilter = "Food & Drink"
	Gas           CategoryFilter = "Gas"
	Groceries     CategoryFilter = "Groceries"
	Home          CategoryFilter = "Home"
	Personal      CategoryFilter = "Personal"
	Shopping      CategoryFilter = "Shopping"
	Travel        CategoryFilter = "Travel"
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

func Filter(t Transaction, categoryFilter CategoryFilter, filterFunc func(t Transaction, categoryFilter CategoryFilter) bool) bool {
	if filterFunc(t, categoryFilter) {
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

func sumTransactions(transactions []Transaction) float64 {
	sum := 0.0
	for _, transaction := range transactions {
		sum += transaction.Amount
	}
	return sum
}

func transactionFilter(t Transaction, categoryFilter CategoryFilter) bool {
	return t.Category == string(categoryFilter)
}

func numCategories(transactions []Transaction, categoryFilter CategoryFilter) int {
	count := 0
	for _, transaction := range transactions {
		if transactionFilter(transaction, categoryFilter) {
			count++
		}
	}
	return count
}

func createCategorizedTransactionsList(transactions []Transaction, categoryFilter CategoryFilter) []Transaction {
	var categorizedTransactions []Transaction
	for _, transaction := range transactions {
		if transactionFilter(transaction, categoryFilter) {
			categorizedTransactions = append(categorizedTransactions, transaction)
		}
	}
	return categorizedTransactions
}

func main() {
	fmt.Println("Hello world")

	f, _ := os.Open("sheets/Chase0588_Activity20190319.CSV")

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		panic(err)
	}

	var allTransactions []Transaction

	for _, line := range lines {
		data := Transaction{
			Date:        line[0],
			Description: line[2],
			Category:    line[3],
			Type:        line[4],
			Amount:      returnFloat(line[5]),
		}
		if data.Amount > 0 {
			continue //Filter out paying off credit card
		}
		allTransactions = append(allTransactions, data)

	}

	// billTransactions := createCategorizedTransactionsList(allTransactions, Bills)
	// entertainmentTransactions := createCategorizedTransactionsList(allTransactions, Entertainment)
	// foodTransactions := createCategorizedTransactionsList(allTransactions, Food)
	// gasTransactions := createCategorizedTransactionsList(allTransactions, Gas)
	groceriesTransactions := createCategorizedTransactionsList(allTransactions, Groceries)
	// homeTransactions := createCategorizedTransactionsList(allTransactions, Home)
	// personalTransactions := createCategorizedTransactionsList(allTransactions, Personal)
	// shoppingTransactions := createCategorizedTransactionsList(allTransactions, Shopping)
	// travelTransactions := createCategorizedTransactionsList(allTransactions, Travel)

	fmt.Println(groceriesTransactions)
	groceriesMoney := sumTransactions(groceriesTransactions)
	groceriesTimes := numCategories(allTransactions, Groceries)
	fmt.Println(groceriesMoney)
	fmt.Println(groceriesTimes)
	fmt.Println(len(groceriesTransactions))
}
