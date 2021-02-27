package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
  )

func factorial(numero int)int{
	if numero == 0 {
		return 1
	}
	return numero * factorial(numero -1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
  	fmt.Println("Angel Lopez")
	fmt.Println("201807299")
  	fmt.Println("---------------------")

  	for {
		fmt.Println("Ingrese el numero")
    	text0, _ := reader.ReadString('\n')
    	// convert CRLF to LF
		texto := strings.ReplaceAll(text0,"\r\n","")
		i1, err := strconv.Atoi(texto)
		if err == nil {
			fmt.Print("El factorial es ")
			fmt.Print(factorial(i1))
			fmt.Println()
		}else{
			fmt.Println("NO es un numero lo ingresado")
		}
		
	}	
}