package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var newTable string
	if table < 10 {
		newTable = "00"
	} else if table < 100 {
		newTable = "0"
	}
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s%v. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, newTable, table, direction, distance, neighbor)
}
