Install dep
Install kustomize
Install kubebuilder

#创建CRD
$ kubebuilder init --domain k8s.io --license apache2 --owner "The Kubernetes Authors"


#创建API
$ kubebuilder create api --group mygroup --version v1beta1 --kind MyKind


# 安装 CRDs into the cluster
$ make install

# Build and run the manager
$ make run

# 创建CRD测试样本
$ kubectl apply -f sample/<resource>.yaml

---
# Create a docker image
$ make docker-build IMG=<img-name>

# Push the docker image to a configured container registry
$ make docker-push IMG=<img-name>

# Deploy the controller manager manifests to the cluster.
$ make deploy
