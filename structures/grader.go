package structures

const DefaultGrade = 4

type Grader struct {
	First     string   `json:"first"`
	Last      string   `json:"last"`
	Email     string   `json:"email"`
	Grade     int      `json:"grade"`
	Conflicts []string `json:"conflicts"`
}

type Graders struct {
	G *[]Grader `json:"graders"`
}
