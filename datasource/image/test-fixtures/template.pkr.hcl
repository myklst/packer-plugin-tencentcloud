variable "secret_id" {
  type    = string
  default = "${env("TENCENTCLOUD_SECRET_ID")}"
}

variable "secret_key" {
  type    = string
  default =  "${env("TENCENTCLOUD_SECRET_KEY")}"
}

variable "region_id" {
  type    = string
  default = "ap-hongkong"
}

variable "image_name" {
  type    = string
  default = "golden-image-v1-15-5"
}

data "tencentcloud-images" "test_image" {
  secret_id = var.secret_id
  secret_key = var.secret_key
  region  = var.region_id
  filters = {
      // image-name =  "golden-image-v1-15-5"
      // image-type = ["PRIVATE_IMAGE"]
      // "tag-key" = "registrar"
      // "tag-value" = "namecheap"
      "tag:registrar" = "namecheap"
    }
  // image_ids = ["img-altiyjog","img-1az6pxke","img-its3np62"]
  // instance_type = "S1.SMALL1"
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = ["source.null.basic-example"]

  provisioner "shell-local" {
    inline = [
      // "echo image_id: ${data.tencentcloud-images.test_image.images[0].image_id}",
      "echo image_id: ${local.prodImageID[0]}",
    ]

  }
}

locals {
  prodImageID = compact(flatten([for v in data.tencentcloud-images.test_image.images :
     v.tags != null ? [for tag in v.tags :
     tag.key == "registrar" && tag.value == "namecheap"? v.image_id : null ] : []
    ]))

//   devImageID = compact(flatten([for v in data.tencentcloud-images.test_image.images : [
//     v.tags != null ? [for tag in v.tags : (
//      tag.key == "registrar" && tag.value == "namecheap")? v.image_id : null
//  ]
//   ]]))
}
