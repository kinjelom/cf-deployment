package legacy_test

import (
	"testing"

	"tests/helpers"
)

const testDirectory = "operations/legacy"

var legacyTests = map[string]helpers.OpsFileTestParams{
	"keep-haproxy-ssl-pem.yml": {
		Ops:  []string{"../use-haproxy.yml", "keep-haproxy-ssl-pem.yml"},
		Vars: []string{"haproxy_ssl_pem=i-am-a-pem", "haproxy_private_ip=1.2.3.4"},
	},
	"keep-original-blobstore-directory-keys.yml": {
		VarsFiles: []string{"example-vars-files/vars-keep-original-blobstore-directory-keys.yml"},
	},
	"keep-original-internal-usernames.yml": {
		VarsFiles: []string{"example-vars-files/vars-keep-original-internal-usernames.yml"},
	},
	"keep-original-postgres-configuration.yml": {
		Ops:       []string{"../use-postgres.yml", "keep-original-postgres-configuration.yml"},
		VarsFiles: []string{"example-vars-files/vars-keep-original-postgres-configuration.yml"},
	},
	"keep-original-routing-postgres-configuration.yml": {
		Ops:  []string{"../use-postgres.yml", "keep-original-routing-postgres-configuration.yml"},
		Vars: []string{"routing_api_database_name=routing-api-db", "routing_api_database_username=routing-api-user"},
	},
	"keep-static-ips.yml": {
		VarsFiles: []string{"example-vars-files/vars-keep-static-ips.yml"},
	},
	"old-droplet-mitigation.yml": {},
}

func TestLegacy(t *testing.T) {
	cfDeploymentHome, err := helpers.SetPath()
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	fileNames, err := helpers.FindFiles(cfDeploymentHome, testDirectory)
	if err != nil {
		t.Fatalf("setup: %v", err)
	}
	for _, fileName := range fileNames {
		t.Run("ensure "+fileName+" has test", func(t *testing.T) {
			// TODO: only skip with some sort of flag
			// t.Skip()
			if _, hasTest := legacyTests[fileName]; !hasTest {
				t.Error("Missing test for:", fileName)
			}
		})
	}
	if t.Failed() {
		t.FailNow()
	}

	for name, params := range legacyTests {
		name, params := name, params
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := helpers.CheckInterpolate(cfDeploymentHome, testDirectory, name, params); err != nil {
				t.Error("interpolate failed:", err)
			}
		})
	}
}
