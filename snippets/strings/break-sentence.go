package main

import "fmt"

func main() {
	sentence := "Tiramisu powder marshmallow donut lemon drops carrot cake liquorice gummi bears. Macaroon powder carrot cake cotton candy caramels. Chocolate topping candy canes tootsie roll jelly. Croissant sesame snaps cheesecake bear claw pudding cupcake gummi bears pudding. Ice cream sweet roll toffee liquorice tootsie roll carrot cake bear claw muffin. Gingerbread chocolate gummi bears tiramisu soufflé powder. Pastry shortbread chocolate marshmallow powder fruitcake bonbon dragée cookie. Toffee bear claw sugar plum sweet roll gummi bears. Candy canes cake icing cake chocolate bar jelly. Tootsie roll cheesecake fruitcake sweet sesame snaps tiramisu macaroon. Shortbread jelly beans dessert ice cream halvah. Cheesecake chocolate pudding brownie shortbread powder bear claw."
	sentenceBreaks := []string{}
	limit := 42
	cursorPosition := limit

	for len(sentence) > 0 {
		if cursorPosition > len(sentence) {
			sentenceBreaks = append(sentenceBreaks, sentence)
			break
		}
		if sentence[cursorPosition-1] != 32 {
			cursorPosition--
		} else {
			partOfSentence := sentence[:cursorPosition-1]
			sentenceBreaks = append(sentenceBreaks, partOfSentence)
			sentence = sentence[cursorPosition:]
			cursorPosition = limit
		}
	}
	fmt.Printf("%q\n", sentenceBreaks)
}
