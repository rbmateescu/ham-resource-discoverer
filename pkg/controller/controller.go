// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"github.com/hybridapp-io/ham-resource-discoverer/pkg/synchronizer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManagerFuncs is a list of non-synchronized functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager, *rest.Config, types.NamespacedName) error

// AddToManagerSyncFuncs is a list of synchronized functions to add all Controllers to the Manager
var AddToManagerSyncFuncs []func(manager.Manager, *rest.Config, types.NamespacedName, synchronizer.HubSynchronizerInterface) error

// AddToManager adds all controllers to the Manager
func AddToManager(m manager.Manager, hubconfig *rest.Config, cluster types.NamespacedName,
	hubSynchronizer synchronizer.HubSynchronizerInterface) error {
	for _, f := range AddToManagerFuncs {
		if err := f(m, hubconfig, cluster); err != nil {
			return err
		}
	}

	for _, f := range AddToManagerSyncFuncs {
		if err := f(m, hubconfig, cluster, hubSynchronizer); err != nil {
			return err
		}
	}
	return nil
}
