package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("ERROR:Could not read from the input %v", err)
	}

}

func checkDomain(Domain string) {

}
