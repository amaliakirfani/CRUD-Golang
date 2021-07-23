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

	err := json.NewDecoder(req.Body).Decode(&AttendanceModel)

	if err != nil {
		panic(err)
	}

	pg := infrastructures.AttendanceConn()
	var MasterKaryawanModel model.MasterKaryawan
	rec, err := pg.Select("*").From("master_karyawan").Where("rfid = ? ", AttendanceModel.Rfid).Load(&MasterKaryawanModel)
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

		/*CEK Karyawan Sudah Absen Masuk Atau Belum*/
		CekAttend, err := pg.Select("*").From("attendance").
			Where("employee_code = ?", MasterKaryawanModel.EmployeeCode).
			Where("date = ?", currentTime.Format("2006-01-02")).
			Load(&AttendanceModel)

		if err != nil {
			fmt.Println(err)
		}

		if CekAttend != 0 {
			status = 2
			message = "Absensi Keluar Berhasil!"
			if AttendanceModel.TimeOut == nil {
				_, err := pg.Update("attendance").
					Set("time_out", curTime).
					Where("id = ?", AttendanceModel.Id).Exec()
				if err != nil {
					panic(err)
				}
			} else {
				LogAttendanceModel.Rfid = AttendanceModel.Rfid
				LogAttendanceModel.Status = 3

				_, err := pg.InsertInto("log_attendance").
					Columns("rfid", "status").
					Record(&LogAttendanceModel).
					Exec()

				if err != nil {
					panic(err)
				}

				helpers.Response(res, http.StatusOK, helpers.ResponsePattern{
					Code:    4400,
					Message: "Anda Sudah Melakukan Absen!",
					Status:  "success",
					Data:    []string{},
				})
				return
			}
		} else {
			if curTime < "09:00:00" {
				curTime = "08:00:00"
			}

			status = 1
			message = "Absensi Masuk Berhasil!"

			result, err := pg.InsertInto("attendance").
				Pair("employee_code", MasterKaryawanModel.EmployeeCode).
				Pair("time_in", curTime).
				Pair("date", currentTime.Format("2006-01-02")).
				Exec()

			if err != nil {
				panic(err)
			}
			fmt.Println(result)
		}
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
