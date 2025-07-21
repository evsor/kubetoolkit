package k8s

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getClientset() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func CreateDeployment(repo, name, namespace string) error {
	clientset, err := getClientset()
	if err != nil {
		return fmt.Errorf("failed to get k8s client: %w", err)
	}
	image := repo + "/" + name + ":latest"
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name + "-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:    name,
						Image:   image,
						Command: []string{"tail", "-f", "/dev/null"},
					}},
				},
			},
		},
	}
	_, err = clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	return err
}

func DeleteDeployment(name, namespace string) error {
	clientset, err := getClientset()
	if err != nil {
		return fmt.Errorf("failed to get k8s client: %w", err)
	}
	deletePolicy := metav1.DeletePropagationForeground
	return clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name+"-deployment", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

func int32Ptr(i int32) *int32 { return &i }
