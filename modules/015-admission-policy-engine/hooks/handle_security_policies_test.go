/*
Copyright 2022 Flant JSC

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

package hooks

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deckhouse/deckhouse/testing/hooks"
)

var _ = Describe("Modules :: admission-policy-engine :: hooks :: handle security policies", func() {
	f := HookExecutionConfigInit(
		`{"admissionPolicyEngine": {"internal": {"bootstrapped": true} } }`,
		`{"admissionPolicyEngine":{}}`,
	)
	f.RegisterCRD("templates.gatekeeper.sh", "v1", "ConstraintTemplate", false)
	f.RegisterCRD("deckhouse.io", "v1alpha1", "SecurityPolicy", false)

	Context("Security Policy is set", func() {
		BeforeEach(func() {
			f.BindingContexts.Set(f.KubeStateSet(testSecurityPolicy))
			f.RunHook()
		})
		It("should have generated resources", func() {
			Expect(f).To(ExecuteSuccessfully())
			Expect(f.ValuesGet("admissionPolicyEngine.internal.securityPolicies").Array()).To(HaveLen(1))
		})
	})

})

var testSecurityPolicy = `
---
apiVersion: deckhouse.io/v1alpha1
kind: SecurityPolicy
metadata:
  name: foo
spec:
  enforcementAction: Deny
  match:
    namespaceSelector:
      labelSelector:
        matchLabels:
          operation-policy.deckhouse.io/enabled: "true"
  policies:
    allowHostNetwork: false
    allowPrivilegeEscalation: false
    allowPrivileged: false
    allowedCapabilities: []
    allowedAppArmor:
    - runtime/default
    allowedFlexVolumes:
    - driver: vmware
    allowedProcMount: Unmasked
    allowedUnsafeSysctls:
    - user/huyser
    allowedVolumes:
    - '*'
    forbiddenSysctls:
    - user/huyser
    fsGroup:
      rule: RunAsAny
    readOnlyRootFilesystem: true
    runAsGroup:
      ranges:
      - max: 500
        min: 300
      rule: RunAsAny
    runAsUser:
      ranges:
      - max: 500
        min: 300
      rule: MustRunAs
    supplementalGroups:
      ranges:
      - max: 1000
        min: 500
      rule: MustRunAs
    seLinux:
    - role: role
      user: user
    - level: level
      type: type
    allowHostIPC: true
    allowHostPID: false
    allowedHostPaths:
    - pathPrefix: /dev
      readOnly: true
    allowedHostPorts:
    - min: 10
      max: 100
    allowedUnsafeSysctls: ["*"]
    forbiddenSysctls:
    - user/example
    requiredDropCapabilities:
    - ALL
    seccompProfiles:
      allowedProfiles:
      - RuntimeDefault
      - Localhost
      allowedLocalhostFiles:
      - '*'

`