package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func main() {

	var DomainResult, EmailResult bool
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")
	var domain string
	for scanner.Scan() {

		Email := scanner.Text()
		parts := strings.Split(Email, "@")
		if len(parts) == 2 {
			domain = parts[1]
			fmt.Println("Extracted domain:", domain)

		}
		DomainResult = checkDomain(domain)
		EmailResult = checkEmailFormat(Email)
		if DomainResult == true && EmailResult == true {
			fmt.Printf("\nEmail is Good")
		} else {
			fmt.Printf("\nEmail is Bad")

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("ERROR:Could not read from the input %v", err)
	}

}

func checkEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func checkDomain(domain string) bool {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	var ReturnValue = true

	if hasDMARC == true && hasMX == true && hasSPF == true {
		ReturnValue = true
	} else {
		ReturnValue = false
	}

	return ReturnValue
}
