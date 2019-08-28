package main

import (
	"testing"
)

func TestDeploymentFromTemplate(t *testing.T) {
	d, err := DeploymentFromTemplate("{\"a\":\"b\"}")
	if err != nil {
		t.Errorf("expected empty error instead %s", err.Error())
	}
	found, exists := d["a"]
	if !exists {
		t.Errorf("expected exists %t instead %t", exists, true)
	}
	if found.(string) != "b" {
		t.Errorf("expected found '%s' instead %s", found, "b")
	}
}
