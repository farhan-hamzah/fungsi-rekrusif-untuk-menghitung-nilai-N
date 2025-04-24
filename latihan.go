package main
import "fmt"

func main(){
	var n, hasil float64
	fmt.Scan(&n)
	hasil = jum(n)
	fmt.Print(hasil*n)
}
func jum(n float64)float64{
	if n == 1{
		return 1
	}else{
		return 1/n + jum(n/2)
	}
}