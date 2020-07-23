package message

// Subject abc
type Subject int32

// Constants for Subject
const (
	SubjectPHYSICS   Subject = 0
	SubjectCHEMISTRY Subject = 1
	SubjectMATHS     Subject = 2
)

// MarkReq abc
type MarkReq struct {
	SlNo    int32
	Name    string
	Subject Subject
}
