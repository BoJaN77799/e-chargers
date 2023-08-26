package utils

import "time"

type DateFromTo struct {
	DateFrom time.Time
	DateTo   time.Time
}

func ConvertDateFromAndDateTo(dateFrom string, dateTo string) (DateFromTo, error) {
	dateFromParsed, err := time.Parse(time.RFC3339, dateFrom)
	if err != nil {
		return DateFromTo{}, err
	}

	dateToParsed, err := time.Parse(time.RFC3339, dateTo)
	if err != nil {
		return DateFromTo{}, err
	}

	return DateFromTo{
		DateFrom: dateFromParsed,
		DateTo:   dateToParsed,
	}, nil
}
