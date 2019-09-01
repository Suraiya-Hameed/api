// +build !ignore_autogenerated

// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Code generated by deepcopy-gen. DO NOT EDIT.

package projectcalico

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditEventsSelection) DeepCopyInto(out *AuditEventsSelection) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]AuditResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditEventsSelection.
func (in *AuditEventsSelection) DeepCopy() *AuditEventsSelection {
	if in == nil {
		return nil
	}
	out := new(AuditEventsSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditResource) DeepCopyInto(out *AuditResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditResource.
func (in *AuditResource) DeepCopy() *AuditResource {
	if in == nil {
		return nil
	}
	out := new(AuditResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalReportType) DeepCopyInto(out *GlobalReportType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalReportType.
func (in *GlobalReportType) DeepCopy() *GlobalReportType {
	if in == nil {
		return nil
	}
	out := new(GlobalReportType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalReportType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalReportTypeList) DeepCopyInto(out *GlobalReportTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalReportType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalReportTypeList.
func (in *GlobalReportTypeList) DeepCopy() *GlobalReportTypeList {
	if in == nil {
		return nil
	}
	out := new(GlobalReportTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalReportTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseKey) DeepCopyInto(out *LicenseKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseKey.
func (in *LicenseKey) DeepCopy() *LicenseKey {
	if in == nil {
		return nil
	}
	out := new(LicenseKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseKeyList) DeepCopyInto(out *LicenseKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseKeyList.
func (in *LicenseKeyList) DeepCopy() *LicenseKeyList {
	if in == nil {
		return nil
	}
	out := new(LicenseKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseKeySpec) DeepCopyInto(out *LicenseKeySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseKeySpec.
func (in *LicenseKeySpec) DeepCopy() *LicenseKeySpec {
	if in == nil {
		return nil
	}
	out := new(LicenseKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportTemplate) DeepCopyInto(out *ReportTemplate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportTemplate.
func (in *ReportTemplate) DeepCopy() *ReportTemplate {
	if in == nil {
		return nil
	}
	out := new(ReportTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportTypeSpec) DeepCopyInto(out *ReportTypeSpec) {
	*out = *in
	out.UISummaryTemplate = in.UISummaryTemplate
	if in.DownloadTemplates != nil {
		in, out := &in.DownloadTemplates, &out.DownloadTemplates
		*out = make([]ReportTemplate, len(*in))
		copy(*out, *in)
	}
	if in.AuditEventsSelection != nil {
		in, out := &in.AuditEventsSelection, &out.AuditEventsSelection
		*out = new(AuditEventsSelection)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportTypeSpec.
func (in *ReportTypeSpec) DeepCopy() *ReportTypeSpec {
	if in == nil {
		return nil
	}
	out := new(ReportTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceID) DeepCopyInto(out *ResourceID) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceID.
func (in *ResourceID) DeepCopy() *ResourceID {
	if in == nil {
		return nil
	}
	out := new(ResourceID)
	in.DeepCopyInto(out)
	return out
}
