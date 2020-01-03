package cmd

import (
	"errors"
	"fmt"

	"github.com/geraldofada/uncut-timer/timer"
)

// CliStart is the function running on the start command
func CliStart(name string) error {
	timers, err := timer.Read("ongoing")
	if err != nil {
		return err
	}

	newTimer := timer.Start(len(timers), name)

	timers = append(timers, newTimer)

	err = timer.Save(timers, "ongoing")
	if err != nil {
		return err
	}

	if newTimer.Name != "" {
		fmt.Printf("Timer [%d]: %s, started at: %s\n", newTimer.ID, newTimer.Name, newTimer.Start)
	} else {
		fmt.Printf("Timer [%d] started at: %s\n", newTimer.ID, newTimer.Start)
	}

	return nil
}

// CliList is the function running on the list command
func CliList(id int, path string) error {
	timers, err := timer.Read(path)
	if err != nil {
		return err
	}

	if id != -1 && id < len(timers) {
		timers = timers[id : id+1]
	} else if id >= len(timers) {
		return errors.New("This timer does not exists")
	}

	for _, timer := range timers {
		if timer.Name != "" {
			fmt.Printf("Timer [%d]: %s\n", timer.ID, timer.Name)
		} else {
			fmt.Printf("Timer [%d]\n", timer.ID)
		}

		fmt.Printf("\tStarted at: %s\n", timer.Start)

		if timer.Stopped {
			fmt.Printf("\tStopped at: %s\n", timer.End)
			fmt.Printf("\tTotal time: %s\n", timer.Elapsed)
		}
	}
	return nil
}
