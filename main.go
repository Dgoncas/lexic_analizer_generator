package main

import "fmt"
import "github.com/dgoncas/lexic_analizer_generator/regex"

func test(input string){
	parser := regex.NewRegexParser()
	expresion, _, err := parser.Parse(input)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("Expresion: %v\n", expresion)
}

func main(){
	test(".as.?.d.*(ad(.a(ds)?)+asd)*")
	test("j(a|b|c)")
}