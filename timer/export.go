package timer

import (
	"encoding/csv"
	"os"
	"time"
)

// Export creates a new csv file on the given
// path with the finished timers
func Export(t []*Timer, path string) error {
	date := time.Now().Format("20060102T150405")
	path += "-" + date + ".csv"

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fieldsName := []string{
		"Name",
		"Started at",
		"Stopped at",
		"Total time",
	}
	err = writer.Write(fieldsName)
	if err != nil {
		return err
	}

	for _, timer := range t {
		data := []string{
			timer.Name,
			timer.Start.String(),
			timer.End.String(),
			timer.Elapsed.String(),
		}

		err := writer.Write(data)
		if err != nil {
			return err
		}

	}

	return nil
}
