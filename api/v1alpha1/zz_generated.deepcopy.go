//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/fedstate/fedstate/pkg/driver/mgo"
	policyv1alpha1 "github.com/karmada-io/api/policy/v1alpha1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthSetting) DeepCopyInto(out *AuthSetting) {
	*out = *in
	if in.RootPasswd != nil {
		in, out := &in.RootPasswd, &out.RootPasswd
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthSetting.
func (in *AuthSetting) DeepCopy() *AuthSetting {
	if in == nil {
		return nil
	}
	out := new(AuthSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSetting) DeepCopyInto(out *ConfigSetting) {
	*out = *in
	if in.ConfigSet != nil {
		in, out := &in.ConfigSet, &out.ConfigSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigRef != nil {
		in, out := &in.ConfigRef, &out.ConfigRef
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSetting.
func (in *ConfigSetting) DeepCopy() *ConfigSetting {
	if in == nil {
		return nil
	}
	out := new(ConfigSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigVar) DeepCopyInto(out *ConfigVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigVar.
func (in *ConfigVar) DeepCopy() *ConfigVar {
	if in == nil {
		return nil
	}
	out := new(ConfigVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurrentInfo) DeepCopyInto(out *CurrentInfo) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceSetting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurrentInfo.
func (in *CurrentInfo) DeepCopy() *CurrentInfo {
	if in == nil {
		return nil
	}
	out := new(CurrentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBUserSpec) DeepCopyInto(out *DBUserSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBUserSpec.
func (in *DBUserSpec) DeepCopy() *DBUserSpec {
	if in == nil {
		return nil
	}
	out := new(DBUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportSetting) DeepCopyInto(out *ExportSetting) {
	*out = *in
	in.Resource.DeepCopyInto(&out.Resource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportSetting.
func (in *ExportSetting) DeepCopy() *ExportSetting {
	if in == nil {
		return nil
	}
	out := new(ExportSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePullSecretReference) DeepCopyInto(out *ImagePullSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePullSecretReference.
func (in *ImagePullSecretReference) DeepCopy() *ImagePullSecretReference {
	if in == nil {
		return nil
	}
	out := new(ImagePullSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePullSecretSpec) DeepCopyInto(out *ImagePullSecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePullSecretSpec.
func (in *ImagePullSecretSpec) DeepCopy() *ImagePullSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ImagePullSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSetting) DeepCopyInto(out *ImageSetting) {
	*out = *in
	out.ImagePullSecret = in.ImagePullSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSetting.
func (in *ImageSetting) DeepCopy() *ImageSetting {
	if in == nil {
		return nil
	}
	out := new(ImageSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberSetting) DeepCopyInto(out *MemberSetting) {
	*out = *in
	if in.MemberConfigRef != nil {
		in, out := &in.MemberConfigRef, &out.MemberConfigRef
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberSetting.
func (in *MemberSetting) DeepCopy() *MemberSetting {
	if in == nil {
		return nil
	}
	out := new(MemberSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsExporterSpec) DeepCopyInto(out *MetricsExporterSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceSetting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsExporterSpec.
func (in *MetricsExporterSpec) DeepCopy() *MetricsExporterSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoCondition) DeepCopyInto(out *MongoCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoCondition.
func (in *MongoCondition) DeepCopy() *MongoCondition {
	if in == nil {
		return nil
	}
	out := new(MongoCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDB) DeepCopyInto(out *MongoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDB.
func (in *MongoDB) DeepCopy() *MongoDB {
	if in == nil {
		return nil
	}
	out := new(MongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBList) DeepCopyInto(out *MongoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBList.
func (in *MongoDBList) DeepCopy() *MongoDBList {
	if in == nil {
		return nil
	}
	out := new(MongoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBSpec) DeepCopyInto(out *MongoDBSpec) {
	*out = *in
	if in.PodSpec != nil {
		in, out := &in.PodSpec, &out.PodSpec
		*out = new(PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricsExporterSpec != nil {
		in, out := &in.MetricsExporterSpec, &out.MetricsExporterSpec
		*out = new(MetricsExporterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceSetting)
		(*in).DeepCopyInto(*out)
	}
	out.DBUserSpec = in.DBUserSpec
	out.Persistence = in.Persistence
	out.ImagePullSecret = in.ImagePullSecret
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigVar, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBSpec.
func (in *MongoDBSpec) DeepCopy() *MongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBStatus) DeepCopyInto(out *MongoDBStatus) {
	*out = *in
	if in.ReplSet != nil {
		in, out := &in.ReplSet, &out.ReplSet
		*out = make([]mgo.MemberStatus, len(*in))
		copy(*out, *in)
	}
	in.CurrentInfo.DeepCopyInto(&out.CurrentInfo)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MongoCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBStatus.
func (in *MongoDBStatus) DeepCopy() *MongoDBStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiCloudMongoDB) DeepCopyInto(out *MultiCloudMongoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiCloudMongoDB.
func (in *MultiCloudMongoDB) DeepCopy() *MultiCloudMongoDB {
	if in == nil {
		return nil
	}
	out := new(MultiCloudMongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiCloudMongoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiCloudMongoDBList) DeepCopyInto(out *MultiCloudMongoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiCloudMongoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiCloudMongoDBList.
func (in *MultiCloudMongoDBList) DeepCopy() *MultiCloudMongoDBList {
	if in == nil {
		return nil
	}
	out := new(MultiCloudMongoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiCloudMongoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiCloudMongoDBSpec) DeepCopyInto(out *MultiCloudMongoDBSpec) {
	*out = *in
	in.Resource.DeepCopyInto(&out.Resource)
	if in.Replicaset != nil {
		in, out := &in.Replicaset, &out.Replicaset
		*out = new(int32)
		**out = **in
	}
	in.Auth.DeepCopyInto(&out.Auth)
	in.Member.DeepCopyInto(&out.Member)
	out.ImageSetting = in.ImageSetting
	out.Storage = in.Storage
	in.Export.DeepCopyInto(&out.Export)
	in.Config.DeepCopyInto(&out.Config)
	in.Scheduler.DeepCopyInto(&out.Scheduler)
	in.SpreadConstraints.DeepCopyInto(&out.SpreadConstraints)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiCloudMongoDBSpec.
func (in *MultiCloudMongoDBSpec) DeepCopy() *MultiCloudMongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(MultiCloudMongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiCloudMongoDBStatus) DeepCopyInto(out *MultiCloudMongoDBStatus) {
	*out = *in
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		*out = make([]*ServiceTopology, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ServiceTopology)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiCloudMongoDBStatus.
func (in *MultiCloudMongoDBStatus) DeepCopy() *MultiCloudMongoDBStatus {
	if in == nil {
		return nil
	}
	out := new(MultiCloudMongoDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceSpec) DeepCopyInto(out *PersistenceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceSpec.
func (in *PersistenceSpec) DeepCopy() *PersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(PersistenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpec) DeepCopyInto(out *PodSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpec.
func (in *PodSpec) DeepCopy() *PodSpec {
	if in == nil {
		return nil
	}
	out := new(PodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetting) DeepCopyInto(out *ResourceSetting) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetting.
func (in *ResourceSetting) DeepCopy() *ResourceSetting {
	if in == nil {
		return nil
	}
	out := new(ResourceSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerSetting) DeepCopyInto(out *SchedulerSetting) {
	*out = *in
	if in.SchedulerName != nil {
		in, out := &in.SchedulerName, &out.SchedulerName
		*out = new(string)
		**out = **in
	}
	if in.SchedulerMode != nil {
		in, out := &in.SchedulerMode, &out.SchedulerMode
		*out = new(string)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerSetting.
func (in *SchedulerSetting) DeepCopy() *SchedulerSetting {
	if in == nil {
		return nil
	}
	out := new(SchedulerSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerCondition) DeepCopyInto(out *ServerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerCondition.
func (in *ServerCondition) DeepCopy() *ServerCondition {
	if in == nil {
		return nil
	}
	out := new(ServerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTopology) DeepCopyInto(out *ServiceTopology) {
	*out = *in
	if in.ReplicasetStatus != nil {
		in, out := &in.ReplicasetStatus, &out.ReplicasetStatus
		*out = new(int)
		**out = **in
	}
	if in.ReplicasetSpec != nil {
		in, out := &in.ReplicasetSpec, &out.ReplicasetSpec
		*out = new(int)
		**out = **in
	}
	if in.ConnectAddrWithRole != nil {
		in, out := &in.ConnectAddrWithRole, &out.ConnectAddrWithRole
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTopology.
func (in *ServiceTopology) DeepCopy() *ServiceTopology {
	if in == nil {
		return nil
	}
	out := new(ServiceTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpreadConstraint) DeepCopyInto(out *SpreadConstraint) {
	*out = *in
	if in.NodeSelect != nil {
		in, out := &in.NodeSelect, &out.NodeSelect
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SpreadConstraints != nil {
		in, out := &in.SpreadConstraints, &out.SpreadConstraints
		*out = make([]policyv1alpha1.SpreadConstraint, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpreadConstraint.
func (in *SpreadConstraint) DeepCopy() *SpreadConstraint {
	if in == nil {
		return nil
	}
	out := new(SpreadConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSetting) DeepCopyInto(out *StorageSetting) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSetting.
func (in *StorageSetting) DeepCopy() *StorageSetting {
	if in == nil {
		return nil
	}
	out := new(StorageSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhook.
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}