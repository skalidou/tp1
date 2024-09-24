package tp1

// Fonction pour trier le tableau d'entiers, ordre décroissant.
func bubbleSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				// Échange des éléments si l'ordre est incorrect (croissant)
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Ft_coin(coins []int, amount int) int {

	result := 0
	compteur := 0
	bubbleSortDesc(coins)

	for i := 0; i < len(coins); i++ {
		if result > amount {
			return -1
		}
		for j := 0; j < len(coins) && result != compteur; j++ {
			if coins[j] == amount {
				return 1
			} else if coins[j] > amount {
				continue
			} else {
				compteur += 1
				result += coins[j]
			}
		}
	}
	return compteur
}

func main() {
	Ft_coin([]int{1, 2, 5}, 11) // resultat : 3 car (11 == 5 + 5 + 1)
	Ft_coin([]int{2}, 3)        // resultat : -1
	Ft_coin([]int{1}, 0)        // resultat : 0

}
