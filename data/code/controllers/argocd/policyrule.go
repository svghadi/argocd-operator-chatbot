package argocd

import (
	"fmt"

	"golang.org/x/mod/semver"

	"github.com/argoproj-labs/argocd-operator/common"

	v1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func policyRuleForApplicationController() []v1.PolicyRule {

	return []v1.PolicyRule{
		{
			APIGroups: []string{
				"*",
			},
			Resources: []string{
				"*",
			},
			Verbs: []string{
				"*",
			},
		},
	}
}

func policyRuleForRedis(client client.Client) []v1.PolicyRule {
	rules := []v1.PolicyRule{}

	// Need additional policy rules if we are running on openshift, else the stateful set won't have the right
	// permissions to start
	rules = appendOpenShiftNonRootSCC(rules, client)

	return rules
}

func policyRuleForRedisHa(client client.Client) []v1.PolicyRule {

	rules := []v1.PolicyRule{
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"endpoints",
			},
			Verbs: []string{
				"get",
			},
		},
	}

	// Need additional policy rules if we are running on openshift, else the stateful set won't have the right
	// permissions to start
	rules = appendOpenShiftNonRootSCC(rules, client)

	return rules
}

func policyRuleForDexServer() []v1.PolicyRule {

	return []v1.PolicyRule{
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"secrets",
				"configmaps",
			},
			Verbs: []string{
				"get",
				"list",
				"watch",
			},
		},
	}
}

func policyRuleForServer() []v1.PolicyRule {
	return []v1.PolicyRule{
		{
			APIGroups: []string{
				"*",
			},
			Resources: []string{
				"*",
			},
			Verbs: []string{
				"get",
				"patch",
				"delete",
			},
		},
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"secrets",
				"configmaps",
			},
			Verbs: []string{
				"create",
				"get",
				"list",
				"watch",
				"update",
				"patch",
				"delete",
			},
		}, {
			APIGroups: []string{
				"argoproj.io",
			},
			Resources: []string{
				"applications",
				"appprojects",
			},
			Verbs: []string{
				"create",
				"get",
				"list",
				"watch",
				"update",
				"delete",
				"patch",
			},
		},
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"events",
			},
			Verbs: []string{
				"create",
				"list",
			},
		},
	}
}

func policyRuleForNotificationsController() []v1.PolicyRule {
	return []v1.PolicyRule{

		{
			APIGroups: []string{
				"argoproj.io",
			},
			Resources: []string{
				"applications",
				"appprojects",
			},
			Verbs: []string{
				"get",
				"list",
				"patch",
				"update",
				"watch",
			},
		},
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"configmaps",
				"secrets",
			},
			Verbs: []string{
				"list",
				"watch",
			},
		},
		{
			APIGroups: []string{
				"",
			},
			ResourceNames: []string{
				"argocd-notifications-cm",
			},
			Resources: []string{
				"configmaps",
			},
			Verbs: []string{
				"get",
			},
		},
		{
			APIGroups: []string{
				"",
			},
			ResourceNames: []string{
				"argocd-notifications-secret",
			},
			Resources: []string{
				"secrets",
			},
			Verbs: []string{
				"get",
			},
		},
	}
}

func policyRuleForServerApplicationSourceNamespaces() []v1.PolicyRule {
	return []v1.PolicyRule{
		{
			APIGroups: []string{
				"argoproj.io",
			},
			Resources: []string{
				"applications",
			},
			Verbs: []string{
				"create",
				"get",
				"list",
				"patch",
				"update",
				"watch",
				"delete",
			},
		},
	}
}

func policyRuleForServerClusterRole() []v1.PolicyRule {
	return []v1.PolicyRule{
		{
			APIGroups: []string{
				"*",
			},
			Resources: []string{
				"*",
			},
			Verbs: []string{
				"get",
				"delete",
				"patch",
			},
		},
		{
			APIGroups: []string{
				"argoproj.io",
			},
			Resources: []string{
				"applications",
			},
			Verbs: []string{
				"list",
				"watch",
			},
		},
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"events",
			},
			Verbs: []string{
				"list",
			},
		},
	}
}

func policyRuleForGrafana(client client.Client) []v1.PolicyRule {
	rules := []v1.PolicyRule{}

	// Need additional policy rules if we are running on openshift, else the stateful set won't have the right
	// permissions to start
	rules = appendOpenShiftNonRootSCC(rules, client)

	return rules
}

func getPolicyRuleList(client client.Client) []struct {
	name       string
	policyRule []v1.PolicyRule
} {
	return []struct {
		name       string
		policyRule []v1.PolicyRule
	}{
		{
			name:       common.ArgoCDApplicationControllerComponent,
			policyRule: policyRuleForApplicationController(),
		}, {
			name:       common.ArgoCDDexServerComponent,
			policyRule: policyRuleForDexServer(),
		}, {
			name:       common.ArgoCDServerComponent,
			policyRule: policyRuleForServer(),
		}, {
			name:       common.ArgoCDRedisHAComponent,
			policyRule: policyRuleForRedisHa(client),
		}, {
			name:       common.ArgoCDRedisComponent,
			policyRule: policyRuleForRedis(client),
		}, {
			name:       common.ArgoCDOperatorGrafanaComponent,
			policyRule: policyRuleForGrafana(client),
		},
	}
}

func getPolicyRuleClusterRoleList() []struct {
	name       string
	policyRule []v1.PolicyRule
} {
	return []struct {
		name       string
		policyRule []v1.PolicyRule
	}{
		{
			name:       common.ArgoCDApplicationControllerComponent,
			policyRule: policyRuleForApplicationController(),
		}, {
			name:       common.ArgoCDServerComponent,
			policyRule: policyRuleForServerClusterRole(),
		},
	}
}

func appendOpenShiftNonRootSCC(rules []v1.PolicyRule, client client.Client) []v1.PolicyRule {
	if IsVersionAPIAvailable() {
		// Starting with OpenShift 4.11, we need to use the resource name "nonroot-v2" instead of "nonroot"
		resourceName := "nonroot"
		version, err := getClusterVersion(client)
		if err != nil {
			log.Error(err, "couldn't get OpenShift version")
		}
		if version == "" || semver.Compare(fmt.Sprintf("v%s", version), "v4.10.999") > 0 {
			resourceName = "nonroot-v2"
		}
		orules := v1.PolicyRule{
			APIGroups: []string{
				"security.openshift.io",
			},
			ResourceNames: []string{
				resourceName,
			},
			Resources: []string{
				"securitycontextconstraints",
			},
			Verbs: []string{
				"use",
			},
		}
		rules = append(rules, orules)
	}
	return rules
}
