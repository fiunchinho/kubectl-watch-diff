package main

import (
	"encoding/json"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

func checkErr(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func writeYaml(data []byte) (*os.File, error) {
	f, err := os.CreateTemp("", "kubectl-watch-diff*.yaml")
	if err != nil {
		return f, err
	}
	if _, err := f.Write(data); err != nil {
		return f, err
	}
	return f, nil
}

func closeAndDelete(f *os.File) {

}

func main() {
	if len(os.Args) < 3 {
		println("Usage: " + os.Args[0] + " <resource> <name>")
		os.Exit(1)
	}
	cmd := exec.Command("kubectl", "get", "-w", os.Args[1], os.Args[2], "-o=json")
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	checkErr(err)

	checkErr(cmd.Start())

	var prev []byte
	dec := json.NewDecoder(stdout)
	first := true
	for dec.More() {
		var obj map[string]interface{}
		checkErr(dec.Decode(&obj))
		data, err := yaml.Marshal(obj)
		checkErr(err)
		if prev == nil {
			prev = data
		} else {
			if first {
				first = false
			} else {
				_, _ = os.Stdout.WriteString("===\n")
			}

			prevFile, err := writeYaml(prev)
			checkErr(err)
			nextFile, err := writeYaml(data)
			checkErr(err)
			cmd := exec.Command("diff", "-u", prevFile.Name(), nextFile.Name())
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
			closeAndDelete(prevFile)
			closeAndDelete(nextFile)
		}
	}

	checkErr(cmd.Wait())
}
