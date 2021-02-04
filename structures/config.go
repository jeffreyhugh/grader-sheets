package structures

type Config struct {
	FileType             string `json:"file_type"`
	DefaultGrade         int    `json:"default_grade"`
	SubmissionsZip       string
	GradersJson          string
	SubmissionsDirectory string
}
