package cloud

const (
	GKE        = "gke"
	OKE        = "oke"
	EKS        = "eks"
	AKS        = "aks"
	AWS        = "aws"
	PKS        = "pks"
	IKS        = "iks"
	MINIKUBE   = "minikube"
	MINISHIFT  = "minishift"
	KUBERNETES = "kubernetes"
	OPENSHIFT  = "openshift"
	ORACLE     = "oracle"
	ICP        = "icp"
	JX_INFRA   = "jx-infra"
)

// KubernetesProviders list of all available Kubernetes providers
var KubernetesProviders = []string{MINIKUBE, GKE, OKE, AKS, AWS, EKS, KUBERNETES, IKS, OPENSHIFT, MINISHIFT, JX_INFRA, PKS, ICP}
