package employee

type EmployeeRepository interface {
	Create(*Employee) *Employee
	FindByID(*int64) (Employee, error)
	FindByName(string) (Employee, error)
}
