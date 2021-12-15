//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	v1beta2 "github.com/cybozu-go/moco/api/v1beta2"
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BackupStatus)(nil), (*v1beta2.BackupStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__BackupStatus_To_v1beta2_BackupStatus(a.(*BackupStatus), b.(*v1beta2.BackupStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.BackupStatus)(nil), (*BackupStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_BackupStatus_To__BackupStatus(a.(*v1beta2.BackupStatus), b.(*BackupStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BucketConfig)(nil), (*v1beta2.BucketConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__BucketConfig_To_v1beta2_BucketConfig(a.(*BucketConfig), b.(*v1beta2.BucketConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.BucketConfig)(nil), (*BucketConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_BucketConfig_To__BucketConfig(a.(*v1beta2.BucketConfig), b.(*BucketConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*JobConfig)(nil), (*v1beta2.JobConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__JobConfig_To_v1beta2_JobConfig(a.(*JobConfig), b.(*v1beta2.JobConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.JobConfig)(nil), (*JobConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_JobConfig_To__JobConfig(a.(*v1beta2.JobConfig), b.(*JobConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MySQLClusterCondition)(nil), (*v1beta2.MySQLClusterCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__MySQLClusterCondition_To_v1beta2_MySQLClusterCondition(a.(*MySQLClusterCondition), b.(*v1beta2.MySQLClusterCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.MySQLClusterCondition)(nil), (*MySQLClusterCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MySQLClusterCondition_To__MySQLClusterCondition(a.(*v1beta2.MySQLClusterCondition), b.(*MySQLClusterCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MySQLClusterList)(nil), (*v1beta2.MySQLClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__MySQLClusterList_To_v1beta2_MySQLClusterList(a.(*MySQLClusterList), b.(*v1beta2.MySQLClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.MySQLClusterList)(nil), (*MySQLClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MySQLClusterList_To__MySQLClusterList(a.(*v1beta2.MySQLClusterList), b.(*MySQLClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MySQLClusterStatus)(nil), (*v1beta2.MySQLClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__MySQLClusterStatus_To_v1beta2_MySQLClusterStatus(a.(*MySQLClusterStatus), b.(*v1beta2.MySQLClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.MySQLClusterStatus)(nil), (*MySQLClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MySQLClusterStatus_To__MySQLClusterStatus(a.(*v1beta2.MySQLClusterStatus), b.(*MySQLClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ObjectMeta)(nil), (*v1beta2.ObjectMeta)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__ObjectMeta_To_v1beta2_ObjectMeta(a.(*ObjectMeta), b.(*v1beta2.ObjectMeta), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.ObjectMeta)(nil), (*ObjectMeta)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ObjectMeta_To__ObjectMeta(a.(*v1beta2.ObjectMeta), b.(*ObjectMeta), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PersistentVolumeClaim)(nil), (*v1beta2.PersistentVolumeClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__PersistentVolumeClaim_To_v1beta2_PersistentVolumeClaim(a.(*PersistentVolumeClaim), b.(*v1beta2.PersistentVolumeClaim), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.PersistentVolumeClaim)(nil), (*PersistentVolumeClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_PersistentVolumeClaim_To__PersistentVolumeClaim(a.(*v1beta2.PersistentVolumeClaim), b.(*PersistentVolumeClaim), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodTemplateSpec)(nil), (*v1beta2.PodTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__PodTemplateSpec_To_v1beta2_PodTemplateSpec(a.(*PodTemplateSpec), b.(*v1beta2.PodTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.PodTemplateSpec)(nil), (*PodTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_PodTemplateSpec_To__PodTemplateSpec(a.(*v1beta2.PodTemplateSpec), b.(*PodTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ReconcileInfo)(nil), (*v1beta2.ReconcileInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__ReconcileInfo_To_v1beta2_ReconcileInfo(a.(*ReconcileInfo), b.(*v1beta2.ReconcileInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.ReconcileInfo)(nil), (*ReconcileInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ReconcileInfo_To__ReconcileInfo(a.(*v1beta2.ReconcileInfo), b.(*ReconcileInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RestoreSpec)(nil), (*v1beta2.RestoreSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__RestoreSpec_To_v1beta2_RestoreSpec(a.(*RestoreSpec), b.(*v1beta2.RestoreSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.RestoreSpec)(nil), (*RestoreSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_RestoreSpec_To__RestoreSpec(a.(*v1beta2.RestoreSpec), b.(*RestoreSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceTemplate)(nil), (*v1beta2.ServiceTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__ServiceTemplate_To_v1beta2_ServiceTemplate(a.(*ServiceTemplate), b.(*v1beta2.ServiceTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta2.ServiceTemplate)(nil), (*ServiceTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_ServiceTemplate_To__ServiceTemplate(a.(*v1beta2.ServiceTemplate), b.(*ServiceTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*MySQLClusterSpec)(nil), (*v1beta2.MySQLClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__MySQLClusterSpec_To_v1beta2_MySQLClusterSpec(a.(*MySQLClusterSpec), b.(*v1beta2.MySQLClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*MySQLCluster)(nil), (*v1beta2.MySQLCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert__MySQLCluster_To_v1beta2_MySQLCluster(a.(*MySQLCluster), b.(*v1beta2.MySQLCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta2.MySQLClusterSpec)(nil), (*MySQLClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MySQLClusterSpec_To__MySQLClusterSpec(a.(*v1beta2.MySQLClusterSpec), b.(*MySQLClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta2.MySQLCluster)(nil), (*MySQLCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta2_MySQLCluster_To__MySQLCluster(a.(*v1beta2.MySQLCluster), b.(*MySQLCluster), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert__BackupStatus_To_v1beta2_BackupStatus(in *BackupStatus, out *v1beta2.BackupStatus, s conversion.Scope) error {
	out.Time = in.Time
	out.Elapsed = in.Elapsed
	out.SourceIndex = in.SourceIndex
	out.SourceUUID = in.SourceUUID
	out.BinlogFilename = in.BinlogFilename
	out.GTIDSet = in.GTIDSet
	out.DumpSize = in.DumpSize
	out.BinlogSize = in.BinlogSize
	out.WorkDirUsage = in.WorkDirUsage
	out.Warnings = *(*[]string)(unsafe.Pointer(&in.Warnings))
	return nil
}

// Convert__BackupStatus_To_v1beta2_BackupStatus is an autogenerated conversion function.
func Convert__BackupStatus_To_v1beta2_BackupStatus(in *BackupStatus, out *v1beta2.BackupStatus, s conversion.Scope) error {
	return autoConvert__BackupStatus_To_v1beta2_BackupStatus(in, out, s)
}

func autoConvert_v1beta2_BackupStatus_To__BackupStatus(in *v1beta2.BackupStatus, out *BackupStatus, s conversion.Scope) error {
	out.Time = in.Time
	out.Elapsed = in.Elapsed
	out.SourceIndex = in.SourceIndex
	out.SourceUUID = in.SourceUUID
	out.BinlogFilename = in.BinlogFilename
	out.GTIDSet = in.GTIDSet
	out.DumpSize = in.DumpSize
	out.BinlogSize = in.BinlogSize
	out.WorkDirUsage = in.WorkDirUsage
	out.Warnings = *(*[]string)(unsafe.Pointer(&in.Warnings))
	return nil
}

// Convert_v1beta2_BackupStatus_To__BackupStatus is an autogenerated conversion function.
func Convert_v1beta2_BackupStatus_To__BackupStatus(in *v1beta2.BackupStatus, out *BackupStatus, s conversion.Scope) error {
	return autoConvert_v1beta2_BackupStatus_To__BackupStatus(in, out, s)
}

func autoConvert__BucketConfig_To_v1beta2_BucketConfig(in *BucketConfig, out *v1beta2.BucketConfig, s conversion.Scope) error {
	out.BucketName = in.BucketName
	out.Region = in.Region
	out.EndpointURL = in.EndpointURL
	out.UsePathStyle = in.UsePathStyle
	return nil
}

// Convert__BucketConfig_To_v1beta2_BucketConfig is an autogenerated conversion function.
func Convert__BucketConfig_To_v1beta2_BucketConfig(in *BucketConfig, out *v1beta2.BucketConfig, s conversion.Scope) error {
	return autoConvert__BucketConfig_To_v1beta2_BucketConfig(in, out, s)
}

func autoConvert_v1beta2_BucketConfig_To__BucketConfig(in *v1beta2.BucketConfig, out *BucketConfig, s conversion.Scope) error {
	out.BucketName = in.BucketName
	out.Region = in.Region
	out.EndpointURL = in.EndpointURL
	out.UsePathStyle = in.UsePathStyle
	return nil
}

// Convert_v1beta2_BucketConfig_To__BucketConfig is an autogenerated conversion function.
func Convert_v1beta2_BucketConfig_To__BucketConfig(in *v1beta2.BucketConfig, out *BucketConfig, s conversion.Scope) error {
	return autoConvert_v1beta2_BucketConfig_To__BucketConfig(in, out, s)
}

func autoConvert__JobConfig_To_v1beta2_JobConfig(in *JobConfig, out *v1beta2.JobConfig, s conversion.Scope) error {
	out.ServiceAccountName = in.ServiceAccountName
	if err := Convert__BucketConfig_To_v1beta2_BucketConfig(&in.BucketConfig, &out.BucketConfig, s); err != nil {
		return err
	}
	out.WorkVolume = in.WorkVolume
	out.Threads = in.Threads
	out.Memory = (*resource.Quantity)(unsafe.Pointer(in.Memory))
	out.MaxMemory = (*resource.Quantity)(unsafe.Pointer(in.MaxMemory))
	out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	return nil
}

// Convert__JobConfig_To_v1beta2_JobConfig is an autogenerated conversion function.
func Convert__JobConfig_To_v1beta2_JobConfig(in *JobConfig, out *v1beta2.JobConfig, s conversion.Scope) error {
	return autoConvert__JobConfig_To_v1beta2_JobConfig(in, out, s)
}

func autoConvert_v1beta2_JobConfig_To__JobConfig(in *v1beta2.JobConfig, out *JobConfig, s conversion.Scope) error {
	out.ServiceAccountName = in.ServiceAccountName
	if err := Convert_v1beta2_BucketConfig_To__BucketConfig(&in.BucketConfig, &out.BucketConfig, s); err != nil {
		return err
	}
	out.WorkVolume = in.WorkVolume
	out.Threads = in.Threads
	out.Memory = (*resource.Quantity)(unsafe.Pointer(in.Memory))
	out.MaxMemory = (*resource.Quantity)(unsafe.Pointer(in.MaxMemory))
	out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	return nil
}

// Convert_v1beta2_JobConfig_To__JobConfig is an autogenerated conversion function.
func Convert_v1beta2_JobConfig_To__JobConfig(in *v1beta2.JobConfig, out *JobConfig, s conversion.Scope) error {
	return autoConvert_v1beta2_JobConfig_To__JobConfig(in, out, s)
}

func autoConvert__MySQLCluster_To_v1beta2_MySQLCluster(in *MySQLCluster, out *v1beta2.MySQLCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert__MySQLClusterSpec_To_v1beta2_MySQLClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert__MySQLClusterStatus_To_v1beta2_MySQLClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta2_MySQLCluster_To__MySQLCluster(in *v1beta2.MySQLCluster, out *MySQLCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_MySQLClusterSpec_To__MySQLClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_MySQLClusterStatus_To__MySQLClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func autoConvert__MySQLClusterCondition_To_v1beta2_MySQLClusterCondition(in *MySQLClusterCondition, out *v1beta2.MySQLClusterCondition, s conversion.Scope) error {
	out.Type = v1beta2.MySQLClusterConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastTransitionTime = in.LastTransitionTime
	return nil
}

// Convert__MySQLClusterCondition_To_v1beta2_MySQLClusterCondition is an autogenerated conversion function.
func Convert__MySQLClusterCondition_To_v1beta2_MySQLClusterCondition(in *MySQLClusterCondition, out *v1beta2.MySQLClusterCondition, s conversion.Scope) error {
	return autoConvert__MySQLClusterCondition_To_v1beta2_MySQLClusterCondition(in, out, s)
}

func autoConvert_v1beta2_MySQLClusterCondition_To__MySQLClusterCondition(in *v1beta2.MySQLClusterCondition, out *MySQLClusterCondition, s conversion.Scope) error {
	out.Type = MySQLClusterConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastTransitionTime = in.LastTransitionTime
	return nil
}

// Convert_v1beta2_MySQLClusterCondition_To__MySQLClusterCondition is an autogenerated conversion function.
func Convert_v1beta2_MySQLClusterCondition_To__MySQLClusterCondition(in *v1beta2.MySQLClusterCondition, out *MySQLClusterCondition, s conversion.Scope) error {
	return autoConvert_v1beta2_MySQLClusterCondition_To__MySQLClusterCondition(in, out, s)
}

func autoConvert__MySQLClusterList_To_v1beta2_MySQLClusterList(in *MySQLClusterList, out *v1beta2.MySQLClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta2.MySQLCluster, len(*in))
		for i := range *in {
			if err := Convert__MySQLCluster_To_v1beta2_MySQLCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert__MySQLClusterList_To_v1beta2_MySQLClusterList is an autogenerated conversion function.
func Convert__MySQLClusterList_To_v1beta2_MySQLClusterList(in *MySQLClusterList, out *v1beta2.MySQLClusterList, s conversion.Scope) error {
	return autoConvert__MySQLClusterList_To_v1beta2_MySQLClusterList(in, out, s)
}

func autoConvert_v1beta2_MySQLClusterList_To__MySQLClusterList(in *v1beta2.MySQLClusterList, out *MySQLClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLCluster, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_MySQLCluster_To__MySQLCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_MySQLClusterList_To__MySQLClusterList is an autogenerated conversion function.
func Convert_v1beta2_MySQLClusterList_To__MySQLClusterList(in *v1beta2.MySQLClusterList, out *MySQLClusterList, s conversion.Scope) error {
	return autoConvert_v1beta2_MySQLClusterList_To__MySQLClusterList(in, out, s)
}

func autoConvert__MySQLClusterSpec_To_v1beta2_MySQLClusterSpec(in *MySQLClusterSpec, out *v1beta2.MySQLClusterSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	if err := Convert__PodTemplateSpec_To_v1beta2_PodTemplateSpec(&in.PodTemplate, &out.PodTemplate, s); err != nil {
		return err
	}
	out.VolumeClaimTemplates = *(*[]v1beta2.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	// WARNING: in.ServiceTemplate requires manual conversion: does not exist in peer-type
	out.MySQLConfigMapName = (*string)(unsafe.Pointer(in.MySQLConfigMapName))
	out.ReplicationSourceSecretName = (*string)(unsafe.Pointer(in.ReplicationSourceSecretName))
	out.Collectors = *(*[]string)(unsafe.Pointer(&in.Collectors))
	out.ServerIDBase = in.ServerIDBase
	out.MaxDelaySeconds = in.MaxDelaySeconds
	out.StartupWaitSeconds = in.StartupWaitSeconds
	out.LogRotationSchedule = in.LogRotationSchedule
	out.BackupPolicyName = (*string)(unsafe.Pointer(in.BackupPolicyName))
	out.Restore = (*v1beta2.RestoreSpec)(unsafe.Pointer(in.Restore))
	out.DisableSlowQueryLogContainer = in.DisableSlowQueryLogContainer
	return nil
}

func autoConvert_v1beta2_MySQLClusterSpec_To__MySQLClusterSpec(in *v1beta2.MySQLClusterSpec, out *MySQLClusterSpec, s conversion.Scope) error {
	out.Replicas = in.Replicas
	if err := Convert_v1beta2_PodTemplateSpec_To__PodTemplateSpec(&in.PodTemplate, &out.PodTemplate, s); err != nil {
		return err
	}
	out.VolumeClaimTemplates = *(*[]PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	// WARNING: in.PrimaryServiceTemplate requires manual conversion: does not exist in peer-type
	// WARNING: in.ReplicaServiceTemplate requires manual conversion: does not exist in peer-type
	out.MySQLConfigMapName = (*string)(unsafe.Pointer(in.MySQLConfigMapName))
	out.ReplicationSourceSecretName = (*string)(unsafe.Pointer(in.ReplicationSourceSecretName))
	out.Collectors = *(*[]string)(unsafe.Pointer(&in.Collectors))
	out.ServerIDBase = in.ServerIDBase
	out.MaxDelaySeconds = in.MaxDelaySeconds
	out.StartupWaitSeconds = in.StartupWaitSeconds
	out.LogRotationSchedule = in.LogRotationSchedule
	out.BackupPolicyName = (*string)(unsafe.Pointer(in.BackupPolicyName))
	out.Restore = (*RestoreSpec)(unsafe.Pointer(in.Restore))
	out.DisableSlowQueryLogContainer = in.DisableSlowQueryLogContainer
	return nil
}

func autoConvert__MySQLClusterStatus_To_v1beta2_MySQLClusterStatus(in *MySQLClusterStatus, out *v1beta2.MySQLClusterStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1beta2.MySQLClusterCondition)(unsafe.Pointer(&in.Conditions))
	out.CurrentPrimaryIndex = in.CurrentPrimaryIndex
	out.SyncedReplicas = in.SyncedReplicas
	out.ErrantReplicas = in.ErrantReplicas
	out.ErrantReplicaList = *(*[]int)(unsafe.Pointer(&in.ErrantReplicaList))
	if err := Convert__BackupStatus_To_v1beta2_BackupStatus(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	out.RestoredTime = (*metav1.Time)(unsafe.Pointer(in.RestoredTime))
	out.Cloned = in.Cloned
	if err := Convert__ReconcileInfo_To_v1beta2_ReconcileInfo(&in.ReconcileInfo, &out.ReconcileInfo, s); err != nil {
		return err
	}
	return nil
}

// Convert__MySQLClusterStatus_To_v1beta2_MySQLClusterStatus is an autogenerated conversion function.
func Convert__MySQLClusterStatus_To_v1beta2_MySQLClusterStatus(in *MySQLClusterStatus, out *v1beta2.MySQLClusterStatus, s conversion.Scope) error {
	return autoConvert__MySQLClusterStatus_To_v1beta2_MySQLClusterStatus(in, out, s)
}

func autoConvert_v1beta2_MySQLClusterStatus_To__MySQLClusterStatus(in *v1beta2.MySQLClusterStatus, out *MySQLClusterStatus, s conversion.Scope) error {
	out.Conditions = *(*[]MySQLClusterCondition)(unsafe.Pointer(&in.Conditions))
	out.CurrentPrimaryIndex = in.CurrentPrimaryIndex
	out.SyncedReplicas = in.SyncedReplicas
	out.ErrantReplicas = in.ErrantReplicas
	out.ErrantReplicaList = *(*[]int)(unsafe.Pointer(&in.ErrantReplicaList))
	if err := Convert_v1beta2_BackupStatus_To__BackupStatus(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	out.RestoredTime = (*metav1.Time)(unsafe.Pointer(in.RestoredTime))
	out.Cloned = in.Cloned
	if err := Convert_v1beta2_ReconcileInfo_To__ReconcileInfo(&in.ReconcileInfo, &out.ReconcileInfo, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_MySQLClusterStatus_To__MySQLClusterStatus is an autogenerated conversion function.
func Convert_v1beta2_MySQLClusterStatus_To__MySQLClusterStatus(in *v1beta2.MySQLClusterStatus, out *MySQLClusterStatus, s conversion.Scope) error {
	return autoConvert_v1beta2_MySQLClusterStatus_To__MySQLClusterStatus(in, out, s)
}

func autoConvert__ObjectMeta_To_v1beta2_ObjectMeta(in *ObjectMeta, out *v1beta2.ObjectMeta, s conversion.Scope) error {
	out.Name = in.Name
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert__ObjectMeta_To_v1beta2_ObjectMeta is an autogenerated conversion function.
func Convert__ObjectMeta_To_v1beta2_ObjectMeta(in *ObjectMeta, out *v1beta2.ObjectMeta, s conversion.Scope) error {
	return autoConvert__ObjectMeta_To_v1beta2_ObjectMeta(in, out, s)
}

func autoConvert_v1beta2_ObjectMeta_To__ObjectMeta(in *v1beta2.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	out.Name = in.Name
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_v1beta2_ObjectMeta_To__ObjectMeta is an autogenerated conversion function.
func Convert_v1beta2_ObjectMeta_To__ObjectMeta(in *v1beta2.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	return autoConvert_v1beta2_ObjectMeta_To__ObjectMeta(in, out, s)
}

func autoConvert__PersistentVolumeClaim_To_v1beta2_PersistentVolumeClaim(in *PersistentVolumeClaim, out *v1beta2.PersistentVolumeClaim, s conversion.Scope) error {
	if err := Convert__ObjectMeta_To_v1beta2_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Spec = in.Spec
	return nil
}

// Convert__PersistentVolumeClaim_To_v1beta2_PersistentVolumeClaim is an autogenerated conversion function.
func Convert__PersistentVolumeClaim_To_v1beta2_PersistentVolumeClaim(in *PersistentVolumeClaim, out *v1beta2.PersistentVolumeClaim, s conversion.Scope) error {
	return autoConvert__PersistentVolumeClaim_To_v1beta2_PersistentVolumeClaim(in, out, s)
}

func autoConvert_v1beta2_PersistentVolumeClaim_To__PersistentVolumeClaim(in *v1beta2.PersistentVolumeClaim, out *PersistentVolumeClaim, s conversion.Scope) error {
	if err := Convert_v1beta2_ObjectMeta_To__ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Spec = in.Spec
	return nil
}

// Convert_v1beta2_PersistentVolumeClaim_To__PersistentVolumeClaim is an autogenerated conversion function.
func Convert_v1beta2_PersistentVolumeClaim_To__PersistentVolumeClaim(in *v1beta2.PersistentVolumeClaim, out *PersistentVolumeClaim, s conversion.Scope) error {
	return autoConvert_v1beta2_PersistentVolumeClaim_To__PersistentVolumeClaim(in, out, s)
}

func autoConvert__PodTemplateSpec_To_v1beta2_PodTemplateSpec(in *PodTemplateSpec, out *v1beta2.PodTemplateSpec, s conversion.Scope) error {
	if err := Convert__ObjectMeta_To_v1beta2_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Spec = in.Spec
	return nil
}

// Convert__PodTemplateSpec_To_v1beta2_PodTemplateSpec is an autogenerated conversion function.
func Convert__PodTemplateSpec_To_v1beta2_PodTemplateSpec(in *PodTemplateSpec, out *v1beta2.PodTemplateSpec, s conversion.Scope) error {
	return autoConvert__PodTemplateSpec_To_v1beta2_PodTemplateSpec(in, out, s)
}

func autoConvert_v1beta2_PodTemplateSpec_To__PodTemplateSpec(in *v1beta2.PodTemplateSpec, out *PodTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta2_ObjectMeta_To__ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Spec = in.Spec
	return nil
}

// Convert_v1beta2_PodTemplateSpec_To__PodTemplateSpec is an autogenerated conversion function.
func Convert_v1beta2_PodTemplateSpec_To__PodTemplateSpec(in *v1beta2.PodTemplateSpec, out *PodTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta2_PodTemplateSpec_To__PodTemplateSpec(in, out, s)
}

func autoConvert__ReconcileInfo_To_v1beta2_ReconcileInfo(in *ReconcileInfo, out *v1beta2.ReconcileInfo, s conversion.Scope) error {
	out.Generation = in.Generation
	out.ReconcileVersion = in.ReconcileVersion
	return nil
}

// Convert__ReconcileInfo_To_v1beta2_ReconcileInfo is an autogenerated conversion function.
func Convert__ReconcileInfo_To_v1beta2_ReconcileInfo(in *ReconcileInfo, out *v1beta2.ReconcileInfo, s conversion.Scope) error {
	return autoConvert__ReconcileInfo_To_v1beta2_ReconcileInfo(in, out, s)
}

func autoConvert_v1beta2_ReconcileInfo_To__ReconcileInfo(in *v1beta2.ReconcileInfo, out *ReconcileInfo, s conversion.Scope) error {
	out.Generation = in.Generation
	out.ReconcileVersion = in.ReconcileVersion
	return nil
}

// Convert_v1beta2_ReconcileInfo_To__ReconcileInfo is an autogenerated conversion function.
func Convert_v1beta2_ReconcileInfo_To__ReconcileInfo(in *v1beta2.ReconcileInfo, out *ReconcileInfo, s conversion.Scope) error {
	return autoConvert_v1beta2_ReconcileInfo_To__ReconcileInfo(in, out, s)
}

func autoConvert__RestoreSpec_To_v1beta2_RestoreSpec(in *RestoreSpec, out *v1beta2.RestoreSpec, s conversion.Scope) error {
	out.SourceName = in.SourceName
	out.SourceNamespace = in.SourceNamespace
	out.RestorePoint = in.RestorePoint
	if err := Convert__JobConfig_To_v1beta2_JobConfig(&in.JobConfig, &out.JobConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert__RestoreSpec_To_v1beta2_RestoreSpec is an autogenerated conversion function.
func Convert__RestoreSpec_To_v1beta2_RestoreSpec(in *RestoreSpec, out *v1beta2.RestoreSpec, s conversion.Scope) error {
	return autoConvert__RestoreSpec_To_v1beta2_RestoreSpec(in, out, s)
}

func autoConvert_v1beta2_RestoreSpec_To__RestoreSpec(in *v1beta2.RestoreSpec, out *RestoreSpec, s conversion.Scope) error {
	out.SourceName = in.SourceName
	out.SourceNamespace = in.SourceNamespace
	out.RestorePoint = in.RestorePoint
	if err := Convert_v1beta2_JobConfig_To__JobConfig(&in.JobConfig, &out.JobConfig, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_RestoreSpec_To__RestoreSpec is an autogenerated conversion function.
func Convert_v1beta2_RestoreSpec_To__RestoreSpec(in *v1beta2.RestoreSpec, out *RestoreSpec, s conversion.Scope) error {
	return autoConvert_v1beta2_RestoreSpec_To__RestoreSpec(in, out, s)
}

func autoConvert__ServiceTemplate_To_v1beta2_ServiceTemplate(in *ServiceTemplate, out *v1beta2.ServiceTemplate, s conversion.Scope) error {
	if err := Convert__ObjectMeta_To_v1beta2_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Spec = (*v1.ServiceSpec)(unsafe.Pointer(in.Spec))
	return nil
}

// Convert__ServiceTemplate_To_v1beta2_ServiceTemplate is an autogenerated conversion function.
func Convert__ServiceTemplate_To_v1beta2_ServiceTemplate(in *ServiceTemplate, out *v1beta2.ServiceTemplate, s conversion.Scope) error {
	return autoConvert__ServiceTemplate_To_v1beta2_ServiceTemplate(in, out, s)
}

func autoConvert_v1beta2_ServiceTemplate_To__ServiceTemplate(in *v1beta2.ServiceTemplate, out *ServiceTemplate, s conversion.Scope) error {
	if err := Convert_v1beta2_ObjectMeta_To__ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.Spec = (*v1.ServiceSpec)(unsafe.Pointer(in.Spec))
	return nil
}

// Convert_v1beta2_ServiceTemplate_To__ServiceTemplate is an autogenerated conversion function.
func Convert_v1beta2_ServiceTemplate_To__ServiceTemplate(in *v1beta2.ServiceTemplate, out *ServiceTemplate, s conversion.Scope) error {
	return autoConvert_v1beta2_ServiceTemplate_To__ServiceTemplate(in, out, s)
}
