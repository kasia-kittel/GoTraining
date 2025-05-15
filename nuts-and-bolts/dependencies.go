package nutsandbolts


import (
	"fmt"
	"io"
	"time"
)

type timer interface {
	Now() time.Time
}

// the io.Writer interface looks like this
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

func Log(t timer, w io.Writer, message string, prefix string) (int, error) {
	dateTime := t.Now().String()
	return fmt.Fprintf(w, "%s - %s: %s", prefix, dateTime, message)
}
