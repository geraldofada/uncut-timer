package timer

import (
	"encoding/gob"
	"os"
)

// Save stores a list of timers in a bin file
func Save(t []*Timer, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		return err
	}

	enc := gob.NewEncoder(file)

	err = enc.Encode(t)
	if err != nil {
		return err
	}

	return nil
}

// Get returns a struct timer from a given bin file
func Get(id int, path string) (*Timer, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var t []*Timer

	stat, _ := file.Stat()
	if stat.Size() > 0 {
		dec := gob.NewDecoder(file)
		err = dec.Decode(&t)
		if err != nil {
			return nil, err
		}
	}

	return t[id], nil
}

// Read returns a list of Timers
func Read(path string) ([]*Timer, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var t []*Timer

	stat, _ := file.Stat()
	if stat.Size() > 0 {
		dec := gob.NewDecoder(file)
		err = dec.Decode(&t)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

// Remove removes a Timer from a given bin file
func Remove(id int, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return nil
	}

	var t []*Timer

	stat, _ := file.Stat()
	if stat.Size() > 0 {
		dec := gob.NewDecoder(file)
		err = dec.Decode(&t)
		if err != nil {
			return err
		}

		t = append(t[:id], t[id+1:]...)

		err = Save(t, path)
		if err != nil {
			return err
		}
	}

	return nil
}
