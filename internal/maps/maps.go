package maps

import (
	"fmt"
	"time"
)

func TimeNowToFormattedDate() string {
	monthMap := map[time.Month]string{
		time.January:   "Janeiro",
		time.February:  "Fevereiro",
		time.March:     "Março",
		time.April:     "Abril",
		time.May:       "Maio",
		time.June:      "Junho",
		time.July:      "Julho",
		time.August:    "Agosto",
		time.September: "Setembro",
		time.October:   "Outubro",
		time.November:  "Novembro",
		time.December:  "Dezembro",
	}

	now := time.Now()
	return fmt.Sprintf("%d de %s de %d", now.Day(), monthMap[now.Month()], now.Year())
}