# Production Environment Setup Instructions [(reference)](https://kubernetes.io/docs/setup/production-environment/)
1. Install any [container-runtime](https://kubernetes.io/docs/setup/production-environment/container-runtimes/) like `Docker`, `containerd`.
2. Install Kubernetes deployment tools.
    - Install any one of the [bootstrapping-clusters](https://kubernetes.io/docs/setup/production-environment/tools/) like `kubeadm`, `kops` or `kubespray`.
    - Install `kubelet` and `kubectl`. Note the kubelet version may never exceed the API server version.
