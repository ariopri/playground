package main

import "fmt"

// Check Point:
// Search Roles
// - Input: Role
// - Output: Rearch Result

// Contoh:
// Data Users: [{Aditira 20 Programmer} {Dimas 18 Designer} {Yuni 21 DevOps} {Dito 22 Programmer} {Juno 25 DevOps}]

// Input:
//   - Masukan Role : Programmer

// Output:
// Programmer Found:
// Name:  Aditira  Age:  20  Role:  Programmer
// Name:  Dito  Age:  22  Role:  Programmer

// Input:
//   - Masukan Role : Secretary

// Output:
// Role: Sercretary Not Found!

type User struct {
	name string
	age  int
	role string
}

func main() {

	users := []User{
		{
			name: "Aditira",
			age:  20,
			role: "Programmer",
		},
		{
			name: "Dimas",
			age:  18,
			role: "Designer",
		},
		{
			name: "Yuni",
			age:  21,
			role: "DevOps",
		},
		{
			name: "Dito",
			age:  22,
			role: "Programmer",
		},
		{
			name: "Juno",
			age:  25,
			role: "DevOps",
		},
	}

	// TODO: answer here
	for {
		fmt.Println("Data Users:")
		for _, u := range users {
			fmt.Println("Name: ", u.name, " Age", u.age, "Role: ", u.role)
		}
		
		var name, role string
		var age int
		fmt.Println("Add Name : ")
		fmt.Scan(&name)
		fmt.Println("Add Age : ")
		fmt.Scan(&age)
		fmt.Println("Add Role : ")
		fmt.Scan(&role)

		var userStruct = User{
			name: name,
			age:     age,
			role:     role,
		}

		users = append(users, userStruct)
		fmt.Println("Users Added:", users)

		var addAgain string
		fmt.Printf("Ingign menambahkan user lain?(yes/no):")
		fmt.Scan(&addAgain)
		
		if addAgain == "no" {
			fmt.Println("Data Users Update: ")
			for _,u := range users {
				fmt.Println("Name:", u.name, "Age", u.age, " Role:", u.role)
			}
			break		
		}
	}
}