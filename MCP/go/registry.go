package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_tenantaccess "github.com/apimanagementclient/mcp-server/tools/tenantaccess"
	tools_tenantconfiguration "github.com/apimanagementclient/mcp-server/tools/tenantconfiguration"
	tools_tenantconfigurationsyncstate "github.com/apimanagementclient/mcp-server/tools/tenantconfigurationsyncstate"
	tools_tenantaccessgit "github.com/apimanagementclient/mcp-server/tools/tenantaccessgit"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tenantaccess.CreateTenantaccess_regeneratesecondarykeyTool(cfg),
		tools_tenantconfiguration.CreateTenantconfiguration_saveTool(cfg),
		tools_tenantconfigurationsyncstate.CreateTenantconfiguration_getsyncstateTool(cfg),
		tools_tenantaccess.CreateTenantaccess_getTool(cfg),
		tools_tenantaccess.CreateTenantaccess_updateTool(cfg),
		tools_tenantaccessgit.CreateTenantaccessgit_regenerateprimarykeyTool(cfg),
		tools_tenantconfiguration.CreateTenantconfiguration_validateTool(cfg),
		tools_tenantaccessgit.CreateTenantaccessgit_getTool(cfg),
		tools_tenantaccessgit.CreateTenantaccessgit_regeneratesecondarykeyTool(cfg),
		tools_tenantaccess.CreateTenantaccess_regenerateprimarykeyTool(cfg),
		tools_tenantconfiguration.CreateTenantconfiguration_deployTool(cfg),
	}
}
