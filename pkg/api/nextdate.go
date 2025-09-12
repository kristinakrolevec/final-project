package api

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	return date.After(now)

}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if dstart == "" || repeat == "" {
		return "", errors.New("NextDate isn't")
	}
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		log.Println("Проверка dstart на time.Parse не пройдена")
		return "", err
	}
	log.Printf("значение dstart = %v/", date)

	sliceInterval := strings.Split(repeat, " ")

	if sliceInterval[0] == "d" {
		interval, err := strconv.Atoi(sliceInterval[len(sliceInterval)-1])

		if err != nil || interval > 400 {
			return "", err
		}
		log.Printf("значение d = %d/", interval)

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}

		}

	} else if sliceInterval[0] == "y" {

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}

		}
	} else {
		log.Println("Повторение не назначено")
		return "", errors.New("Unsupported format")
	}
	log.Println("все равно продолжает работать")
	return date.Format(DateFormat), nil
}

func nextDayHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Обработчик nextDayHandler")

	nowGet := req.URL.Query().Get("now")

	nowTime, err := time.Parse(DateFormat, nowGet)
	if err != nil {
		http.Error(res, "Ошибка при парсинге даты", http.StatusBadRequest)
		return
	}

	dstart := req.URL.Query().Get("date")
	repeat := req.URL.Query().Get("repeat")

	log.Printf("now = %v, dstart = %s, repeat = %s", nowTime, dstart, repeat)

	nextDateStr, err := NextDate(nowTime, dstart, repeat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, nextDateStr)

}
