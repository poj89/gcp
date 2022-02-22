package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"google_sql_backup_run": dataSourceSqlBackupRun(),
			// "google_sql_database_instance": dataSourceSqlDatabaseInstance(),
		},
	}
}

func dataSourceSqlBackupRun() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceSqlBackupRunRead,

		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The identifier for this backup run. Unique only for a specific Cloud SQL instance. If left empty and multiple backups exist for the instance, most_recent must be set to true.`,
			},
			"instance": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the database instance.`,
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Location of the backups.`,
			},
			"start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the backup operation actually started in UTC timezone in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.`,
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The status of this run.`,
			},
			"most_recent": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Toggles use of the most recent backup run if multiple backups exist for a Cloud SQL instance.`,
			},
		},
	}
}

func dataSourceSqlBackupRunRead(d *schema.ResourceData, m interface{}) error {
	instance_id := d.Get("instance").(string)
	d.SetId(instance_id)
	return nil
}
