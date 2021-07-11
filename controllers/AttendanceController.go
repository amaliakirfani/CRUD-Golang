package controllers

import (
	"AttendanceApi/helpers"
	"AttendanceApi/infrastructures"
	"AttendanceApi/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Attendance(res http.ResponseWriter, req *http.Request) {

	var LogAttendanceModel model.LogAttendance
	var AttendanceModel model.Attendance

	var currentTime = time.Now()
	var curTime = currentTime.Format("15:04:05")

	// if curTime < "17:00:00" {

	// 	helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
	// 		Code:    2200,
	// 		Message: "Absen Tidak Boleh Kurang Dari Jam 8",
	// 		Status:  "success",
	// 		Data:    []string{},
	// 	})
	// }

	err := json.NewDecoder(req.Body).Decode(&AttendanceModel)

	if err != nil {
		panic(err)
	}

	pg := infrastructures.AttendanceConn()
	var MasterKaryawanModel model.MasterKaryawan
	rec, err := pg.Select("*").From("master_karyawan").Where("id = ? ", 6).Load(&MasterKaryawanModel)
	fmt.Println("master", MasterKaryawanModel)
	if err != nil {
		panic(err)
	}
	var status int
	var message string

	if rec == 0 {
		status = 0
		message = "ID Karyawan Tidak Terdaftar!"
	} else {
		status = 1
		message = "Absensi Berhasil!"
		result, err := pg.InsertInto("attendance").
			Pair("employee_code", MasterKaryawanModel.KodeKaryawan).
			Pair("time_in", curTime).
			Pair("date", currentTime.Format("2006-01-02")).
			Exec()

		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}

	fmt.Println(status)

	LogAttendanceModel.Rfid = AttendanceModel.Rfid
	LogAttendanceModel.Status = status

	result, err := pg.InsertInto("log_attendance").
		Columns("rfid", "status").
		Record(&LogAttendanceModel).
		Exec()

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(LogAttendanceModel)

	helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
		Code:    2200,
		Message: message,
		Status:  "success",
		Data:    LogAttendanceModel.Rfid,
	})
}
