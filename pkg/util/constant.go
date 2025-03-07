/*
Copyright 2022 The Kruise Authors.

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

package util

// For Rollout and BatchRelease
const (
	// BatchReleaseControlAnnotation is controller info about batchRelease when rollout
	BatchReleaseControlAnnotation = "batchrelease.rollouts.kruise.io/control-info"
	// InRolloutProgressingAnnotation marks workload as entering the rollout progressing process
	//and does not allow paused=false during this process
	InRolloutProgressingAnnotation = "rollouts.kruise.io/in-progressing"
	// RolloutHashAnnotation record observed rollout spec hash
	RolloutHashAnnotation = "rollouts.kruise.io/hash"
	// RollbackInBatchAnnotation allow use disable quick rollback, and will roll back in batch style.
	RollbackInBatchAnnotation = "rollouts.kruise.io/rollback-in-batch"
)

// For Workloads
const (
	// CanaryDeploymentLabel is to label canary deployment that is created by batchRelease controller
	CanaryDeploymentLabel = "rollouts.kruise.io/canary-deployment"
	// CanaryDeploymentFinalizer is a finalizer to resources patched by batchRelease controller
	CanaryDeploymentFinalizer = "finalizer.rollouts.kruise.io/batch-release"
	// KruiseRolloutFinalizer is a finalizer for deployment/service/ingress/gateway/etc
	KruiseRolloutFinalizer = "rollouts.kruise.io/rollout"
	// WorkloadTypeLabel is a label to identify workload type
	WorkloadTypeLabel = "rollouts.kruise.io/workload-type"
)

// For Pods
const (
	// RolloutIDLabel is designed to distinguish each workload revision publications.
	// The value of RolloutIDLabel corresponds Rollout.Spec.RolloutID.
	RolloutIDLabel = "rollouts.kruise.io/rollout-id"
	// RolloutBatchIDLabel is the label key of batch id that will be patched to pods during rollout.
	// Only when RolloutIDLabel is set, RolloutBatchIDLabel will be patched.
	// Users can use RolloutIDLabel and RolloutBatchIDLabel to select the pods that are upgraded in some certain batch and release.
	RolloutBatchIDLabel = "rollouts.kruise.io/rollout-batch-id"
	// NoNeedUpdatePodLabel will be patched to pod when rollback in batches if the pods no need to rollback
	NoNeedUpdatePodLabel = "rollouts.kruise.io/no-need-update"
)

// For Others
const (
	// We omit vowels from the set of available characters to reduce the chances
	// of "bad words" being formed.
	alphanums = "bcdfghjklmnpqrstvwxz2456789"

	// CloneSetType DeploymentType and StatefulSetType are values to WorkloadTypeLabel
	CloneSetType    WorkloadType = "cloneset"
	DeploymentType  WorkloadType = "deployment"
	StatefulSetType WorkloadType = "statefulset"

	AddFinalizerOpType    FinalizerOpType = "Add"
	RemoveFinalizerOpType FinalizerOpType = "Remove"
)

type WorkloadType string

type FinalizerOpType string
