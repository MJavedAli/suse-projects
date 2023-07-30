![image](https://github.com/MJavedAli/suse-projects/assets/30997178/efac35ad-81f6-4300-9bfc-b83e484b1dd2)
# GCP Virtual Machine Deployment with Terraform

This Terraform project automates the deployment of a virtual machine (VM) on Google Cloud Platform (GCP). It creates a VM instance running Debian GNU/Linux 11 (bullseye) in a new GCP project, using the provided GCP credentials for authentication.

## Prerequisites

Before you can use this Terraform configuration, ensure you have the following prerequisites:

1. **Terraform**: Install Terraform on your local machine. You can download it from the official website: https://www.terraform.io/downloads.html

2. **Google Cloud Platform (GCP) Account**: Create a GCP account if you don't have one already.

3. **GCP Project**: Create a new GCP project using the GCP Console. 

4. **GCP Service Account**: Create a service account with the necessary permissions (Compute Network Admin, Compute Instance Admin) to create and manage resources in the GCP project. Generate a JSON key for the service account and save it in a secure location.

5. **GCP Project ID**: Note down the GCP Project ID for the newly created project.

## Getting Started

1. Clone this repository to your local machine:

2. Create a `.env` file in the root directory of the project and add the following content, replacing `path/to/your/credentials.json` with the path to your GCP credentials JSON file:


Replace `your-gcp-project-id` with your GCP Project ID, and include the entire GCP credentials JSON content within the single quotes.

3. Initialize Terraform:
    ``terraform init``


4. Apply the Terraform configuration:
 ``terraform apply``

Terraform will prompt you to review the changes and ask for confirmation before proceeding with the VM deployment. Type `yes` and press Enter to proceed.

## Cleanup

To remove the resources created by Terraform (VM, network), run the following command:

``terraform destroy``
