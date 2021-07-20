package main

import "fmt"

func sayHello(){
	fmt.Println("Hellooo")
}
func sayHelloToSomeone(name string){
	fmt.Println("Hellloooo", name , "👋")
}
func getHello(name string) string{
	return "Hello "+name
}
// klo lbh dari 1 return pake kurung
func getData() (string, int8){
	return "Andrew",20
}
// named return value
func getDataAgain() (name string, age int8){
	name = "Andi"
	age = 20
	return
}
func main() {
	sayHello()
	sayHelloToSomeone("Andrew")
	fmt.Println(getHello("Andi"))
	name, _ := getData()
	a, b := getDataAgain()
	fmt.Println(name, a , b)
}
