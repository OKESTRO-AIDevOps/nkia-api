package kuberesource

type HorizontalPodAutoscaler struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ScaleTargetRef struct {
			APIVersion string `yaml:"apiVersion"`
			Kind       string `yaml:"kind"`
			Name       string `yaml:"name"`
		} `yaml:"scaleTargetRef"`
		MinReplicas                    int `yaml:"minReplicas"`
		MaxReplicas                    int `yaml:"maxReplicas"`
		TargetCPUUtilizationPercentage int `yaml:"targetCPUUtilizationPercentage"`
	} `yaml:"spec"`
}

type Ingress struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Annotations struct {
			NginxIngressKubernetesIoProxyBodySize string `yaml:"nginx.ingress.kubernetes.io/proxy-body-size"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec struct {
		Rules []struct {
			Host string `yaml:"host"`
			HTTP struct {
				Paths []struct {
					Path     string `yaml:"path"`
					PathType string `yaml:"pathType"`
					Backend  struct {
						Service struct {
							Name string `yaml:"name"`
							Port struct {
								Number string `yaml:"number"`
							} `yaml:"port"`
						} `yaml:"service"`
					} `yaml:"backend"`
				} `yaml:"paths"`
			} `yaml:"http"`
		} `yaml:"rules"`
	} `yaml:"spec"`
}

type Service struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			App string `yaml:"app"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Type  string `yaml:"type"`
		Ports []struct {
			Port       int    `yaml:"port"`
			TargetPort int    `yaml:"targetPort"`
			Protocol   string `yaml:"protocol"`
		} `yaml:"ports"`
		Selector struct {
			App string `yaml:"app"`
		} `yaml:"selector"`
	} `yaml:"spec"`
}

type Deployment struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Replicas int `yaml:"replicas"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				ImagePullSecrets []struct {
					Name string `yaml:"name"`
				} `yaml:"imagePullSecrets"`
				Containers []struct {
					Name            string `yaml:"name"`
					Image           string `yaml:"image"`
					ImagePullPolicy string `yaml:"imagePullPolicy"`
					Ports           []struct {
						ContainerPort int `yaml:"containerPort"`
					} `yaml:"ports"`
				} `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}
