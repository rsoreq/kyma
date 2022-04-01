//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileMount) DeepCopyInto(out *FileMount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileMount.
func (in *FileMount) DeepCopy() *FileMount {
	if in == nil {
		return nil
	}
	out := new(FileMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogPipeline) DeepCopyInto(out *LogPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogPipeline.
func (in *LogPipeline) DeepCopy() *LogPipeline {
	if in == nil {
		return nil
	}
	out := new(LogPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogPipelineCondition) DeepCopyInto(out *LogPipelineCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogPipelineCondition.
func (in *LogPipelineCondition) DeepCopy() *LogPipelineCondition {
	if in == nil {
		return nil
	}
	out := new(LogPipelineCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogPipelineList) DeepCopyInto(out *LogPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogPipelineList.
func (in *LogPipelineList) DeepCopy() *LogPipelineList {
	if in == nil {
		return nil
	}
	out := new(LogPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogPipelineSpec) DeepCopyInto(out *LogPipelineSpec) {
	*out = *in
	if in.Parsers != nil {
		in, out := &in.Parsers, &out.Parsers
		*out = make([]Parser, len(*in))
		copy(*out, *in)
	}
	if in.MultiLineParsers != nil {
		in, out := &in.MultiLineParsers, &out.MultiLineParsers
		*out = make([]MultiLineParser, len(*in))
		copy(*out, *in)
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]Filter, len(*in))
		copy(*out, *in)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]Output, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileMount, len(*in))
		copy(*out, *in)
	}
	if in.SecretRefs != nil {
		in, out := &in.SecretRefs, &out.SecretRefs
		*out = make([]SecretReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogPipelineSpec.
func (in *LogPipelineSpec) DeepCopy() *LogPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(LogPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogPipelineStatus) DeepCopyInto(out *LogPipelineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]LogPipelineCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogPipelineStatus.
func (in *LogPipelineStatus) DeepCopy() *LogPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(LogPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiLineParser) DeepCopyInto(out *MultiLineParser) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiLineParser.
func (in *MultiLineParser) DeepCopy() *MultiLineParser {
	if in == nil {
		return nil
	}
	out := new(MultiLineParser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Output) DeepCopyInto(out *Output) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Output.
func (in *Output) DeepCopy() *Output {
	if in == nil {
		return nil
	}
	out := new(Output)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parser) DeepCopyInto(out *Parser) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parser.
func (in *Parser) DeepCopy() *Parser {
	if in == nil {
		return nil
	}
	out := new(Parser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}
