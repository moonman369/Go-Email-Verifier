package emailverifier

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type VerificationResult struct {
	Domain      string `json:"domain"`
	HasMX       bool   `json:"hasMX"`
	HasSPF      bool   `json:"hasSPF"`
	SPFRecord   string `json:"spfRecord"`
	HasDMARC    bool   `json:"hasDMARC"`
	DMARCRecord string `json:"dmarcRecord"`
}

func CheckDomain(domain string) VerificationResult {
	var ver VerificationResult

	ver.Domain = domain

	// Checking MX Records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		ver.HasMX = true
	}

	// Checking SPF Records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			ver.HasSPF = true
			ver.SPFRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			ver.HasDMARC = true
			ver.DMARCRecord = record
			break
		}
	}

	fmt.Printf("Domain:\t\t%v\n", ver.Domain)
	fmt.Printf("Has MX:\t\t%v\n", ver.HasMX)
	fmt.Printf("Has SPF:\t%v\n", ver.HasSPF)
	fmt.Printf("SPF Record:\t%v\n", ver.SPFRecord)
	fmt.Printf("Has DMARC:\t%v\n", ver.HasDMARC)
	fmt.Printf("DMARC Record:\t%v\n", ver.DMARCRecord)
	fmt.Println()

	return ver
}
