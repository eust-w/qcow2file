package src

import "testing"

func TestGenerateUuid(t *testing.T) {
	uuid, err := generateUuid()
	if err != nil {
		t.Errorf("generateUuid error: %v", err)
	}
	if uuid == "" {
		t.Errorf("generateUuid return empty string")
	}
	uuid2, err := generateUuid()
	if err != nil {
		t.Errorf("generateUuid error: %v", err)
	}
	if uuid == uuid2 {
		t.Errorf("generateUuid return same string")
	}
}

func TestSafeUsedUuidMap(t *testing.T) {
	uuid, err := generateUuid()
	if err != nil {
		t.Errorf("generateUuid error: %v", err)
	}
	if safeUsedUuidMapInstance.Get(uuid) != true {
		t.Errorf("safeUsedUuidMapInstance.Get return false")
	}
}
