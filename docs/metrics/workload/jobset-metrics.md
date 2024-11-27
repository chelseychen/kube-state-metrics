# Job Metrics

| Metric name                           | Metric type | Description                | Labels/tags                                                                                                                                                                                                 | Status       |
| ------------------------------------- | ----------- | ------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------ |
| kube_jobset_specified_replicas        | Gauge       |                            | `jobset_name`=&lt;jobset-name&gt; <br> `namespace`=&lt;jobset-namespace&gt; <br> `replicated_job_name`=&lt;replicated-job-name&gt;                                   | EXPERIMENTAL |
| kube_jobset_status_replicas           | Gauge       |                            | `jobset_name`=&lt;jobset-name&gt; <br> `namespace`=&lt;jobset-namespace&gt; <br> `replicated_job_name`=&lt;replicated-job-name&gt; <br> `status`=&lt;ready\|succeeded\|failed\|active\|suspended&gt;             | EXPERIMENTAL |
| kube_jobset_status_condition          | Gauge       |                            | `jobset_name`=&lt;jobset-name&gt; <br> `namespace`=&lt;jobset-namespace&gt; <br> `condition`=&lt;deployment-condition&gt; <br> `status`=&lt;true\|false\|unknown&gt; | EXPERIMENTAL |
| kube_jobset_annotations               | Gauge       |Kubernetes annotations converted to Prometheus labels controlled via [--metric-annotations-allowlist](../../developer/cli-arguments.md) | `jobset_name`=&lt;jobset-name&gt; <br> `namespace`=&lt;jobset-namespace&gt; <br> `annotation_JOBSET_ANNOTATION`=&lt;JOBSET_ANNOTATION&gt;                        | EXPERIMENTAL |
| kube_jobset_labels                    | Gauge       |Kubernetes labels converted to Prometheus labels controlled via [--metric-labels-allowlist](../../developer/cli-arguments.md)           | `jobset_name`=&lt;jobset-name&gt; <br> `namespace`=&lt;jobset-namespace&gt; <br> `label_JOBSET_LABEL`=&lt;JOBSET_LABEL&gt;                                       | EXPERIMENTAL |