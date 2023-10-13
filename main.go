package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	cmd := exec.Command("kubectl", "config", "get-contexts", "--output=name")
	stdout, err := cmd.Output()
	clusternames := strings.Split(strings.TrimSuffix(string(stdout), "\n"), "\n")
	for _, name := range clusternames {
		if strings.HasPrefix(name, "aks-pr-") || strings.HasPrefix(name, "aks-non") {
			continue
		}
		preCommand := append([]string{"--context", name}, args...)
		fmt.Printf("Clustername: %s\n", name)
		cmd = exec.Command("kubectl", preCommand...)
		stdout, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(stdout))

	}
	if err != nil {
		fmt.Println(err.Error())
	}

}
