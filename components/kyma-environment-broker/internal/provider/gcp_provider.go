package provider

import (
	"github.com/kyma-incubator/compass/components/kyma-environment-broker/internal"
	"github.com/kyma-incubator/compass/components/provisioner/pkg/gqlschema"
)

type GcpInput struct{}

func (p *GcpInput) Defaults() *gqlschema.ClusterConfigInput {
	return &gqlschema.ClusterConfigInput{
		GardenerConfig: &gqlschema.GardenerConfigInput{
			KubernetesVersion: "1.15.11",
			DiskType:          "pd-standard",
			VolumeSizeGb:      30,
			MachineType:       "n1-standard-4",
			Region:            "europe-west4",
			Provider:          "gcp",
			WorkerCidr:        "10.250.0.0/19",
			AutoScalerMin:     2,
			AutoScalerMax:     4,
			MaxSurge:          4,
			MaxUnavailable:    1,
			ProviderSpecificConfig: &gqlschema.ProviderSpecificInput{
				GcpConfig: &gqlschema.GCPProviderConfigInput{
					Zone: "europe-west4-b",
				},
			},
		},
	}
}

func (p *GcpInput) ApplyParameters(input *gqlschema.ClusterConfigInput, params internal.ProvisioningParametersDTO) {
	updateString(&input.GardenerConfig.ProviderSpecificConfig.GcpConfig.Zone, params.Zone)
}
