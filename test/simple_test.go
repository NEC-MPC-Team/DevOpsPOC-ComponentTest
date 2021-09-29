package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSimpleTerraform(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform",
		VarFiles: []string{
			"../envs/dev/variables.tfvars",
		},
	}

	terraform.InitAndApply(t, terraformOptions)

	t.Run("Helloworld", func(t *testing.T) {
		test1_output := terraform.Output(t, terraformOptions, "hello_world")
		assert.Equal(t, "more, more, more", test1_output)
	})

	t.Run("Resource group location", func(t *testing.T) {
		test2_output := terraform.Output(t, terraformOptions, "rg_location")
		assert.Equal(t, "australiaeast", test2_output)
	})

	t.Run("Resource Group Exists", func(t *testing.T) {
		resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
		subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
		exists := azure.ResourceGroupExists(t, resourceGroupName, subscriptionId)
		assert.True(t, exists, "Resource group does not exist")
	})

	defer terraform.Destroy(t, terraformOptions)
}
