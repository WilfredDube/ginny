package http

import "time"

func HumanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func HumanDateShort(t time.Time) string {
	return t.Format("02 Jan 2006")
}

func TrimGuid(g string) string {
	return g[:8]
}
