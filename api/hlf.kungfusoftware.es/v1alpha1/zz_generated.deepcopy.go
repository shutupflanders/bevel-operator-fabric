// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/operator-framework/operator-lib/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CA) DeepCopyInto(out *CA) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CA.
func (in *CA) DeepCopy() *CA {
	if in == nil {
		return nil
	}
	out := new(CA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Catls) DeepCopyInto(out *Catls) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Catls.
func (in *Catls) DeepCopy() *Catls {
	if in == nil {
		return nil
	}
	out := new(Catls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	out.Catls = in.Catls
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cors) DeepCopyInto(out *Cors) {
	*out = *in
	if in.Origins != nil {
		in, out := &in.Origins, &out.Origins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cors.
func (in *Cors) DeepCopy() *Cors {
	if in == nil {
		return nil
	}
	out := new(Cors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Csr) DeepCopyInto(out *Csr) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Csr.
func (in *Csr) DeepCopy() *Csr {
	if in == nil {
		return nil
	}
	out := new(Csr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Enrollment) DeepCopyInto(out *Enrollment) {
	*out = *in
	out.Component = in.Component
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Enrollment.
func (in *Enrollment) DeepCopy() *Enrollment {
	if in == nil {
		return nil
	}
	out := new(Enrollment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCA) DeepCopyInto(out *FabricCA) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCA.
func (in *FabricCA) DeepCopy() *FabricCA {
	if in == nil {
		return nil
	}
	out := new(FabricCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricCA) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCABCCSP) DeepCopyInto(out *FabricCABCCSP) {
	*out = *in
	out.SW = in.SW
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCABCCSP.
func (in *FabricCABCCSP) DeepCopy() *FabricCABCCSP {
	if in == nil {
		return nil
	}
	out := new(FabricCABCCSP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCABCCSPSW) DeepCopyInto(out *FabricCABCCSPSW) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCABCCSPSW.
func (in *FabricCABCCSPSW) DeepCopy() *FabricCABCCSPSW {
	if in == nil {
		return nil
	}
	out := new(FabricCABCCSPSW)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCACFG) DeepCopyInto(out *FabricCACFG) {
	*out = *in
	out.Identities = in.Identities
	out.Affiliations = in.Affiliations
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCACFG.
func (in *FabricCACFG) DeepCopy() *FabricCACFG {
	if in == nil {
		return nil
	}
	out := new(FabricCACFG)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCACFGAffilitions) DeepCopyInto(out *FabricCACFGAffilitions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCACFGAffilitions.
func (in *FabricCACFGAffilitions) DeepCopy() *FabricCACFGAffilitions {
	if in == nil {
		return nil
	}
	out := new(FabricCACFGAffilitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCACFGIdentities) DeepCopyInto(out *FabricCACFGIdentities) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCACFGIdentities.
func (in *FabricCACFGIdentities) DeepCopy() *FabricCACFGIdentities {
	if in == nil {
		return nil
	}
	out := new(FabricCACFGIdentities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCACRL) DeepCopyInto(out *FabricCACRL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCACRL.
func (in *FabricCACRL) DeepCopy() *FabricCACRL {
	if in == nil {
		return nil
	}
	out := new(FabricCACRL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCACSR) DeepCopyInto(out *FabricCACSR) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]FabricCANames, len(*in))
		copy(*out, *in)
	}
	out.CA = in.CA
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCACSR.
func (in *FabricCACSR) DeepCopy() *FabricCACSR {
	if in == nil {
		return nil
	}
	out := new(FabricCACSR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCACSRCA) DeepCopyInto(out *FabricCACSRCA) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCACSRCA.
func (in *FabricCACSRCA) DeepCopy() *FabricCACSRCA {
	if in == nil {
		return nil
	}
	out := new(FabricCACSRCA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCADatabase) DeepCopyInto(out *FabricCADatabase) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCADatabase.
func (in *FabricCADatabase) DeepCopy() *FabricCADatabase {
	if in == nil {
		return nil
	}
	out := new(FabricCADatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAIdentity) DeepCopyInto(out *FabricCAIdentity) {
	*out = *in
	out.Attrs = in.Attrs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAIdentity.
func (in *FabricCAIdentity) DeepCopy() *FabricCAIdentity {
	if in == nil {
		return nil
	}
	out := new(FabricCAIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAIdentityAttrs) DeepCopyInto(out *FabricCAIdentityAttrs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAIdentityAttrs.
func (in *FabricCAIdentityAttrs) DeepCopy() *FabricCAIdentityAttrs {
	if in == nil {
		return nil
	}
	out := new(FabricCAIdentityAttrs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAIntermediate) DeepCopyInto(out *FabricCAIntermediate) {
	*out = *in
	out.ParentServer = in.ParentServer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAIntermediate.
func (in *FabricCAIntermediate) DeepCopy() *FabricCAIntermediate {
	if in == nil {
		return nil
	}
	out := new(FabricCAIntermediate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAIntermediateEnrollment) DeepCopyInto(out *FabricCAIntermediateEnrollment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAIntermediateEnrollment.
func (in *FabricCAIntermediateEnrollment) DeepCopy() *FabricCAIntermediateEnrollment {
	if in == nil {
		return nil
	}
	out := new(FabricCAIntermediateEnrollment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAIntermediateParentServer) DeepCopyInto(out *FabricCAIntermediateParentServer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAIntermediateParentServer.
func (in *FabricCAIntermediateParentServer) DeepCopy() *FabricCAIntermediateParentServer {
	if in == nil {
		return nil
	}
	out := new(FabricCAIntermediateParentServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAIntermediateTLS) DeepCopyInto(out *FabricCAIntermediateTLS) {
	*out = *in
	if in.CertFiles != nil {
		in, out := &in.CertFiles, &out.CertFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Client = in.Client
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAIntermediateTLS.
func (in *FabricCAIntermediateTLS) DeepCopy() *FabricCAIntermediateTLS {
	if in == nil {
		return nil
	}
	out := new(FabricCAIntermediateTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAIntermediateTLSClient) DeepCopyInto(out *FabricCAIntermediateTLSClient) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAIntermediateTLSClient.
func (in *FabricCAIntermediateTLSClient) DeepCopy() *FabricCAIntermediateTLSClient {
	if in == nil {
		return nil
	}
	out := new(FabricCAIntermediateTLSClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAItemConf) DeepCopyInto(out *FabricCAItemConf) {
	*out = *in
	out.CFG = in.CFG
	out.Subject = in.Subject
	in.CSR.DeepCopyInto(&out.CSR)
	out.CRL = in.CRL
	in.Registry.DeepCopyInto(&out.Registry)
	out.Intermediate = in.Intermediate
	out.BCCSP = in.BCCSP
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAItemConf.
func (in *FabricCAItemConf) DeepCopy() *FabricCAItemConf {
	if in == nil {
		return nil
	}
	out := new(FabricCAItemConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAList) DeepCopyInto(out *FabricCAList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FabricCA, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAList.
func (in *FabricCAList) DeepCopy() *FabricCAList {
	if in == nil {
		return nil
	}
	out := new(FabricCAList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricCAList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAMetrics) DeepCopyInto(out *FabricCAMetrics) {
	*out = *in
	out.Statsd = in.Statsd
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAMetrics.
func (in *FabricCAMetrics) DeepCopy() *FabricCAMetrics {
	if in == nil {
		return nil
	}
	out := new(FabricCAMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAMetricsStatsd) DeepCopyInto(out *FabricCAMetricsStatsd) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAMetricsStatsd.
func (in *FabricCAMetricsStatsd) DeepCopy() *FabricCAMetricsStatsd {
	if in == nil {
		return nil
	}
	out := new(FabricCAMetricsStatsd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCANames) DeepCopyInto(out *FabricCANames) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCANames.
func (in *FabricCANames) DeepCopy() *FabricCANames {
	if in == nil {
		return nil
	}
	out := new(FabricCANames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCARegistry) DeepCopyInto(out *FabricCARegistry) {
	*out = *in
	if in.Identities != nil {
		in, out := &in.Identities, &out.Identities
		*out = make([]FabricCAIdentity, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCARegistry.
func (in *FabricCARegistry) DeepCopy() *FabricCARegistry {
	if in == nil {
		return nil
	}
	out := new(FabricCARegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCASpec) DeepCopyInto(out *FabricCASpec) {
	*out = *in
	out.Database = in.Database
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Service = in.Service
	out.TLS = in.TLS
	in.CA.DeepCopyInto(&out.CA)
	in.TLSCA.DeepCopyInto(&out.TLSCA)
	in.Cors.DeepCopyInto(&out.Cors)
	out.Resources = in.Resources
	out.Storage = in.Storage
	out.Metrics = in.Metrics
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCASpec.
func (in *FabricCASpec) DeepCopy() *FabricCASpec {
	if in == nil {
		return nil
	}
	out := new(FabricCASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCASpecService) DeepCopyInto(out *FabricCASpecService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCASpecService.
func (in *FabricCASpecService) DeepCopy() *FabricCASpecService {
	if in == nil {
		return nil
	}
	out := new(FabricCASpecService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCAStatus) DeepCopyInto(out *FabricCAStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCAStatus.
func (in *FabricCAStatus) DeepCopy() *FabricCAStatus {
	if in == nil {
		return nil
	}
	out := new(FabricCAStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCASubject) DeepCopyInto(out *FabricCASubject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCASubject.
func (in *FabricCASubject) DeepCopy() *FabricCASubject {
	if in == nil {
		return nil
	}
	out := new(FabricCASubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricCATLSConf) DeepCopyInto(out *FabricCATLSConf) {
	*out = *in
	out.Subject = in.Subject
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricCATLSConf.
func (in *FabricCATLSConf) DeepCopy() *FabricCATLSConf {
	if in == nil {
		return nil
	}
	out := new(FabricCATLSConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrdererNode) DeepCopyInto(out *FabricOrdererNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrdererNode.
func (in *FabricOrdererNode) DeepCopy() *FabricOrdererNode {
	if in == nil {
		return nil
	}
	out := new(FabricOrdererNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricOrdererNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrdererNodeList) DeepCopyInto(out *FabricOrdererNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FabricOrdererNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrdererNodeList.
func (in *FabricOrdererNodeList) DeepCopy() *FabricOrdererNodeList {
	if in == nil {
		return nil
	}
	out := new(FabricOrdererNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricOrdererNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrdererNodeSpec) DeepCopyInto(out *FabricOrdererNodeSpec) {
	*out = *in
	out.Storage = in.Storage
	out.Service = in.Service
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrdererNodeSpec.
func (in *FabricOrdererNodeSpec) DeepCopy() *FabricOrdererNodeSpec {
	if in == nil {
		return nil
	}
	out := new(FabricOrdererNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrdererNodeStatus) DeepCopyInto(out *FabricOrdererNodeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrdererNodeStatus.
func (in *FabricOrdererNodeStatus) DeepCopy() *FabricOrdererNodeStatus {
	if in == nil {
		return nil
	}
	out := new(FabricOrdererNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrderingService) DeepCopyInto(out *FabricOrderingService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrderingService.
func (in *FabricOrderingService) DeepCopy() *FabricOrderingService {
	if in == nil {
		return nil
	}
	out := new(FabricOrderingService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricOrderingService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrderingServiceList) DeepCopyInto(out *FabricOrderingServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FabricOrderingService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrderingServiceList.
func (in *FabricOrderingServiceList) DeepCopy() *FabricOrderingServiceList {
	if in == nil {
		return nil
	}
	out := new(FabricOrderingServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricOrderingServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrderingServiceSpec) DeepCopyInto(out *FabricOrderingServiceSpec) {
	*out = *in
	in.Enrollment.DeepCopyInto(&out.Enrollment)
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]OrdererNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Service = in.Service
	out.Storage = in.Storage
	out.SystemChannel = in.SystemChannel
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrderingServiceSpec.
func (in *FabricOrderingServiceSpec) DeepCopy() *FabricOrderingServiceSpec {
	if in == nil {
		return nil
	}
	out := new(FabricOrderingServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricOrderingServiceStatus) DeepCopyInto(out *FabricOrderingServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricOrderingServiceStatus.
func (in *FabricOrderingServiceStatus) DeepCopy() *FabricOrderingServiceStatus {
	if in == nil {
		return nil
	}
	out := new(FabricOrderingServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricPeer) DeepCopyInto(out *FabricPeer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricPeer.
func (in *FabricPeer) DeepCopy() *FabricPeer {
	if in == nil {
		return nil
	}
	out := new(FabricPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricPeer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricPeerList) DeepCopyInto(out *FabricPeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FabricPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricPeerList.
func (in *FabricPeerList) DeepCopy() *FabricPeerList {
	if in == nil {
		return nil
	}
	out := new(FabricPeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricPeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricPeerSpec) DeepCopyInto(out *FabricPeerSpec) {
	*out = *in
	in.Secret.DeepCopyInto(&out.Secret)
	out.Service = in.Service
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OperationHosts != nil {
		in, out := &in.OperationHosts, &out.OperationHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OperationIPs != nil {
		in, out := &in.OperationIPs, &out.OperationIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricPeerSpec.
func (in *FabricPeerSpec) DeepCopy() *FabricPeerSpec {
	if in == nil {
		return nil
	}
	out := new(FabricPeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricPeerStatus) DeepCopyInto(out *FabricPeerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricPeerStatus.
func (in *FabricPeerStatus) DeepCopy() *FabricPeerStatus {
	if in == nil {
		return nil
	}
	out := new(FabricPeerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererEnrollment) DeepCopyInto(out *OrdererEnrollment) {
	*out = *in
	out.Component = in.Component
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererEnrollment.
func (in *OrdererEnrollment) DeepCopy() *OrdererEnrollment {
	if in == nil {
		return nil
	}
	out := new(OrdererEnrollment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererNode) DeepCopyInto(out *OrdererNode) {
	*out = *in
	in.Enrollment.DeepCopyInto(&out.Enrollment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererNode.
func (in *OrdererNode) DeepCopy() *OrdererNode {
	if in == nil {
		return nil
	}
	out := new(OrdererNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererNodeEnrollment) DeepCopyInto(out *OrdererNodeEnrollment) {
	*out = *in
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererNodeEnrollment.
func (in *OrdererNodeEnrollment) DeepCopy() *OrdererNodeEnrollment {
	if in == nil {
		return nil
	}
	out := new(OrdererNodeEnrollment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererNodeEnrollmentTLS) DeepCopyInto(out *OrdererNodeEnrollmentTLS) {
	*out = *in
	in.Csr.DeepCopyInto(&out.Csr)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererNodeEnrollmentTLS.
func (in *OrdererNodeEnrollmentTLS) DeepCopy() *OrdererNodeEnrollmentTLS {
	if in == nil {
		return nil
	}
	out := new(OrdererNodeEnrollmentTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererNodeService) DeepCopyInto(out *OrdererNodeService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererNodeService.
func (in *OrdererNodeService) DeepCopy() *OrdererNodeService {
	if in == nil {
		return nil
	}
	out := new(OrdererNodeService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererService) DeepCopyInto(out *OrdererService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererService.
func (in *OrdererService) DeepCopy() *OrdererService {
	if in == nil {
		return nil
	}
	out := new(OrdererService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererSystemChannel) DeepCopyInto(out *OrdererSystemChannel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererSystemChannel.
func (in *OrdererSystemChannel) DeepCopy() *OrdererSystemChannel {
	if in == nil {
		return nil
	}
	out := new(OrdererSystemChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerService) DeepCopyInto(out *PeerService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerService.
func (in *PeerService) DeepCopy() *PeerService {
	if in == nil {
		return nil
	}
	out := new(PeerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requests) DeepCopyInto(out *Requests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requests.
func (in *Requests) DeepCopy() *Requests {
	if in == nil {
		return nil
	}
	out := new(Requests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestsLimit) DeepCopyInto(out *RequestsLimit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestsLimit.
func (in *RequestsLimit) DeepCopy() *RequestsLimit {
	if in == nil {
		return nil
	}
	out := new(RequestsLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	in.Enrollment.DeepCopyInto(&out.Enrollment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	out.Catls = in.Catls
	in.Csr.DeepCopyInto(&out.Csr)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}
