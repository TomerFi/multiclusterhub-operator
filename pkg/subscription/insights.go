// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package subscription

import (
	operatorsv1 "github.com/stolostron/multiclusterhub-operator/api/v1"
	"github.com/stolostron/multiclusterhub-operator/pkg/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Insights overrides the insights chart
func Insights(m *operatorsv1.MultiClusterHub, overrides map[string]string) *unstructured.Unstructured {
	sub := &Subscription{
		Name:      "policyreport",
		Namespace: m.Namespace,
		Overrides: map[string]interface{}{
			"hubconfig": map[string]interface{}{
				"name":         m.Name,
				"namespace":    m.Namespace,
				"replicaCount": utils.DefaultReplicaCount(m),
				"nodeSelector": m.Spec.NodeSelector,
				"tolerations":  utils.GetTolerations(m),
			},
			"global": map[string]interface{}{
				"imageOverrides": overrides,
				"pullPolicy":     utils.GetImagePullPolicy(m),
				"pullSecret":     m.Spec.ImagePullSecret,
			},
			"clusterImageSets": map[string]interface{}{
				"subscriptionPause": utils.GetDisableClusterImageSets(m),
			},
		},
	}
	setCustomCA(m, sub)

	return newSubscription(m, sub)
}
