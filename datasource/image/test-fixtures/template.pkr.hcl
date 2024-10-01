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
      image-name = "TencentOS Server 3.1 (TK4)"
      image-type = "PUBLIC_IMAGE"
    }
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = ["source.null.basic-example"]

  provisioner "shell-local" {
    inline = [
      "echo image_id: ${data.tencentcloud-images.test_image.images[0].image_id}",
    ]
  }
}
