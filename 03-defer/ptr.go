package main
import "fmt"

func inc(i *int){
	*i = *i + 1
}

func main(){
	i := 1
	inc(&i)
	fmt.Println(i)
}
