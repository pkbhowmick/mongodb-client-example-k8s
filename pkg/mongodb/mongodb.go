package mongodb

import (
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Connect() error {
	ns := metav1.NamespaceDefault
	svc := "mongodb-dev.default.svc.cluster.local:27017"

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "curl",
							Image: "curlimages/curl",
						},
					},
				},
			},
		},
	}
}
