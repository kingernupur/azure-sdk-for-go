//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/Firewall/stable/2023-12-30/examples/FirewallRuleCreate.json
func ExampleFirewallRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFirewallRulesClient().BeginCreateOrUpdate(ctx, "TestGroup", "testserver", "rule1", armmysqlflexibleservers.FirewallRule{
		Properties: &armmysqlflexibleservers.FirewallRuleProperties{
			EndIPAddress:   to.Ptr("255.255.255.255"),
			StartIPAddress: to.Ptr("0.0.0.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallRule = armmysqlflexibleservers.FirewallRule{
	// 	Name: to.Ptr("rule1"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/firewallRules"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/firewallRules/rule1"),
	// 	Properties: &armmysqlflexibleservers.FirewallRuleProperties{
	// 		EndIPAddress: to.Ptr("255.255.255.255"),
	// 		StartIPAddress: to.Ptr("0.0.0.0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/Firewall/stable/2023-12-30/examples/FirewallRuleDelete.json
func ExampleFirewallRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFirewallRulesClient().BeginDelete(ctx, "TestGroup", "testserver", "rule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/Firewall/stable/2023-12-30/examples/FirewallRuleGet.json
func ExampleFirewallRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallRulesClient().Get(ctx, "TestGroup", "testserver", "rule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallRule = armmysqlflexibleservers.FirewallRule{
	// 	Name: to.Ptr("rule1"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/firewallRules"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/firewallRules/rule1"),
	// 	Properties: &armmysqlflexibleservers.FirewallRuleProperties{
	// 		EndIPAddress: to.Ptr("255.255.255.255"),
	// 		StartIPAddress: to.Ptr("0.0.0.0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/Firewall/stable/2023-12-30/examples/FirewallRulesListByServer.json
func ExampleFirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallRulesClient().NewListByServerPager("TestGroup", "testserver", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FirewallRuleListResult = armmysqlflexibleservers.FirewallRuleListResult{
		// 	Value: []*armmysqlflexibleservers.FirewallRule{
		// 		{
		// 			Name: to.Ptr("rule1"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/firewallRules/rule1"),
		// 			Properties: &armmysqlflexibleservers.FirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("255.255.255.255"),
		// 				StartIPAddress: to.Ptr("0.0.0.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("rule2"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/firewallRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/firewallRules/rule2"),
		// 			Properties: &armmysqlflexibleservers.FirewallRuleProperties{
		// 				EndIPAddress: to.Ptr("255.0.0.0"),
		// 				StartIPAddress: to.Ptr("1.0.0.0"),
		// 			},
		// 	}},
		// }
	}
}
