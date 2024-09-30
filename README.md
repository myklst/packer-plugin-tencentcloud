# packer-plugin-tencentcloud

## Inputs

### Required:
|    Name   |  Type  |
|-----------|--------|
|secret_id  | string |
|secret_key | string |
|region     | string |

### Optional:

For more details, refer to this [Docs](https://www.tencentcloud.com/document/product/213/33272).<br>
ImageIds and Filters cannot be specified at the same time.*

|    Name      | Type          | Description                                                                                            |
|--------------|---------------|--------------------------------------------------------------------------------------------------------|
|image_ids     | list of string        | list of ID of the image.                                                                                       |
|image_name_regex | string        | Regex that is used to query image by name.                                                                                       |
|filters       | map of string | <pre>image_id   = "img-tes1r4" <br>image-type = "PRIVATE_IMAGE" <br>image-name = "webserver" // full image name<br>platform   = "CentOS" <br>tag-key    = "env" <br>tag-value  = "prod" <br>tag:env    = "prod" // tag:[key] = "value"</pre>|
|instance_type | string        | Instance type, `e.g. S1.SMALL1`                                                                                  |


## Outputs
|    Name     | Type           |
|-------------|----------------|
|images       | <pre>list of object([{<br>  image_id      = string<br>  image_name    = string<br>  instance_type = string<br>  tags          = list of object([{<br>      key   = string<br>      value = string<br>   }])<br>}])</pre> |


## Example
```
packer {
  required_plugins {
     st-tencentcloud = {
      source  = "github.com/myklst/tencentcloud"
      version = "~> 0.1"
    }
  }
}

// Example inputs

data "st-tencentcloud-images" "test_image" {
  secret_id  = "v1-gastisthisisnotmyaccesskey"
  secret_key = "v9-adftthisfathisisnotmysecretkey"
  region  = "cn-hongkong"
  image_name_regex = "^TencentOS\\s+Server\\s+\\d+\\.\\d+\\s+\\(TK4\\)$"
  // image_ids = ["img-altiyjog","img-1az6pxke","img-its3np62"]
  filters = {
      "image-name"    = "img-5566"
      "image-type"    = "PRIVATE_IMAGE"
      "tag-key"       = "registrar"
      "tag-value"     = "namecheap"
      "tag:registrar" = "namecheap" // tag:[key] = "value"
    }
  instance_type = "S1.SMALL1"
}

build {
  sources = ["source.null.basic-example"]

  provisioner "shell-local" {
    inline = [
      "echo image_id: ${data.tencentcloud-images.test_image.images[0].image_id}",
    ]
  }
}
```
