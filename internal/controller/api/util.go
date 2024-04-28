package api

import (
	"github.com/zncdata-labs/dolphinscheduler-operator/pkg/core"
	"github.com/zncdata-labs/dolphinscheduler-operator/pkg/util"
)

func createDeploymentName(instanceName string, groupName string) string {
	return util.NewResourceNameGenerator(instanceName, string(core.Api), groupName).GenerateResourceName("")
}

func createSvcName(instanceName string, groupName string) string {
	return util.NewResourceNameGenerator(instanceName, string(core.Api), groupName).GenerateResourceName("")
}

func configVolumeName() string {
	return "config"
}
