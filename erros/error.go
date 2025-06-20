package erros
type CustomError struct{
	Err error
	Message string
	Type string
}
func (c CustomError) Error() string {
return "Message:"+c.Message+"type"+c.Type
}
func Test()error{
return CustomError{
	Err: nil,
	Message: "test",
	Type: "test",
}
}