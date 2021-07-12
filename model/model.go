package model

import "time"

type ToDoList struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	Status    string  `json:"status"`
}

type AddToDoList struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Attendance struct {
	Id           int    `json:"id"`
	Rfid         string `json:"rfid"`
	EmployeeCode string
	TimeIn       time.Time
	TimeOut      *time.Time
	Date         time.Time
}

type LogAttendance struct {
	Rfid   string `json:"rfid"`
	Status int
}

type MasterKaryawan struct {
	EmployeeCode   string  `json:"employee_code"`
	EmployeeName   string  `json:"employee_name"`
	DivisiId       int     `json:"division_id"`
	PositionId     string  `json:"position_id"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
	RFID           *string `json:"rfid"`
	MasterSalaryId int     `json:"master_salary_id"`
}
