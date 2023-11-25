package main

import (
	"context"
	"github.com/ekomobile/dadata/v2/client"
	"reflect"
	"testing"
)

func TestAddCredential(t *testing.T) {
	main()
	tests := []struct {
		name        string
		credentials client.Credentials
	}{
		{"Test get credentials", NewClientCredentials()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedCredentials := client.Credentials{
				ApiKeyValue:    "86ba621e7a8aa0efa0fd8057cd4e889d734cbbd8",
				SecretKeyValue: "9a72b95eb027c3120069c6da0755a7fcfd3299d9",
			}
			AddCredential(expectedCredentials)
			api := NewWorkApi()
			_, err := api.Address(context.Background(), "москва сухонская 11")
			if err != nil {
				t.Errorf("Function %s does not work", tt.name)
			}
		})
	}
}

func TestNewClientCredentials(t *testing.T) {
	tests := []struct {
		name string
		want client.Credentials
	}{
		{"Test New Credentials", client.Credentials{
			ApiKeyValue:    "86ba621e7a8aa0efa0fd8057cd4e889d734cbbd8",
			SecretKeyValue: "9a72b95eb027c3120069c6da0755a7fcfd3299d9",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientCredentials(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWorkApi(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test NewWorkApi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewWorkApi()
			_, err := api.Address(context.Background(), "москва сухонская 11")
			if err != nil {
				t.Errorf("Function %s does not work", tt.name)
			}
		})
	}
}
