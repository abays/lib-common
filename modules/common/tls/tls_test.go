/*
Copyright 2023 Red Hat

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tls

import (
	"testing"
)

func TestCreateVolumeMounts(t *testing.T) {
	tests := []struct {
		name          string
		service       *Service
		ca            *Ca
		wantMountsLen int
	}{
		{
			name:          "No Secrets",
			service:       &Service{},
			ca:            &Ca{},
			wantMountsLen: 0,
		},
		{
			name:          "Only TLS Secret",
			service:       &Service{SecretName: "test-tls-secret"},
			ca:            &Ca{},
			wantMountsLen: 2,
		},
		{
			name:          "Only CA Secret",
			service:       &Service{},
			ca:            &Ca{CaSecretName: "test-ca1"},
			wantMountsLen: 1,
		},
		{
			name:          "TLS and CA Secrets",
			service:       &Service{SecretName: "test-tls-secret"},
			ca:            &Ca{CaSecretName: "test-ca1"},
			wantMountsLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlsInstance := &TLS{Service: tt.service, Ca: tt.ca}
			mounts := tlsInstance.CreateVolumeMounts()
			if len(mounts) != tt.wantMountsLen {
				t.Errorf("CreateVolumeMounts() got = %v mounts, want %v mounts", len(mounts), tt.wantMountsLen)
			}
		})
	}
}

func TestCreateVolumes(t *testing.T) {
	tests := []struct {
		name       string
		service    *Service
		ca         *Ca
		wantVolLen int
	}{
		{
			name:       "No Secrets",
			service:    &Service{},
			ca:         &Ca{},
			wantVolLen: 0,
		},
		{
			name:       "Only TLS Secret",
			service:    &Service{SecretName: "test-tls-secret"},
			ca:         &Ca{},
			wantVolLen: 1,
		},
		{
			name:       "Only CA Secret",
			service:    &Service{},
			ca:         &Ca{CaSecretName: "test-ca1"},
			wantVolLen: 1,
		},
		{
			name:       "TLS and CA Secrets",
			service:    &Service{SecretName: "test-tls-secret"},
			ca:         &Ca{CaSecretName: "test-ca1"},
			wantVolLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlsInstance := &TLS{Service: tt.service, Ca: tt.ca}
			volumes := tlsInstance.CreateVolumes()
			if len(volumes) != tt.wantVolLen {
				t.Errorf("CreateVolumes() got = %v volumes, want %v volumes", len(volumes), tt.wantVolLen)
			}
		})
	}
}