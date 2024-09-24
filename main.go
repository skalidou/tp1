// ft_coin.go
package main

import (
	"sort"
)

// Fonction pour trier le tableau d'entiers, ordre décroissant.
func bubbleSortDesc(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				// Échange des éléments pour trier en décroissant
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Fonction pour calculer le nombre minimum de pièces pour atteindre `amount`
func Ft_coin(coins []int, amount int) int {
	result := 0
	compteur := 0

	// Trier les pièces en ordre décroissant
	coins = bubbleSortDesc(coins)

	// Parcourir les pièces
	for i := 0; i < len(coins); i++ {
		for result+coins[i] <= amount {
			result += coins[i]
			compteur++
		}
	}

	// Si le montant exact est atteint, retourner le compteur, sinon -1
	if result == amount {
		return compteur
	} else {
		return -1
	}
}

// Fonction pour trouver le nombre manquant
func Ft_missing(nums []int) int {
	// Trier le tableau
	nums = bubbleSortDesc(nums)

	// Calculer la somme attendue
	expectedSum := (len(nums) * (len(nums) + 1)) / 2 // selon la règles qui dit que la somme de 0 à n = n(n+1) / 2

	// Calculer la somme actuelle
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}

	// Le nombre manquant est la différence entre les deux sommes
	return expectedSum - actualSum
}

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Trier les intervalles par leur point de fin
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	countToRemove := 0
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastEnd {
			countToRemove++
		} else {
			lastEnd = intervals[i][1]
		}
	}

	return countToRemove
}

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0] // Prix d'achat minimum
	maxProfit := 0        // Bénéfice maximal

	for _, price := range prices {
		// Mettre à jour le prix d'achat minimum
		if price < minPrice {
			minPrice = price
		}

		// Calculer le bénéfice potentiel
		potentialProfit := price - minPrice
		if potentialProfit > maxProfit {
			maxProfit = potentialProfit
		}
	}

	return maxProfit
}

func main() {
}
