package oids

import (
	"encoding/asn1"
	"testing"
)

// TestGetOID just verifies that GetOID isn't doing anything weird
func TestGetOID(t *testing.T) {
	expects1 := asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 55594}
	result1 := GetOID(nil, "1.3.6.1.4.1.55594")
	if !result1.Equal(expects1) {
		t.Errorf("GetOID(\"%s\") produced an unexpected result: %v", "1.3.6.1.4.1.55594", result1)
	}

	expects2 := asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 55594, 1}
	result2 := GetOID(result1, "1")
	if !result2.Equal(expects2) {
		t.Errorf("GetOID(\"%s\") produced an unexpected result: %v", "1.3.6.1.4.1.55594.1", result2)
	}
}
