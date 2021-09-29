resource "azurerm_resource_group" "target_rg" {
  name     = var.rg_name
  location = var.location
}