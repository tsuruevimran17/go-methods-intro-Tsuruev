package main

import "fmt"

type Passport struct {
	FirstName string
	lastName  string
	Active    bool
}

func (r Passport) FullName() string {
	return r.FirstName + r.lastName
}

func (r *Passport) Deactivate() bool {
	r.Active = false
	return r.Active
}

func (r Passport) IsActive() bool {
	return r.Active

}

func main() {
	name := Passport{
		FirstName: "Imran",
		lastName:  "Tsuruev",
	}
	fmt.Println(name.FullName(), name.Deactivate(), name.IsActive())
}
