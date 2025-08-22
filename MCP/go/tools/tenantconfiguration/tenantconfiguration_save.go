package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Tenantconfiguration_saveHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		resourceGroupNameVal, ok := args["resourceGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroupName"), nil
		}
		resourceGroupName, ok := resourceGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroupName"), nil
		}
		serviceNameVal, ok := args["serviceName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serviceName"), nil
		}
		serviceName, ok := serviceNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serviceName"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		configurationNameVal, ok := args["configurationName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: configurationName"), nil
		}
		configurationName, ok := configurationNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: configurationName"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.SaveConfigurationParameter
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/tenant/%s/save%s", cfg.BaseURL, resourceGroupName, serviceName, subscriptionId, configurationName, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.OperationResultContract
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateTenantconfiguration_saveTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft_ApiManagement_service_serviceName_tenant_configurationName_save",
		mcp.WithDescription("This operation creates a commit with the current configuration snapshot to the specified branch in the repository. This is a long running operation and could take several minutes to complete."),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group.")),
		mcp.WithString("serviceName", mcp.Required(), mcp.Description("The name of the API Management service.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Version of the API to be used with the client request.")),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.")),
		mcp.WithString("configurationName", mcp.Required(), mcp.Description("The identifier of the Git Configuration Operation.")),
		mcp.WithString("branch", mcp.Required(), mcp.Description("Input parameter: The name of the Git branch in which to commit the current configuration snapshot.")),
		mcp.WithBoolean("force", mcp.Description("Input parameter: The value if true, the current configuration database is committed to the Git repository, even if the Git repository has newer changes that would be overwritten.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Tenantconfiguration_saveHandler(cfg),
	}
}
