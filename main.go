package main

import "fmt"
import "time"

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}
func main() {
	print := func(m interface{}) {
		fmt.Println(m)
	}
	slice := make([]string, 3)

	array := [3]int{1, 2, 3}
	fmt.Println("Hello World")
	fmt.Println(time.Now())
	print(array[0])
	fmt.Println("tes:", slice)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	//MAPS
	//make(map[tipo-chave]tipovalor)

	//======================================== RANGE ========================================
	//	kvs := map[string]string{"a": "apple", "b": "banana"}
	// k = key, v = value
	//semelhante ao map do javscript
	for k, v := range m {
		fmt.Printf("%s -> %v\n", k, v)
	}

	//============================= functions ===========================
	// fun nome(parametros) tipo do retorno {}
	plus := sum(1, 2)
	mult := multiply(5, 4)
	fmt.Println(plus, mult)

	// ============== funçao com multiplos retornos ===============

	multiple := func() (int, int) {
		return 3, 4
	}
	a, b := multiple()

	fmt.Println(a, b)

	// Funções que aceitam n quantidade de parametros

	soma := func(nums ...int) int {
		var total int
		for _, num := range nums {
			total += num
		}
		return total
	}

	print(soma(1, 3, 5, 5))

	// =============CLOSURES	==================== uma fuma função que retorna outra função
	intSeq := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	nextInt := intSeq()
	print(nextInt())
	print(nextInt())
	print(nextInt())
	print(nextInt())
	print(nextInt())

	/// ======= RECURSIVIDADE============================
	print("============== recursividade ==========\n")
	print(fatorial(8))

	print("===============PONTEIROS =====================\n")

	zeroVal := func(ival int) {
		ival = 0
	}

	zeroptr := func(ptr *int) {
		*ptr = 0
	}
	i := 9
	zeroVal(i)
	fmt.Println("zeroVal", i)
	zeroptr(&i)
	fmt.Println("ponteiro ", i)

}
