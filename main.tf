# This is required for Terraform 0.13+
terraform {
  required_providers {
    terraform-provider-gcp = {
      version = "~> 1.0.0"
      # source  = "C:\Users\mario.tejada\devWorkspace\terraform_updates\GCP\gcptest_02082022\5064-tfmod-cloudsql-google\test\integration\cloudsql\.terraform\plugins\windows_amd64\terraform-provider-gcp.exe"
    }
  }
}
data "gcp" "backup" {
    instance = google_sql_database_instance.master.name
    most_recent = true
}