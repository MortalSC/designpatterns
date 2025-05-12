#include <iostream>
#include <string>
#include <vector>

using namespace std;

struct Teacher {
    string name;    
    vector<int> classes;
};

struct Student {
    string name;  
    int class_num;  
};

struct Class {
    Teacher* teacher;  
    Student* student; 
};

Teacher* createTeacher(const string& name, const vector<int>& classes) {
    return new Teacher{name, classes};
}

Student* createStudent(const string& name, int cls) {
    return new Student{name, cls};
}

Class* createClass(Teacher* teacher, Student* student) {
    return new Class{teacher, student};
}

int main() {
    Teacher* teacher = createTeacher("John", {1, 2});
    Student* student = createStudent("Doe", 1);
    Class* cls = createClass(teacher, student);

    cout << "Teacher: " << cls->teacher->name << endl;
    cout << "Student: " << cls->student->name << endl;

    // free memory
    // In a real-world scenario, you would use smart pointers or proper memory management
    delete cls->teacher;
    delete cls->student;
    delete cls;

    return 0;
}