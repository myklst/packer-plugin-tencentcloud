// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package datasource

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	SecretId       *string           `mapstructure:"secret_id" required:"true" cty:"secret_id" hcl:"secret_id"`
	SecretKey      *string           `mapstructure:"secret_key" required:"true" cty:"secret_key" hcl:"secret_key"`
	Region         *string           `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	ImageIds       []string          `mapstructure:"image_ids" cty:"image_ids" hcl:"image_ids"`
	ImageNameRegex *string           `mapstructure:"image_name_regex" cty:"image_name_regex" hcl:"image_name_regex"`
	Filters        map[string]string `mapstructure:"filters" cty:"filters" hcl:"filters"`
	InstanceType   *string           `mapstructure:"instance_type" cty:"instance_type" hcl:"instance_type"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"secret_id":        &hcldec.AttrSpec{Name: "secret_id", Type: cty.String, Required: false},
		"secret_key":       &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"region":           &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"image_ids":        &hcldec.AttrSpec{Name: "image_ids", Type: cty.List(cty.String), Required: false},
		"image_name_regex": &hcldec.AttrSpec{Name: "image_name_regex", Type: cty.String, Required: false},
		"filters":          &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"instance_type":    &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
	}
	return s
}

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	Images []FlatImage `mapstructure:"images" cty:"images" hcl:"images"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"images": &hcldec.BlockListSpec{TypeName: "images", Nested: hcldec.ObjectSpec((*FlatImage)(nil).HCL2Spec())},
	}
	return s
}

// FlatImage is an auto-generated flat version of Image.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatImage struct {
	ImageId      *string   `mapstructure:"image_id" cty:"image_id" hcl:"image_id"`
	ImageName    *string   `mapstructure:"image_name" cty:"image_name" hcl:"image_name"`
	InstanceType *string   `mapstructure:"instance_type" cty:"instance_type" hcl:"instance_type"`
	Tags         []FlatTag `mapstructure:"tags" cty:"tags" hcl:"tags"`
}

// FlatMapstructure returns a new FlatImage.
// FlatImage is an auto-generated flat version of Image.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Image) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatImage)
}

// HCL2Spec returns the hcl spec of a Image.
// This spec is used by HCL to read the fields of Image.
// The decoded values from this spec will then be applied to a FlatImage.
func (*FlatImage) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"image_id":      &hcldec.AttrSpec{Name: "image_id", Type: cty.String, Required: false},
		"image_name":    &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"instance_type": &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"tags":          &hcldec.BlockListSpec{TypeName: "tags", Nested: hcldec.ObjectSpec((*FlatTag)(nil).HCL2Spec())},
	}
	return s
}

// FlatTag is an auto-generated flat version of Tag.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatTag struct {
	TagKey   *string `mapstructure:"key" cty:"key" hcl:"key"`
	TagValue *string `mapstructure:"value" cty:"value" hcl:"value"`
}

// FlatMapstructure returns a new FlatTag.
// FlatTag is an auto-generated flat version of Tag.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Tag) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatTag)
}

// HCL2Spec returns the hcl spec of a Tag.
// This spec is used by HCL to read the fields of Tag.
// The decoded values from this spec will then be applied to a FlatTag.
func (*FlatTag) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"key":   &hcldec.AttrSpec{Name: "key", Type: cty.String, Required: false},
		"value": &hcldec.AttrSpec{Name: "value", Type: cty.String, Required: false},
	}
	return s
}
