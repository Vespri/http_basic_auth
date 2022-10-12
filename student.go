package main

var students = []*Student{}

type Student struct {
	ID    string
	Name  string
	Grade int32
}

func GetStudent() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.ID == id {
			return each
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{ID: "s001", Name: "Kresna", Grade: 1})
	students = append(students, &Student{ID: "s002", Name: "Taufiiq", Grade: 2})
	students = append(students, &Student{ID: "s003", Name: "Rian", Grade: 3})
}
