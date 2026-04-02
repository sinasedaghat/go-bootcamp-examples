package main

import (
	"fmt"
	"os"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

func formatSize(size int64) string {
	switch {
	case size >= PB:
		return fmt.Sprintf("%.2f PB", float64(size)/PB)
	case size >= TB:
		return fmt.Sprintf("%.2f TB", float64(size)/TB)
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func monitorDir() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter directory location")
		return
	}

	entries, _ := os.ReadDir(args[0])
	var dirs [2]int
	var files [2]int

	for _, entry := range entries {
		if entry.IsDir() {
			dirs[0]++
			dirs[1] += len(entry.Name()) + 1 // name byte + \n
		} else {
			files[0]++
			files[1] += len(entry.Name()) + 1
		}
	}

	d := make([]byte, 0, dirs[1])
	f := make([]byte, 0, files[1])

	for _, entry := range entries {
		info, _ := entry.Info()

		if entry.IsDir() {
			fmt.Printf(
				"%s is directory\n",
				entry.Name(),
			)
			d = append(d, entry.Name()...)
			d = append(d, '\n')
		} else {
			fmt.Printf(
				"%s file has %s data.\n",
				entry.Name(), formatSize(info.Size()),
			)
			f = append(f, entry.Name()...)
			f = append(f, '\n')
		}
	}

	t, err := title(dirs[0], files[0], args[0])
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	writeFile([]byte(t), d, f)
}

func title(dc, fc int, dir string) (string, error) {
	wd, err := os.Getwd() // working directory
	if err != nil {
		return "", err
	}

	if dir != "." {
		wd = fmt.Sprintf("%s/%s", wd, dir)
	}

	return fmt.Sprintf(
		"In the \"%s\" path, there are %d directories and %d files.",
		wd, dc, fc,
	), nil
}

func writeFile(t, d, f []byte) {
	data := make([]byte, 0, len(t)+len(d)+len(f)+34)
	data = append(data, t...)
	data = append(data, "\n\nDirectories names:\n"...)
	data = append(data, d...)
	data = append(data, "\nFiles names:\n"...)
	data = append(data, f...)

	err := os.WriteFile("output.txt", data, 0644)
	if err != nil {
		fmt.Println("Err: ", err)
	} else {
		fmt.Println("\"output.txt\" WAS CREATED")
	}
}
