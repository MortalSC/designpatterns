package main

type Class struct {
	Teacher *Teacher
	Student *Student
}

type Teacher struct {
	Name    string
	Class   int // bad: a teacher can teach multiple classes
	Classes []int
}

type Student struct {
	Name  string
	Class int
}

// Bad example of SRP
// This function creates a class with a teacher and a student.
func badCreateClass(teacherName string, studentName string, classRoom int) *Class {
	teacher := &Teacher{
		Name:  teacherName,
		Class: classRoom,
	}
	student := &Student{
		Name:  studentName,
		Class: classRoom,
	}
	class := &Class{
		Teacher: teacher,
		Student: student,
	}
	return class
}

// Good example of SRP
// This function creates a teacher.
func createTeacher(name string, classes []int) *Teacher {
	teacher := &Teacher{
		Name:    name,
		Classes: classes,
	}
	return teacher
}

// This function creates a student.
func createStudent(name string, class int) *Student {
	student := &Student{
		Name:  name,
		Class: class,
	}
	return student
}

// This function creates a class with a teacher and a student.
func createClass(teacher *Teacher, student *Student) *Class {
	class := &Class{
		Teacher: teacher,
		Student: student,
	}
	return class
}

func main() {
	badCreateClass("John", "Doe", 1)

	teacher := createTeacher("John", []int{1, 2})
	student := createStudent("Doe", 1)
	class := createClass(teacher, student)
	println("Teacher:", class.Teacher.Name)
	println("Student:", class.Student.Name)
}
