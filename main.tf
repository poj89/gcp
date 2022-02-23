
terraform {
  required_providers {
    gcp = {
      # source  = "local/poj89/gcp"
      version = "~> 1.0.1"
    }
  }
}

# provider "gcp" {
#   backup_id = ""
# }

data "gcp" "backup" {
    instance = google_sql_database_instance.master.name
    most_recent = true
}