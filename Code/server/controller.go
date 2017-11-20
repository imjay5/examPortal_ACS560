/**************************************************************
 * controller.go
 *
 * Implements handler for different request types from C# client
 *
 * 	                Version History
 *   Author          Type of change          Description
 *   Kanika			 Addition				 Created File

 **************************************************************/

package main

import (
	"encoding/json"
)

type Controller struct {
}

func (controller *Controller) process_data(msg *json.Decoder, handler string) string {
	var response string
	switch handler {
	case "register":
		obj := UserDetails{}
		response = obj.register(msg)

	case "login":
		obj := UserDetails{}
		response = obj.login(msg)

	case "addExam":
		obj := Exam{}
		response = obj.create_exam(msg)

	case "addQuestion":
		obj := Question{}
		response = obj.add_question(msg)

	case "selectExam":
		obj := ExamDetails{}
		response = obj.selectExam(msg)

	case "takeExam":
		obj := QuestionDetails{}
		response = obj.takeExam(msg)

	case "getExam":
		obj := Exam{}
		response = obj.get_exam(msg)

	case "updateExam":
		obj := Exam{}
		response = obj.update_exam(msg)
	case "getQuestion":
		obj := Question{}
		response = obj.get_next_question(msg)
	}
	return response
}
