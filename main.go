package oids

import (
	"encoding/asn1"
	"strconv"
	"strings"
)

// RootArc : Cynthia's Root Arc
var RootArc = GetOID(nil, "1.3.6.1.4.1.55594")

// PKIArc : PKI Arc
var PKIArc = GetOID(RootArc, "1")

// PKICPArc : Certificate Policy Arc
var PKICPArc = GetOID(PKIArc, "1")

// PKIExtArc : X.509 Extensions Arc
var PKIExtArc = GetOID(PKIArc, "2")

// PKISDNArc : Subject DN Arc
var PKISDNArc = GetOID(PKIArc, "3")

// PKICPExtValidation : EV Certificate Policy
var PKICPExtValidation = GetOID(PKICPArc, "1")

// PKICPUserAuth : User Authentication Policy
var PKICPUserAuth = GetOID(PKICPArc, "1")

// PKIExtSSHPrincipals : SSH Certificate Principals Extension
var PKIExtSSHPrincipals = GetOID(PKIExtArc, "1")

// PKIExtOwO : OwO Extension
var PKIExtOwO = GetOID(PKIExtArc, "2")

// PKIExtL33T509 : l33t509 Extension
var PKIExtL33T509 = GetOID(PKIExtArc, "1337")

// PKISDNEntityID : Entity ID Subject DN Attribute
var PKISDNEntityID = GetOID(PKISDNArc, "1")

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
