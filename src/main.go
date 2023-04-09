package main

import (
	"flag"
	"os"
)

func memscan(pid string) {
	os.Open("/proc/" + pid)

}

func main() {
	pid := flag.String("pid", "", "Process ID of proc memory to scan -- likely the webserver process")
	flag.Parse()
	// content, _ := ioutil.ReadDir("/proc/")

	// for _, file := range content {
	// 	fmt.Println(file.Name())
	// }

}
