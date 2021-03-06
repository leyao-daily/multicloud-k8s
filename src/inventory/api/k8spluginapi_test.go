/*
Copyright 2020  Tech Mahindra.
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

package api

import (
	"crypto/tls"
	con "github.com/onap/multicloud-k8s/src/inventory/constants"
	util "github.com/onap/multicloud-k8s/src/inventory/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestListInstances(t *testing.T) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	MK8S_URI := os.Getenv("onap-multicloud-k8s")
	MK8S_Port := os.Getenv("multicloud-k8s-port")

	instancelist := MK8S_URI + ":" + MK8S_Port + con.MK8S_EP
	req, err := http.NewRequest(http.MethodGet, instancelist, nil)

	util.SetRequestHeaders(req)

	if err != nil {
		t.Fatal(err)
	}

	client := http.DefaultClient
	res, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {

		t.Fail()
	}

	if p, err := ioutil.ReadAll(res.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s", p)
		} else if !strings.Contains(string(p), `cloud-region`) {
			t.Errorf("header response doen't match:\n%s", p)
		}
	}
}
