package main

func main() {
	var c byte = 'd'
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		fmt.println("the entered char is vowel")
	} else {
		fmt.println("the entered char is a consonant")
	}
}
