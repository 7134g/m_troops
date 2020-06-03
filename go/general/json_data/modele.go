package json_data

type Grade struct {
	School string `json:"school_name"`
	Grade  int    `json:"grade"`
}

type Student struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

type Teacher struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Grade   Grade     `json:"grade_info"`
	Student []Student `json:"student_info"`
}
