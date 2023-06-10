package main
import "fmt"

func main(){
	f1()
	fmt.Println("==================")
	f3()
}

func f1(){
	fmt.Println("f1.start")
	defer f2()
	fmt.Println("f1.end")
}

func f2(){
	fmt.Println("f2.start")
	defer fmt.Println("f2.deferred")
	fmt.Println("f2.end")
}

func f3(){
	for i := 0; i < 5;i++{
		defer fmt.Printf("%d", i)
		//defer uses stack
	}
}
