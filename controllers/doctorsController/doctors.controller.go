package doctorsController

import (
	models "fita-test-02/models"
	responseRepository "fita-test-02/repositories/responseRepository"
	timezoneRepository "fita-test-02/repositories/timezoneRepository"
	doctorsService "fita-test-02/services/doctorsService"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(Context *gin.Context) {

	datasDoctors, err := doctorsService.ReturnAllDoctors()
	if err != nil {
		arr := make(map[string]interface{})
		arr["message"] = err
		res := responseRepository.Result{Status: 500, Datas: nil, Errors: arr}
		Context.JSON(500, res)
		return
	}

	dt := time.Now()
	// dayName := dt.Weekday()

	var doctorsAvailable []*models.Doctors

	for _, dataDoctor := range datasDoctors {

		locDoctor, errDoctorDate := time.LoadLocation(dataDoctor.TimezoneArea)
		if errDoctorDate != nil {
			continue
		}

		dtDoctor := dt.In(locDoctor)
		dayNameDoctor := dtDoctor.Weekday()
		formatDateDoctorOnly := dtDoctor.Format("2006-01-02")
		// formatDateDoctorPM := dtDoctor.Format("2006-01-02 3:4:5 PM")
		formatDateDoctorRFC822 := dtDoctor.Format(time.RFC822)

		if dayNameDoctor.String() != dataDoctor.DayOfWeek {
			continue
		}

		namedtShanghai := formatDateDoctorOnly[0:10]

		startDoctor, _ := time.Parse("2006-01-02 3:4:5 PM", namedtShanghai+" "+dataDoctor.AvailableAt+":00 "+dataDoctor.AvailableAtMeridiem)
		endDoctor, _ := time.Parse("2006-01-02 3:4:5 PM", namedtShanghai+" "+dataDoctor.AvailableUtil+":00 "+dataDoctor.AvailableUtilMeridiem)

		startDoctorRFC822, _ := time.Parse(time.RFC822, startDoctor.Format(time.RFC822))
		endDoctorRFC822, _ := time.Parse(time.RFC822, endDoctor.Format(time.RFC822))
		inDoctorRFC822, _ := time.Parse(time.RFC822, formatDateDoctorRFC822)

		fmt.Println(startDoctorRFC822)
		fmt.Println(endDoctorRFC822)
		fmt.Println(inDoctorRFC822)

		if inTimeSpan(startDoctorRFC822, endDoctorRFC822, inDoctorRFC822) {

			var doctor models.Doctors
			doctor.ID = dataDoctor.ID
			doctor.Name = dataDoctor.Name
			doctor.TimezoneGmt = dataDoctor.TimezoneGmt
			doctor.TimezoneArea = dataDoctor.TimezoneArea
			doctor.DayOfWeek = dataDoctor.DayOfWeek
			doctor.AvailableAt = dataDoctor.AvailableAt
			doctor.AvailableAtMeridiem = dataDoctor.AvailableAtMeridiem
			doctor.AvailableUtil = dataDoctor.AvailableUtil
			doctor.AvailableUtilMeridiem = dataDoctor.AvailableUtilMeridiem
			doctorsAvailable = append(doctorsAvailable, &doctor)

		}

	}

	res := responseRepository.Result{Status: 200, Datas: doctorsAvailable, Errors: nil}
	Context.JSON(http.StatusOK, res)
	return
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func LocalTime(Context *gin.Context) {

	dt := time.Now()
	dayName := dt.Weekday()

	fmt.Println(dayName)
	fmt.Println(dt.Format("2006-01-02"))
	fmt.Println(dt.Format("3:4"))
	fmt.Println(dt.Format("PM"))
	fmt.Println(dt.Format(time.RFC822))
	fmt.Println(dt.Format("2006-01-02 3:4:5 PM"))
	fmt.Println(dt)

	loc, _ := time.LoadLocation("America/Yakutat")
	dtShanghai := dt.In(loc)
	collegedtShanghai := dtShanghai.String()
	namedtShanghai := collegedtShanghai[0:16]
	fmt.Println(namedtShanghai)
	dayNameShanghai := dtShanghai.Weekday()
	fmt.Println(dayNameShanghai)
	fmt.Println(dtShanghai.Format("2006-01-02"))
	fmt.Println(dtShanghai.Format("3:4"))
	fmt.Println(dtShanghai.Format("PM"))
	fmt.Println(dtShanghai.Format(time.RFC822))
	fmt.Println(dtShanghai.Format("2006-01-02 3:4:5 PM"))
	fmt.Println(dtShanghai.String())

	// dateShanghai, error := time.Parse("2006-01-02", dtShanghai.String())
	// if error != nil {
	// 	fmt.Println(error)
	// }

	// dayNameShanghai := dateShanghai.Weekday()
	// fmt.Println(dayNameShanghai)
	// fmt.Println(dateShanghai.Format("2006-01-02"))
	// fmt.Println(dateShanghai.Format(time.Kitchen))

	timezoneRepository := timezoneRepository.GetTimezone()
	res := responseRepository.Result{Status: 200, Datas: timezoneRepository, Errors: nil}
	Context.JSON(http.StatusOK, res)
	return
}
