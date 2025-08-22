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

func Tenantaccess_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		accessNameVal, ok := args["accessName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: accessName"), nil
		}
		accessName, ok := accessNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: accessName"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
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
		var requestBody models.AccessInformationUpdateParameters
		
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
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/tenant/%s%s", cfg.BaseURL, resourceGroupName, serviceName, accessName, subscriptionId, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["If-Match"]; ok {
			req.Header.Set("If-Match", fmt.Sprintf("%v", val))
		}

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
		var result interface{}
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

func CreateTenantaccess_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft_ApiManagement_service_serviceName_tenant_accessName",
		mcp.WithDescription("Update tenant access information details."),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group.")),
		mcp.WithString("serviceName", mcp.Required(), mcp.Description("The name of the API Management service.")),
		mcp.WithString("accessName", mcp.Required(), mcp.Description("The identifier of the Access configuration.")),
		mcp.WithString("If-Match", mcp.Required(), mcp.Description("ETag of the Entity. ETag should match the current entity state from the header response of the GET request or it should be * for unconditional update.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Version of the API to be used with the client request.")),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.")),
		mcp.WithBoolean("enabled", mcp.Description("Input parameter: Tenant access information of the API Management service.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Tenantaccess_updateHandler(cfg),
	}
}
