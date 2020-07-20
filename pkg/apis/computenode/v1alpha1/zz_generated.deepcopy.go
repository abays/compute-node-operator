// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStack) DeepCopyInto(out *ComputeNodeOpenStack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStack.
func (in *ComputeNodeOpenStack) DeepCopy() *ComputeNodeOpenStack {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeNodeOpenStack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStackList) DeepCopyInto(out *ComputeNodeOpenStackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputeNodeOpenStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStackList.
func (in *ComputeNodeOpenStackList) DeepCopy() *ComputeNodeOpenStackList {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeNodeOpenStackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStackSpec) DeepCopyInto(out *ComputeNodeOpenStackSpec) {
	*out = *in
	if in.InfraDaemonSets != nil {
		in, out := &in.InfraDaemonSets, &out.InfraDaemonSets
		*out = make([]InfraDaemonSet, len(*in))
		copy(*out, *in)
	}
	if in.NodesToDelete != nil {
		in, out := &in.NodesToDelete, &out.NodesToDelete
		*out = make([]NodeToDelete, len(*in))
		copy(*out, *in)
	}
	out.Drain = in.Drain
	out.Compute = in.Compute
	in.Network.DeepCopyInto(&out.Network)
	in.Performance.DeepCopyInto(&out.Performance)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStackSpec.
func (in *ComputeNodeOpenStackSpec) DeepCopy() *ComputeNodeOpenStackSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeOpenStackStatus) DeepCopyInto(out *ComputeNodeOpenStackStatus) {
	*out = *in
	if in.InfraDaemonSets != nil {
		in, out := &in.InfraDaemonSets, &out.InfraDaemonSets
		*out = make([]InfraDaemonSet, len(*in))
		copy(*out, *in)
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]Node, len(*in))
		copy(*out, *in)
	}
	if in.NodesToDelete != nil {
		in, out := &in.NodesToDelete, &out.NodesToDelete
		*out = make([]NodeToDelete, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeOpenStackStatus.
func (in *ComputeNodeOpenStackStatus) DeepCopy() *ComputeNodeOpenStackStatus {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeOpenStackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DrainParam) DeepCopyInto(out *DrainParam) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DrainParam.
func (in *DrainParam) DeepCopy() *DrainParam {
	if in == nil {
		return nil
	}
	out := new(DrainParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HugepagesConfig) DeepCopyInto(out *HugepagesConfig) {
	*out = *in
	if in.Pages != nil {
		in, out := &in.Pages, &out.Pages
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HugepagesConfig.
func (in *HugepagesConfig) DeepCopy() *HugepagesConfig {
	if in == nil {
		return nil
	}
	out := new(HugepagesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraDaemonSet) DeepCopyInto(out *InfraDaemonSet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraDaemonSet.
func (in *InfraDaemonSet) DeepCopy() *InfraDaemonSet {
	if in == nil {
		return nil
	}
	out := new(InfraDaemonSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronNetwork) DeepCopyInto(out *NeutronNetwork) {
	*out = *in
	if in.Sriov != nil {
		in, out := &in.Sriov, &out.Sriov
		*out = make([]SriovConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronNetwork.
func (in *NeutronNetwork) DeepCopy() *NeutronNetwork {
	if in == nil {
		return nil
	}
	out := new(NeutronNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeToDelete) DeepCopyInto(out *NodeToDelete) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeToDelete.
func (in *NodeToDelete) DeepCopy() *NodeToDelete {
	if in == nil {
		return nil
	}
	out := new(NodeToDelete)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NovaCompute) DeepCopyInto(out *NovaCompute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NovaCompute.
func (in *NovaCompute) DeepCopy() *NovaCompute {
	if in == nil {
		return nil
	}
	out := new(NovaCompute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceConfig) DeepCopyInto(out *PerformanceConfig) {
	*out = *in
	in.Hugepages.DeepCopyInto(&out.Hugepages)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceConfig.
func (in *PerformanceConfig) DeepCopy() *PerformanceConfig {
	if in == nil {
		return nil
	}
	out := new(PerformanceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovConfig) DeepCopyInto(out *SriovConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovConfig.
func (in *SriovConfig) DeepCopy() *SriovConfig {
	if in == nil {
		return nil
	}
	out := new(SriovConfig)
	in.DeepCopyInto(out)
	return out
}
