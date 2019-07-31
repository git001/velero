/*
Copyright 2017, 2019 the Velero contributors.

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

package builder

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	velerov1api "github.com/heptio/velero/pkg/apis/velero/v1"
)

// BackupStorageLocationBuilder builds BackupStorageLocation objects.
type BackupStorageLocationBuilder struct {
	object *velerov1api.BackupStorageLocation
}

// ForBackupStorageLocation is the constructor for a BackupStorageLocationBuilder.
func ForBackupStorageLocation(ns, name string) *BackupStorageLocationBuilder {
	return &BackupStorageLocationBuilder{
		object: &velerov1api.BackupStorageLocation{
			TypeMeta: metav1.TypeMeta{
				APIVersion: velerov1api.SchemeGroupVersion.String(),
				Kind:       "BackupStorageLocation",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns,
				Name:      name,
			},
		},
	}
}

// Result returns the built BackupStorageLocation.
func (b *BackupStorageLocationBuilder) Result() *velerov1api.BackupStorageLocation {
	return b.object
}

// ObjectMeta applies functional options to the BackupStorageLocation's ObjectMeta.
func (b *BackupStorageLocationBuilder) ObjectMeta(opts ...ObjectMetaOpt) *BackupStorageLocationBuilder {
	for _, opt := range opts {
		opt(b.object)
	}

	return b
}

// Provider sets the BackupStorageLocation's provider.
func (b *BackupStorageLocationBuilder) Provider(name string) *BackupStorageLocationBuilder {
	b.object.Spec.Provider = name
	return b
}

// ObjectStorage sets the BackupStorageLocation's object storage.
func (b *BackupStorageLocationBuilder) ObjectStorage(bucketName string) *BackupStorageLocationBuilder {
	if b.object.Spec.StorageType.ObjectStorage == nil {
		b.object.Spec.StorageType.ObjectStorage = new(velerov1api.ObjectStorageLocation)
	}
	b.object.Spec.ObjectStorage.Bucket = bucketName
	return b
}

// AccessMode sets the BackupStorageLocation's access mode.
func (b *BackupStorageLocationBuilder) AccessMode(accessMode velerov1api.BackupStorageLocationAccessMode) *BackupStorageLocationBuilder {
	b.object.Spec.AccessMode = accessMode
	return b
}