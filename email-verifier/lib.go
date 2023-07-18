package emailverifier

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func CheckDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Checking MX Records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Checking SPF Records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
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
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("Domain:\t\t%v\n", domain)
	fmt.Printf("Has MX:\t\t%v\n", hasMX)
	fmt.Printf("Has SPF:\t%v\n", hasSPF)
	fmt.Printf("SPF Record:\t%v\n", spfRecord)
	fmt.Printf("Has DMARC:\t%v\n", hasDMARC)
	fmt.Printf("DMARC Record:\t%v\n", dmarcRecord)
	fmt.Println()
}
