// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package memorydb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/memorydb"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceACL,
			TypeName: "aws_memorydb_acl",
			Name:     "ACL",
		},
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_memorydb_cluster",
			Name:     "Cluster",
		},
		{
			Factory:  dataSourceParameterGroup,
			TypeName: "aws_memorydb_parameter_group",
			Name:     "Parameter Group",
		},
		{
			Factory:  dataSourceSnapshot,
			TypeName: "aws_memorydb_snapshot",
			Name:     "Snapshot",
		},
		{
			Factory:  dataSourceSubnetGroup,
			TypeName: "aws_memorydb_subnet_group",
			Name:     "Subnet Group",
		},
		{
			Factory:  dataSourceUser,
			TypeName: "aws_memorydb_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceACL,
			TypeName: "aws_memorydb_acl",
			Name:     "ACL",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_memorydb_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceParameterGroup,
			TypeName: "aws_memorydb_parameter_group",
			Name:     "Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSnapshot,
			TypeName: "aws_memorydb_snapshot",
			Name:     "Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSubnetGroup,
			TypeName: "aws_memorydb_subnet_group",
			Name:     "Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_memorydb_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.MemoryDB
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*memorydb.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return memorydb.NewFromConfig(cfg,
		memorydb.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
