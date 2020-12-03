package octopusdeploy

import (
	"context"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandProjectGroup(d *schema.ResourceData) *octopusdeploy.ProjectGroup {
	name := d.Get("name").(string)

	projectGroup := octopusdeploy.NewProjectGroup(name)
	projectGroup.ID = d.Id()

	if v, ok := d.GetOk("description"); ok {
		projectGroup.Description = v.(string)
	}

	if v, ok := d.GetOk("environments"); ok {
		projectGroup.EnvironmentIDs = getSliceFromTerraformTypeList(v)
	}

	if v, ok := d.GetOk("retention_policy_id"); ok {
		projectGroup.RetentionPolicyID = v.(string)
	}

	return projectGroup
}

func flattenProjectGroup(projectGroup *octopusdeploy.ProjectGroup) map[string]interface{} {
	if projectGroup == nil {
		return nil
	}

	return map[string]interface{}{
		"description":         projectGroup.Description,
		"environments":        projectGroup.EnvironmentIDs,
		"id":                  projectGroup.GetID(),
		"name":                projectGroup.Name,
		"retention_policy_id": projectGroup.RetentionPolicyID,
	}
}

func getProjectGroupDataSchema() map[string]*schema.Schema {
	dataSchema := getProjectGroupSchema()
	setDataSchema(&dataSchema)

	return map[string]*schema.Schema{
		"id":           getIDDataSchema(),
		"ids":          getIDsQuery(),
		"partial_name": getPartialNameQuery(),
		"project_groups": {
			Computed:    true,
			Description: "A list of project groups that match the filter(s).",
			Elem:        &schema.Resource{Schema: dataSchema},
			Optional:    true,
			Type:        schema.TypeList,
		},
		"skip": getSkipQuery(),
		"take": getTakeQuery(),
	}
}

func getProjectGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description":  getDescriptionSchema(),
		"environments": getEnvironmentsSchema(),
		"id":           getIDSchema(),
		"name":         getNameSchema(true),
		"retention_policy_id": {
			Optional: true,
			Type:     schema.TypeString,
		},
	}
}

func setProjectGroup(ctx context.Context, d *schema.ResourceData, projectGroup *octopusdeploy.ProjectGroup) {
	d.Set("description", projectGroup.Description)
	d.Set("environments", projectGroup.EnvironmentIDs)
	d.Set("name", projectGroup.Name)
	d.Set("retention_policy_id", projectGroup.RetentionPolicyID)

	d.SetId(projectGroup.GetID())
}
