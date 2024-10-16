package data

type Employee struct {
	Id   int
	Age  int
	City string
	Name string
}

// MEDLEMSFUNKTION
// emp.CalculateSalary()
func (emp Employee) CalculateSalary() int {
	if emp.Name == "Stefan" {
		return 1000
	}
	return 10
}
