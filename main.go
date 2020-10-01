package oids

import (
	"encoding/asn1"
	"strconv"
	"strings"
)

// RootArc : Cynthia's root OID
var RootArc = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 55594}

// PKISDNEntityID : PKI Subject DN Entity ID
var PKISDNEntityID = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 55594, 1, 3, 1}

// GetOID creates an OID object from a dotted string
func GetOID(oidBase asn1.ObjectIdentifier, oidStr string) asn1.ObjectIdentifier {
	parts := strings.Split(oidStr, ".")
	var result asn1.ObjectIdentifier

	if oidBase != nil {
		result = oidBase
	}

	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}

	return result
}
