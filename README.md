# packer-plugin-tencentcloud

## Inputs

### Required:
|    Name   |  Type  |
|-----------|--------|
|secret_id  | string |
|secret_key | string |
|region     | string |

### Optional:

For more details, refer to this [docs](https://www.tencentcloud.com/document/product/213/33272).

|    Name      | Type          | Description                                                                                            |
|--------------|---------------|--------------------------------------------------------------------------------------------------------|
|image_ids     | string        | list of ID of the image.                                                                                       |
|filters       | map of string | <pre>image-id   = img-tes1r4 <br>image-type = PRIVATE_IMAGE <br>image-name = webserver<br>platform   = CentOS <br>tag-key    = env <br>tag-value  = prod <br>tag:env    = prod</pre>|
|instance_type | string        | Instance type, `e.g. S1.SMALL1`                                                                                  |


## Outputs
|    Name     | Type           |
|-------------|----------------|
|images       | <pre>list of object([{<br>  image_id     = string<br>  image_name   = string<br>  instance_type = string<br>  tags         = list of object([{<br>    key   = string<br>    value = string<br>  }])<br>}])</pre> |


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

data "st-tencentcloud-images" "test_image" {
  secret_id  = "v1-gastisthisisnotmyaccesskey"
  secret_key = "v9-adftthisfathisisnotmysecretkey"
  region  = "cn-hongkong"
  filters = {
      "image-name"    = "img-5566"
      "image-type"    = "PRIVATE_IMAGE"
      "tag-key"       = "registrar"
      "tag-value"     = "namecheap"
      "tag:registrar" = "namecheap"
    }
  image_ids = ["img-altiyjog","img-1az6pxke","img-its3np62"]
  instance_type = "S1.SMALL1"
}

locals {
  prodImageID = compact(flatten([for v in data.st-tencentcloud-images.test_image.images :
     v.tags != null ? [for tag in v.tags :
     tag.key == "registrar" && tag.value == "namecheap"? v.image_id : null ] : []
    ]))
}
```
