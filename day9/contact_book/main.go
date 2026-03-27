package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	City  string
}

func main() {
	book := make(map[string]Contact)

	for {
		fmt.Println("1=Add 2=Search 3=List 4=Quit")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, phone, city string
			fmt.Print("Name: ")
			fmt.Scan(&name)
			fmt.Print("Phone: ")
			fmt.Scan(&phone)
			fmt.Print("City: ")
			fmt.Scan(&city)
			book[name] = Contact{name, phone, city}
			fmt.Println("Saved!")
		case 2:
			var name string
			fmt.Print("Search name: ")
			fmt.Scan(&name)
			c, ok := book[name]
			if ok {
				fmt.Printf("%s | %s | %s",
					c.Name, c.Phone, c.City)
			} else {
				fmt.Println("Not found!")
			}
		case 3:
			for _, c := range book {
				fmt.Printf("%s | %s | %s",
					c.Name, c.Phone, c.City)
			}
		case 4:
			fmt.Println("Bye!")
			return
		}
	}
}
