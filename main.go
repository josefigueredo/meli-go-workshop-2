package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
 3º Ejercicio (struct, punteros, slices)
  Hacer un programa que reciba por parametro una lista de enteros, los almacene en un slice, popule un arbol binario y
  despues imprima los valores en orden ascendente
*/

type tree struct {
	value int
	left, right *tree
}

func add(t *tree, v int) *tree {
	if t == nil {
		return &tree{value: v}
	}
	if v < t.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

func printAsc(t *tree) {
	if t == nil {
		return
	}
	printAsc(t.left)
	fmt.Println(t.value)
	printAsc(t.right)
}

func main() {

	if len(os.Args) == 0 {
		fmt.Println("Error: Se requieren argumentos de tipo int")
		os.Exit(1)
	}

	arr := make([]int,0)
	for k, v := range os.Args {
		if k != 0 {
			fmt.Println("El argumento de la posicion: " + strconv.Itoa(k) + " es: " + v)
			vi, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error: El argumento " + v + " no es casteable a int")
				os.Exit(1)
			}
			arr = append(arr, vi)
		}
	}

	fmt.Println(arr)

	var t *tree;
	for _, v := range arr {
		t = add(t, v)
	}

	printAsc(t)
}