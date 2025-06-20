package main

import (
	"errors"
	"learning/learning/erros"
	usersmodule "learning/learning/users_module"
	"log"
	"strconv"
)

func main() {
	println("Ritnesh Thakur ")
	_,err:=DivideWithError(1,0)
	if err!=nil {
		println(err.Error())
	}
	divideResuly,_:=DivideWithError(30,10)
	println(*divideResuly)
	user:=usersmodule.NewUser("Ritnesh",20,"Thimi",true)
	log.Println(user)
	_,err=parseString("abcd")
		if err!=nil {
			println(err.Error())
		}
	_,err=parseString("20000")
	result,_:=parseString("100")
	println(*result)
	
}
func parseString(value string)(*int,error){
	res,err:=strconv.Atoi(value)
	if  err!=nil {
		return nil,erros.CustomError{
			Message:"Cannot ",
			Type:"type",
			Err:err,
		}
	}
	if res<1 ||res>1000 {
		return nil,erros.CustomError{
			Message:"Value must be between 1 and 10000 ",
			Type:"range(1-1000)",
			Err:err,
		}
	}
	return &res,nil;
}
func DivideWithError(a, b float64) (*float64, error) {
	if b == 0 {
		return nil, errors.New("cannot divide by zero")
	}
	res:=a/b
	return &res,nil
}