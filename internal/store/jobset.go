/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package store

import (
	"context"
	"strconv"

	basemetrics "k8s.io/component-base/metrics"

	"k8s.io/kube-state-metrics/v2/pkg/metric"
	generator "k8s.io/kube-state-metrics/v2/pkg/metric_generator"

	jobsetv1alpha2 "sigs.k8s.io/jobset/api/jobset/v1alpha2"
	jobsetclientset "sigs.k8s.io/jobset/client-go/clientset/versioned"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	descJobSetAnnotationsName     = "kube_jobset_annotations"
	descJobSetAnnotationsHelp     = "Kubernetes annotations converted to Prometheus labels."
	descJobSetLabelsName          = "kube_jobset_labels"
	descJobSetLabelsHelp          = "Kubernetes labels converted to Prometheus labels."
	descJobSetLabelsDefaultLabels = []string{"namespace", "jobset_name"}
)

func jobSetMetricFamilies(allowAnnotationsList, allowLabelsList []string) []generator.FamilyGenerator {
	return []generator.FamilyGenerator{
		*generator.NewFamilyGeneratorWithStability(
			descJobSetAnnotationsName,
			descJobSetAnnotationsHelp,
			metric.Gauge,
			basemetrics.ALPHA,
			"",
			wrapJobFunc(func(js *jobsetv1alpha2.JobSet) *metric.Family {
				if len(allowAnnotationsList) == 0 {
					return &metric.Family{}
				}
				annotationKeys, annotationValues := createPrometheusLabelKeysValues("annotation", js.Annotations, allowAnnotationsList)
				return &metric.Family{
					Metrics: []*metric.Metric{
						{
							LabelKeys:   annotationKeys,
							LabelValues: annotationValues,
							Value:       1,
						},
					},
				}
			}),
		),
		*generator.NewFamilyGeneratorWithStability(
			descJobSetLabelsName,
			descJobSetLabelsHelp,
			metric.Gauge,
			basemetrics.ALPHA,
			"",
			wrapJobFunc(func(js *jobsetv1alpha2.JobSet) *metric.Family {
				if len(allowLabelsList) == 0 {
					return &metric.Family{}
				}
				labelKeys, labelValues := createPrometheusLabelKeysValues("label", js.Labels, allowLabelsList)
				return &metric.Family{
					Metrics: []*metric.Metric{
						{
							LabelKeys:   labelKeys,
							LabelValues: labelValues,
							Value:       1,
						},
					},
				}
			}),
		),
	}
}

func wrapJobSetFunc(f func(*jobsetv1alpha2.JobSet) *metric.Family) func(interface{}) *metric.Family {
	return func(obj interface{}) *metric.Family {
		jobset := obj.(*jobsetv1alpha2.JobSet)

		metricFamily := f(jobset)

		for _, m := range metricFamily.Metrics {
			m.LabelKeys, m.LabelValues = mergeKeyValues(descJobSetLabelsDefaultLabels, []string{jobset.Namespace, jobset.Name}, m.LabelKeys, m.LabelValues)
		}

		return metricFamily
	}
}

func createJobSetListWatch(kubeClient jobsetclientset.Interface, ns string, fieldSelector string) cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			opts.FieldSelector = fieldSelector
			return kubeClient.JobsetV1alpha2().JobSets(ns).List(context.TODO(), opts)
		},
		WatchFunc: func(opts metav1.ListOptions) (watch.Interface, error) {
			opts.FieldSelector = fieldSelector
			return kubeClient.JobsetV1alpha2().JobSets(ns).Watch(context.TODO(), opts)
		},
	}
}
