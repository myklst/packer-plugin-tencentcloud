//go:generate packer-sdc mapstructure-to-hcl2 -type DatasourceOutput,Image,Tag,Config
package datasource

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Datasource struct {
	config Config
}

type Config struct {
	SecretId       string            `mapstructure:"secret_id" required:"true"`
	SecretKey      string            `mapstructure:"secret_key" required:"true"`
	Region         string            `mapstructure:"region" required:"true"`
	ImageIds       []string          `mapstructure:"image_ids"`
	ImageNameRegex string            `mapstructure:"image_name_regex"`
	Filters        map[string]string `mapstructure:"filters"`
	InstanceType   string            `mapstructure:"instance_type"`
}

type DatasourceOutput struct {
	Images []Image `mapstructure:"images"`
}

type Image struct {
	ImageId      string `mapstructure:"image_id"`
	ImageName    string `mapstructure:"image_name"`
	InstanceType string `mapstructure:"instance_type"`
	Tags         []Tag  `mapstructure:"tags"`
}

type Tag struct {
	TagKey   string `mapstructure:"key"`
	TagValue string `mapstructure:"value"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	if err := config.Decode(&d.config, nil, raws...); err != nil {
		return fmt.Errorf("error parsing configuration: %v", err)
	}

	var errs *packer.MultiError
	if d.config.SecretId == "" {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("secret_id is missing"))
	}

	if d.config.SecretKey == "" {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("secret_key is missing"))
	}

	if d.config.Region == "" {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("region is missing"))
	}

	if errs != nil && len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func CreateClient(d *Datasource) (client *cvm.Client, err error) {
	credential := common.NewCredential(
		d.config.SecretId,
		d.config.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"

	return cvm.NewClient(credential, d.config.Region, cpf)
}

func (d *Datasource) Execute() (cty.Value, error) {
	client, err := CreateClient(d)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	request := cvm.NewDescribeImagesRequest()

	var filters []*cvm.Filter
	for key, value := range d.config.Filters {
		filter := &cvm.Filter{
			Name:   common.StringPtr(key),
			Values: common.StringPtrs([]string{value}),
		}
		filters = append(filters, filter)
	}
	request.Filters = filters

	// If the imageIds and instanceType inputs are null, it will cause an SDK error.
	// Null/Empty checking
	if len(d.config.ImageIds) > 0 {
		request.ImageIds = common.StringPtrs(d.config.ImageIds)
	}

	if d.config.InstanceType != "" {
		request.InstanceType = common.StringPtr(d.config.InstanceType)
	}

	// Get Images
	resp, err := client.DescribeImages(request)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	filteredImages, err := getFilteredImage(d, resp)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	var dataOutput DatasourceOutput
	dataOutput.Images = filteredImages

	return hcl2helper.HCL2ValueFromConfig(dataOutput, d.OutputSpec()), nil
}

func getFilteredImage(d *Datasource, resp *cvm.DescribeImagesResponse) (images []Image, err error) {
	if *resp.Response.TotalCount == 0 {
		return images, fmt.Errorf("no image found matching the filters")
	}

	for _, img := range resp.Response.ImageSet {
		// check if image name's regex is specified
		if d.config.ImageNameRegex != "" {
			imageNameRegex, err := regexp.Compile(d.config.ImageNameRegex)
			if err != nil {
				return images, err
			}
			if !imageNameRegex.MatchString(*img.ImageName) {
				continue
			}
		}

		var tags []Tag
		for _, imgTag := range img.Tags {
			tag := Tag{
				TagKey:   *imgTag.Key,
				TagValue: *imgTag.Value,
			}
			tags = append(tags, tag)
		}

		images = append(images, Image{
			ImageId:      *img.ImageId,
			ImageName:    *img.ImageName,
			InstanceType: *img.ImageType,
			Tags:         tags,
		})
	}

	return images, nil
}
