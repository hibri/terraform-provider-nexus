module github.com/hibri/terraform-provider-nexus

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.8.0
	github.com/hibri/nexus v0.0.0-20200401151852-2df4ad5151cd
)

replace github.com/hibri/nexus v0.0.0-20200401151852-2df4ad5151cd => ../nexus
