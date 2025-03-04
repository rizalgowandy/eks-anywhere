---
title: "Artifacts"
linkTitle: "Artifacts"
weight: 60
description: >
  Artifacts associated with this release: OVAs and container images.
---

# OVAs

## Bottlerocket

Bottlerocket vends its VMware variant OVAs using a secure distribution tool called tuftool. Please follow instructions down below to
download Bottlerocket OVA.
1. Install Rust and Cargo
```
curl https://sh.rustup.rs -sSf | sh
```
2. Install tuftool using Cargo
```
CARGO_NET_GIT_FETCH_WITH_CLI=true cargo install --force tuftool
```
3. Download the root role tuftool will use to download the OVA
```
curl -O "https://cache.bottlerocket.aws/root.json"
sha512sum -c <<<"90393204232a1ad6b0a45528b1f7df1a3e37493b1e05b1c149f081849a292c8dafb4ea5f7ee17bcc664e35f66e37e4cfa4aae9de7a2a28aa31ae6ac3d9bea4d5  root.json"
```
4. Export the desired Kubernetes Version. EKS Anywhere currently supports 1.21 and 1.20
```
export KUBEVERSION="1.21"
```
5. Download the OVA
```
OVA="bottlerocket-vmware-k8s-${KUBEVERSION}-x86_64-v1.2.0.ova"
tuftool download . --target-name "${OVA}" \
   --root ./root.json \
   --metadata-url "https://updates.bottlerocket.aws/2020-07-07/vmware-k8s-${KUBEVERSION}/x86_64/" \
   --targets-url "https://updates.bottlerocket.aws/targets/"
```

Bottlerocket Tags

OS Family - `os:bottlerocket`

EKS-D Release

1.21 - `eksdRelease:kubernetes-1-21-eks-4`

1.20 - `eksdRelease:kubernetes-1-20-eks-6`

## Ubuntu with Kubernetes 1.21

* https://eks-anywhere-beta.s3.amazonaws.com/0.4.0/ova/ubuntu-v1.21.2-eks-d-1-21-4-eks-a-0.4.0-amd64.ova
* `os:ubuntu`
* `eksdRelease:kubernetes-1-21-eks-4`

## Ubuntu with Kubernetes 1.20

* https://eks-anywhere-beta.s3.amazonaws.com/0.4.0/ova/ubuntu-v1.20.7-eks-d-1-20-6-eks-a-0.4.0-amd64.ova
* `os:ubuntu`
* `eksdRelease:kubernetes-1-20-eks-6`

## Building your own Ubuntu OVA
The EKS Anywhere project OVA building process leverages upstream [image-builder repository.](https://github.com/kubernetes-sigs/image-builder)
If you want to build an OVA with a custom Ubuntu base image to use for an EKS Anywhere cluster, please follow the instructions below.

Having access to a vSphere environment and docker running locally are prerequisites for building your own images.

### Required vSphere Permissions
#### Virtual machine
Inventory:
* Create new

Configuration:
* Change configuration
* Add new disk
* Add or remove device
* Change memory
* Change settings
* Set annotation

Interaction:
* Power on
* Power off
* Console interaction
* Configure CD media
* Device connection

Snapshot management:
* Create snapshot

Provisioning
* Mark as template

#### Resource Pool
* Assign vm to resource pool

#### Datastore
* Allocate space
* Browse data
* Low level file operations

#### Network
* Assign network to vm

### Steps to build an OVA
1. Spin up a builder-base docker container and exec into it. Please use the most recent tag for the image on its repository [here](https://gallery.ecr.aws/eks-distro-build-tooling/builder-base)
```
docker exec -it public.ecr.aws/eks-distro-build-tooling/builder-base:930624e251df041349f3d3089c983fcf394f1c60 bash
```
2. Clone the [eks-anywhere-build-tooling repo.](https://github.com/aws/eks-anywhere-build-tooling)
```
git clone https://github.com/aws/eks-anywhere-build-tooling.git
```
3. Navigate to the image-builder directory.
```
cd eks-anywhere-build-tooling/projects/kubernetes-sigs/image-builder
```
4. Get the vSphere connection details and create a json file named `vsphere.json` with the following template.
```
{
    "cluster": "<vSphere cluster name>",
    "datacenter": "<datacenter name on vSphere>",
    "datastore": "<datastore to be used on vSphere>",
    "folder": "<folder path to use for building ova>",
    "network": "<dhcp enabled network name>",
    "resource_pool": "<vSphere resource pool to use>",
    "vcenter_server": "<vSphere server URL>",
    "username": "<vSphere username>",
    "password": "<vSphere password>",
    "template": "",
    "insecure_connection": "false",
    "linked_clone": "false",
    "convert_to_template": "false",
    "create_snapshot": "true"
}

```
4. Export the vSphere connection data file, escaping all the quotes
```
export VSPHERE_CONNECTION_DATA=\"$(cat vsphere.json | jq -c . | sed 's/"/\\"/g')\"
```
5. Download the most recent release bundle manifest and get the latest URLs for `etcdadm` and `crictl` for the intended Kubernetes version.
```
wget https://anywhere-assets.eks.amazonaws.com/bundle-release.yaml
```
7. Export the CRICTL_URL and ETCADM_HTTP_SOURCE environment variables with the URLs from previous step.
```
export CRICTL_URL=<crictl url>
export ETCDADM_HTTP_SOURCE=<etcdadm url>
```
7. Create a library on vSphere for image-builder.
```
govc library.create "CodeBuild"
```
8. Update the Ubuntu configuration file with the new custom ISO URL and its checksum at
`image-builder/images/capi/packer/ova/ubuntu-2004.json`
9. Setup image-builder and run the OVA build for the Kubernetes version.
```
make release-ova-ubuntu-2004-1-21
```

# Images

The various images for EKS Anywhere can be found [in the EKS Anywhere ECR repository](https://gallery.ecr.aws/eks-anywhere/).
The various images for EKS Distro can be found [in the EKS Distro ECR repository](https://gallery.ecr.aws/eks-distro/).
