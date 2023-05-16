package util

import (
	"fmt"
	"horus/storage"
	"os/exec"
	"regexp"
)

func GetServicesStatus() {

	sNames := storage.GetServicesNames()
	fmt.Println(sNames)
	for _, sName := range sNames {

		cmd := exec.Command("sc", "query", sName)

		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}

		// fmt.Println(string(output))
		serviceInfo := string(output)
		re := regexp.MustCompile(`ESTADO\s+:\s+(\d+)\s+([A-Z_]+)`)
		match := re.FindStringSubmatch(serviceInfo)

		if len(match) > 2 {
			estadoValue := match[2]
			fmt.Printf("ESTADO da aplicação %s: %s\n", sName, estadoValue)
		} else {
			fmt.Printf("Could not find ESTADO field for %s ", sName)
		}
	}
	fmt.Println("\n--------------------------------------------------")
}
