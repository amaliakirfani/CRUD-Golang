package model

import "time"



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
