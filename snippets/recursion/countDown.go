package main

func countDown(number int) {
    // condition d'arrêt
	if number <= 0 {
		return
	}
    // résolution du problème
	print(number)
	print("\n")
    // appel récursif
	number--
    countDown(number)
}

func main() {
	countDown(9)
}
