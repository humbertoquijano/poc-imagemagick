package imgkcaller

import (
	"log"
	"os/exec"
	"sync"
)

func MakeCall(batPath string, wg *sync.WaitGroup) {
	cmd := exec.Command("CMD", "/C", batPath)
	log.Printf("convirtiendo...")

	if err := cmd.Run(); err != nil {
		log.Printf("Error: %v", err.Error())
	}

	defer wg.Done()
}
