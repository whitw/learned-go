package main
import "fmt"

func main(){
	c := 'a'
	switch {
	case '0' <= c && c <= '9':
		fmt.Printf("%c is a number", c)
	case 'a' <= c && c <= 'z':
		fmt.Printf("%c is a small letter", c)
	case 'A' <= c && c <= 'z':
		fmt.Printf("%c is a large letter", c)
	}
}
