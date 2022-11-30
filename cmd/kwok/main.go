/*
Copyright 2022 The Kubernetes Authors.

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

package main

import (
	"os"

	"sigs.k8s.io/kwok/pkg/kwok/cmd"
	"sigs.k8s.io/kwok/pkg/log"
	"sigs.k8s.io/kwok/pkg/utils/signals"
)

func main() {
	command := cmd.NewCommand()
	ctx := signals.SetupSignalContext()
	ctx, logger := log.InitFlags(ctx, command.PersistentFlags())
	err := command.ExecuteContext(ctx)
	if err != nil {
		logger.Error("Execute exit", err)
		os.Exit(1)
	}
}
