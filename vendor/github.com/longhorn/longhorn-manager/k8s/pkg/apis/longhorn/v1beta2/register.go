package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

var SchemeGroupVersion = schema.GroupVersion{Group: longhorn.GroupName, Version: "v1beta2"}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&BackingImage{},
		&BackingImageList{},
		&BackingImageDataSource{},
		&BackingImageDataSourceList{},
		&BackingImageManager{},
		&BackingImageManagerList{},
		&Backup{},
		&BackupList{},
		&BackupTarget{},
		&BackupTargetList{},
		&BackupVolume{},
		&BackupVolumeList{},
		&Engine{},
		&EngineList{},
		&EngineImage{},
		&EngineImageList{},
		&InstanceManager{},
		&InstanceManagerList{},
		&Node{},
		&NodeList{},
		&Orphan{},
		&OrphanList{},
		&RecurringJob{},
		&RecurringJobList{},
		&Replica{},
		&ReplicaList{},
		&Setting{},
		&SettingList{},
		&ShareManager{},
		&ShareManagerList{},
		&Snapshot{},
		&SnapshotList{},
		&SupportBundle{},
		&SupportBundleList{},
		&SystemBackup{},
		&SystemBackupList{},
		&SystemRestore{},
		&SystemRestoreList{},
		&Volume{},
		&VolumeList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
