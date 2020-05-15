package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/geraldofada/uncut-timer/timer"
)

// OngoingFileName is a constant for the ongoing file name
const OngoingFileName = "ongoing"

// FinishedFileName is a constant for the finished file name
const FinishedFileName = "finished"

// ExportedFileName is a constant for the exported file name
const ExportedFileName = "exported"

// CliStart is the function running on the start command
func CliStart(name string) error {
	timers, err := timer.Read(OngoingFileName)
	if err != nil {
		return err
	}

	newTimer := timer.Start(name)

	timers = append(timers, newTimer)

	err = timer.Save(timers, OngoingFileName)
	if err != nil {
		return err
	}

	if newTimer.Name != "" {
		fmt.Printf("[%d]: %s, started at: %s\n", len(timers)-1, newTimer.Name, newTimer.Start)
	} else {
		fmt.Printf("Timer [%d] started at: %s\n", len(timers)-1, newTimer.Start)
	}

	return nil
}

// CliList is the function running on the list command
func CliList(id int, path string) error {
	today := time.Now()

	timers, err := timer.Read(path)
	if err != nil {
		return err
	}

	if id != -1 && id < len(timers) {
		timers = timers[id : id+1]
	} else if id >= len(timers) {
		return errors.New("This timer does not exists")
	}

	for i, timer := range timers {
		if timer.Name != "" {
			fmt.Printf("[%d]: %s\n", i, timer.Name)
		} else {
			fmt.Printf("Timer [%d]\n", i)
		}

		fmt.Printf("\tStarted at: %s\n", timer.Start)

		if timer.Stopped {
			fmt.Printf("\tStopped at: %s\n", timer.End)
			fmt.Printf("\tTotal time: %s\n", timer.Elapsed)
		} else {
			fmt.Printf("\tElapsed time: %s\n", today.Sub(timer.Start))
		}
	}
	return nil
}

// CliStop is the function running on the stop command
func CliStop(id int) error {
	ongoing, err := timer.Read(OngoingFileName)
	if err != nil {
		return nil
	}

	if id >= len(ongoing) {
		return errors.New("This timer does not exists")
	}

	finished, err := timer.Read(FinishedFileName)
	if err != nil {
		return err
	}

	err = timer.Remove(id, OngoingFileName)
	if err != nil {
		return err
	}

	timer.Stop(ongoing[id])

	finished = append(finished, ongoing[id])

	err = timer.Save(finished, FinishedFileName)
	if err != nil {
		return err
	}

	if ongoing[id].Name != "" {
		fmt.Printf("[%d]: %s, stopped at: %s\n", id, ongoing[id].Name, ongoing[id].End)
	} else {
		fmt.Printf("Timer [%d] stopped at: %s\n", id, ongoing[id].End)
	}

	return nil
}

// CliRemove is the function running on cmd remove
func CliRemove(id int, path string) error {
	timers, err := timer.Read(path)
	if err != nil {
		return err
	}

	if id >= len(timers) {
		return errors.New("This timer does not exists")
	}

	err = timer.Remove(id, path)
	if err != nil {
		return err
	}

	if timers[id].Name != "" {
		fmt.Printf("[%d]: %s, removed\n", id, timers[id].Name)
	} else {
		fmt.Printf("Timer [%d] removed\n", id)
	}

	return nil
}

// CliExport is the function running on cmd export
func CliExport(path string) error {
	timers, err := timer.Read(FinishedFileName)
	if err != nil {
		return err
	}

	err = timer.Export(timers, path)
	if err != nil {
		return err
	}

	return nil
}
