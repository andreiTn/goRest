package apiMessages

type ApiMessage struct {
	Message string`json:message`
}
//implement Stringer interface
func (a ApiMessage) String()  string {
	return a.Message
}