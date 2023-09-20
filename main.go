package main

import (
	"fmt"
)

type Person struct {
	Name          string
	Relationships map[string][]*Person
}

type FamilyTree struct {
	People map[string]*Person
}

func CreatePerson(name string) *Person {
	person := &Person{
		Name:          name,
		Relationships: make(map[string][]*Person),
	}
	return person
}

func (p *Person) AddRelationship(relationship string, person *Person) {
	p.Relationships[relationship] = append(p.Relationships[relationship], person)
}

func (p *Person) GetRelationships(relationship string) []*Person {
	return p.Relationships[relationship]
}

func NewFamilyTree() *FamilyTree {
	familyTree := FamilyTree{
		People: make(map[string]*Person),
	}
	return &familyTree
}

func (ft *FamilyTree) AddPerson(name string) {
	person := CreatePerson(name)
	ft.People[name] = person
}

func (ft *FamilyTree) Connect(relationship, name1, name2 string) {
	person1, person2 := ft.People[name1], ft.People[name2]
	if person1 != nil && person2 != nil {
		person1.AddRelationship(relationship, person2)
	}
	if relationship == "wife" {
		sons := person1.GetRelationships("son")
		daughter := person1.GetRelationships("daughter")
		for _, son := range sons {
			person2.AddRelationship("son", son)
		}
		for _, daughter := range daughter {
			person2.AddRelationship("daughter", daughter)
		}
	}
}

func (ft *FamilyTree) CountSons(name string) int {
	person := ft.People[name]
	if person != nil {
		sons := person.GetRelationships("son")
		return len(sons)
	}
	return 0
}

func (ft *FamilyTree) CountDaughters(name string) int {
	person := ft.People[name]
	if person != nil {
		daughters := person.GetRelationships("daughter")
		return len(daughters)
	}
	return 0
}

func (ft *FamilyTree) CountWives(name string) int {
	person := ft.People[name]
	if person != nil {
		wives := person.GetRelationships("wife")
		return len(wives)
	}
	return 0
}

func main() {
	ft := NewFamilyTree()

	for {
		var command string
		fmt.Print("family-tree> {add/connect/count/get}> ")
		fmt.Scanln(&command)

		switch command {
		case "add":
			var name string
			fmt.Print("family-tree add person {name}> ")
			fmt.Scanln(&name)
			ft.AddPerson(name)

		case "connect":
			var name1, relationship, name2 string
			fmt.Print("family-tree connect {name1} as {relationship} of {name2}> ")
			fmt.Scanln(&name1, &relationship, &name2)
			if relationship == "father" {
				ft.Connect(relationship, name2, name1)
				ft.Connect("son", name1, name2)

			} else if relationship == "son" {
				ft.Connect(relationship, name2, name1)
				ft.Connect("father", name1, name2)

			} else if relationship == "daughter" {
				ft.Connect(relationship, name2, name1)
				ft.Connect("father", name1, name2)

			} else if relationship == "wife" {
				ft.Connect(relationship, name2, name1)
				ft.Connect("husband", name1, name2)

			} else if relationship == "husband" {
				ft.Connect(relationship, name2, name1)
				ft.Connect("wife", name1, name2)
			}

		case "count":
			var relationship, name string
			fmt.Print("family-tree count {sons/daughters/wives} of {name}> ")
			fmt.Scanln(&relationship, &name)

			if relationship == "sons" {
				fmt.Println("Number of sons: ", ft.CountSons(name))
			} else if relationship == "daughters" {
				fmt.Println("Number of daughters: ", ft.CountDaughters(name))
			} else if relationship == "wives" {
				fmt.Println("Number of wives: ", ft.CountWives(name))
			}

		case "get":
			var relationship, name string
			fmt.Print("family-tree get {son/daughter/wife/father} of {name}> ")
			fmt.Scanln(&relationship, &name)
			relative := ft.People[name].GetRelationships(relationship)
			for _, person := range relative {
				fmt.Print(person.Name, " ", "\n")
			}

		case "exit":
			return
		}
	}
}
