package maps

import (
	"fmt"
	"time"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
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

func TotalMemoryToGib(mem int) float64 {
	memGib := float64(mem) / 1024
	memGib = memGib / 1024
	return memGib
}

func MapDaysToColorMark(daysRemaining int) string {
	if daysRemaining > 180 {
		return "🟩"
	} else if daysRemaining > 30 {
		return " 🟨"
	} else if daysRemaining > 0 {
		return "🟧"
	} else {
		return "🟥"
	}
}

func GetOrg(orgs []count.Msg, orgFilter string) count.Msg {
	for _, org := range orgs {
		if orgFilter == org.Organization.Name {
			return org
		}
	}
	return count.Msg{}
}