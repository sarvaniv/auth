// Copyright 2017 Istio Authors
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

package na

import (
	"github.com/golang/glog"
)

type Environment int

const (
	ONPREM Environment = iota
	GCP
)

type NodeAgent interface {
	Start()
}

func NewNodeAgent(env Environment, cfg *Config) NodeAgent {
	if cfg == nil {
		glog.Fatalf("Nil configuration passed")
	}
	na := nodeAgentInternal{
		config: cfg,
	}

	switch env {
	case ONPREM:
		na.pr = &onPremPlatfromImpl{}
	default:
		glog.Fatalf("Invalid env %d specified", env)
	}

	return na
}