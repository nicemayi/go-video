kubernetes

master node:
    * API server
    * scheduler
    * controller
    * etcd

slave node:
    * pod, share name spaces
    * Docker (image runtime)
    * kubelet (agent, accpect commands from master)
    * kube-proxy
    * fluentd (logging)
    * DNS, UI (optional)

k8s install
1. minikube local, single node (master & worker), 学习用的
2. kubeadm 搭建真正的集群

windows minikube + virtualbox, kubectl

minikube start
minikube status
minikube stop
minikube delete
minikube dashboard

minikube ssh

kubectl version
kubectl get pod -A



install docker, minikube, kubectl

minikube start --vm-driver=virtualbox
minikube status
minikube stop
minikube delete
```
➜  ~ minikube start --vm-driver=virtualbox
😄  minikube v1.9.0 on Darwin 10.14.1
💥  The existing "minikube" VM was created using the "hyperkit" driver, and is incompatible with the "virtualbox" driver.
👉  To proceed, either:

1) Delete the existing "minikube" cluster using: 'minikube delete'

* or *

2) Start the existing "minikube" cluster using: 'minikube start --driver=hyperkit'

💣  Exiting.
➜  ~ minikube delete
🔥  Deleting "minikube" in hyperkit ...
💀  Removed all traces of the "minikube" cluster.
➜  ~ minikube start --vm-driver=virtualbox
😄  minikube v1.9.0 on Darwin 10.14.1
✨  Using the virtualbox driver based on user configuration
🔥  Creating virtualbox VM (CPUs=2, Memory=4000MB, Disk=20000MB) ...
🐳  Preparing Kubernetes v1.18.0 on Docker 19.03.8 ...
🌟  Enabling addons: default-storageclass, storage-provisioner
🏄  Done! kubectl is now configured to use "minikube"
➜  ~ minikube status
m01
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
```

minikube ssh
docker ps

/////// kubectl

kubectl get pods -A
```
➜  ~ kubectl get pods -A

kubectl get pods --all-namespaces

NAMESPACE     NAME                               READY   STATUS    RESTARTS   AGE
kube-system   coredns-66bff467f8-kkx7d           1/1     Running   0          2m4s
kube-system   coredns-66bff467f8-nzts7           1/1     Running   0          2m4s
kube-system   etcd-minikube                      1/1     Running   0          2m10s
kube-system   kube-apiserver-minikube            1/1     Running   0          2m10s
kube-system   kube-controller-manager-minikube   1/1     Running   0          2m10s
kube-system   kube-proxy-5nhl7                   1/1     Running   0          2m4s
kube-system   kube-scheduler-minikube            1/1     Running   0          2m10s
kube-system   storage-provisioner                1/1     Running   0          2m10s
```

➜  ~ kubectl version --output=yaml
clientVersion:
  buildDate: "2020-03-25T14:58:59Z"
  compiler: gc
  gitCommit: 9e991415386e4cf155a24b1da15becaa390438d8
  gitTreeState: clean
  gitVersion: v1.18.0
  goVersion: go1.13.8
  major: "1"
  minor: "18"
  platform: darwin/amd64
serverVersion:
  buildDate: "2020-03-25T14:50:46Z"
  compiler: gc
  gitCommit: 9e991415386e4cf155a24b1da15becaa390438d8
  gitTreeState: clean
  gitVersion: v1.18.0
  goVersion: go1.13.8
  major: "1"
  minor: "18"
  platform: linux/amd64


minikube dashboard


// linux

all nodes need to have docker, kubeadm, kubelet, kubectl installed

sudo kubeadm init --pod-network-cidr 172.100.0.0/16 --apiserver-advertise-address 157.230.169.141

sudo kubeadm init --pod-network-cidr 172.100.0.0/16 --apiserver-advertise-address 157.230.169.141

```
apt-get update && apt-get install -y \
  apt-transport-https ca-certificates curl software-properties-common gnupg2

### Add Docker’s official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

### Add Docker apt repository.
add-apt-repository \
  "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) \
  stable"

## Install Docker CE.
apt-get update && apt-get install -y \
  containerd.io=1.2.13-1 \
  docker-ce=5:19.03.8~3-0~ubuntu-$(lsb_release -cs) \
  docker-ce-cli=5:19.03.8~3-0~ubuntu-$(lsb_release -cs)

# Setup daemon.
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

mkdir -p /etc/systemd/system/docker.service.d

# Restart docker.
systemctl daemon-reload
systemctl restart docker
```

master
sudo kubeadm init --pod-network-cidr 10.244.0.0/16 --apiserver-advertise-address 157.230.169.141

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config -y
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl get pod --all-namespaces

flannel network
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

kubectl get pod --all-namespaces

client
sudo kubeadm join 157.230.169.141:6443 --token z81lby.c1ahm646ajeq5nyi \
    --discovery-token-ca-cert-hash sha256:fbe3eec8dc73c34a4e4cb3b8a0fbf27a813519e71f1c2e39ea4b81ac7c712dc6



apt-mark unhold kubeadm && \
apt-get update && apt-get install -y kubeadm=1.17.0 && \
apt-mark hold kubeadm


curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add - && \
  echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list && \
  sudo apt-get update -q && \
  sudo apt-get install -qy kubelet=1.17.0-00 kubectl=1.9.0-00 kubeadm=1.17.0-00


////// chapter 11
本地可以切换不同的context
kubectl config current-context
kubectl config get-contexts
kubectl config use-context kubeadm

kubectl get node
kubectl get node node3
kubectl describe node node3

kubectl get node -o wide
kubectl get node -o yaml
kubectl get node -o json

kubectl get node --show-labels

// 增加label
kubectl label node node1 env=test
// 删除一个label
kubectl label node node1 env-

role是一种特殊的label
kubectl label node node3 node-role.kubernetes.io/worker=
kubectl label node ubuntu-s-1vcpu-2gb-sfo2-02 node-role.kubernetes.io/worker=

kubectl create -f ./pod/k.yml
kubectl get pods
kubectl describe pod nginx-busybox
kubectl get pods nginx-busybox -o wide
// 默认进入第一个shell
kubectl exec nginx-busybox -it sh

kubectl delete -f ./pod/k.yml

kubectl exec nginx-busybox date

kubectl exec nginx-busybox -c busybox date
kubectl exec nginx-busybox -it -c busybox sh

kubectl delete -f ./pod/k.yml

kubectl get namespace

kubectl create namespace demo
// namespace可以用来过滤

kubectl create -f ./pod/nginx.yml
kubectl create -f ./pod/nginx_demo.yml

kubectl get pod --namespace demo // 用namespace来过滤
kubectl get pod --all-namespaces

// Context: 可以将default的namespace改成其他的

kubectl config get-contexts

kubectl config use-context minikube
kubectl config delete-context demo

controller是不停的更新当前状态(去deployment)，使之状态尽可能的等于预定状态

kubectl get deployment

// create是从无到有
kubectl create -f 
// apply是更新镜像或者修改relicas的数目
kubectl apply -f

kubectl edit deployment nginx-deployment.yml

kubectl scale --current-replicas=2 --replicas=3 deployment/mysql

kubectl set image nginx_deployment nginx=nginx:1.9.1
kubectl get replicaset
kubectl scale --current-replicas=4 --replicas=6 deployment/nginx-deployment-test

kubectl rollout status deployment nginx-deployment-test

kubectl rollout history
kubectl rollout history deployment nginx-deployment-test --revision 2
kubectl rollout undo deployment nginx-deployment-test --to-revision 2

// 容器监控