package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeployConfigurationParameters represents the DeployConfigurationParameters schema from the OpenAPI specification
type DeployConfigurationParameters struct {
	Branch string `json:"branch"` // The name of the Git branch from which the configuration is to be deployed to the configuration database.
	Force bool `json:"force,omitempty"` // The value enforcing deleting subscriptions to products that are deleted in this update.
}

// OperationResultContract represents the OperationResultContract schema from the OpenAPI specification
type OperationResultContract struct {
	Status string `json:"status,omitempty"` // Status of an async operation.
	Updated string `json:"updated,omitempty"` // Last update time of an async operation. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Actionlog []OperationResultLogItemContract `json:"actionLog,omitempty"` // This property if only provided as part of the TenantConfiguration_Validate operation. It contains the log the entities which will be updated/created/deleted as part of the TenantConfiguration_Deploy operation.
	ErrorField interface{} `json:"error,omitempty"` // Error Body contract.
	Id string `json:"id,omitempty"` // Operation result identifier.
	Resultinfo string `json:"resultInfo,omitempty"` // Optional result info.
	Started string `json:"started,omitempty"` // Start time of an async operation. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
}

// OperationResultLogItemContract represents the OperationResultLogItemContract schema from the OpenAPI specification
type OperationResultLogItemContract struct {
	Objectkey string `json:"objectKey,omitempty"` // Identifier of the entity being created/updated/deleted.
	Objecttype string `json:"objectType,omitempty"` // The type of entity contract.
	Action string `json:"action,omitempty"` // Action like create/update/delete.
}

// SaveConfigurationParameter represents the SaveConfigurationParameter schema from the OpenAPI specification
type SaveConfigurationParameter struct {
	Branch string `json:"branch"` // The name of the Git branch in which to commit the current configuration snapshot.
	Force bool `json:"force,omitempty"` // The value if true, the current configuration database is committed to the Git repository, even if the Git repository has newer changes that would be overwritten.
}

// TenantConfigurationSyncStateContract represents the TenantConfigurationSyncStateContract schema from the OpenAPI specification
type TenantConfigurationSyncStateContract struct {
	Branch string `json:"branch,omitempty"` // The name of Git branch.
	Commitid string `json:"commitId,omitempty"` // The latest commit Id.
	Configurationchangedate string `json:"configurationChangeDate,omitempty"` // The date of the latest configuration change. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Isexport bool `json:"isExport,omitempty"` // value indicating if last sync was save (true) or deploy (false) operation.
	Isgitenabled bool `json:"isGitEnabled,omitempty"` // value indicating whether Git configuration access is enabled.
	Issynced bool `json:"isSynced,omitempty"` // value indicating if last synchronization was later than the configuration change.
	Syncdate string `json:"syncDate,omitempty"` // The date of the latest synchronization. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
}

// AccessInformationContract represents the AccessInformationContract schema from the OpenAPI specification
type AccessInformationContract struct {
	Primarykey string `json:"primaryKey,omitempty"` // Primary access key.
	Secondarykey string `json:"secondaryKey,omitempty"` // Secondary access key.
	Enabled bool `json:"enabled,omitempty"` // Tenant access information of the API Management service.
	Id string `json:"id,omitempty"` // Identifier.
}

// AccessInformationUpdateParameters represents the AccessInformationUpdateParameters schema from the OpenAPI specification
type AccessInformationUpdateParameters struct {
	Enabled bool `json:"enabled,omitempty"` // Tenant access information of the API Management service.
}
