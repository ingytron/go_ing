package main

import "fmt"

type Person struct{
	Name string
}

type Replicant struct{
	Person
	Model string
}


func (p *Person) talk(){
	fmt.Println("Hi, my name is", p.Name)
}

func (r *Replicant) about_me(){
	fmt.Println("Model name", r.Model)
}

func main() {
	a := new(Replicant)
	a.Name = "Deckard"
	a.Person.talk()

	b := new(Replicant)
	b.Name = "Rachel"
	b.Model = "Nexus-7"
	b.Person.talk()
	b.about_me()
}
