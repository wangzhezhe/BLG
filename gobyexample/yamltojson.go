package main

import (
	"fmt"
	//"github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1"
	//"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util/yaml"

	//"gopkg.in/yaml.v2"
	//"log"
)

var data = `
apiVersion: v1
kind: ReplicationController
metadata:
  name: kube-dns-v8
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    version: v8
    kubernetes.io/cluster-service: "true"
spec:
  replicas: 1
  selector:
    k8s-app: kube-dns
    version: v8
  template:
    metadata:
      labels:
        k8s-app: kube-dns
        version: v8
        kubernetes.io/cluster-service: "true"
    spec:
      containers:
      - name: etcd
        image: gcr.io/google_containers/etcd:2.0.9
        resources:
          limits:
            cpu: 100m
            memory: 50Mi
        command:
        - /usr/local/bin/etcd
        - -data-dir
        - /var/etcd/data
        - -listen-client-urls
        - http://127.0.0.1:2379,http://127.0.0.1:4001
        - -advertise-client-urls
        - http://127.0.0.1:2379,http://127.0.0.1:4001
        - -initial-cluster-token
        - skydns-etcd
        volumeMounts:
        - name: etcd-storage
          mountPath: /var/etcd/data
      - name: kube2sky
        image: gcr.io/google_containers/kube2sky:1.11
        resources:
          limits:
            cpu: 100m
            memory: 50Mi
        args:
        # command = "/kube2sky"
        - -kube_master_url=http://121.40.171.96:8080
        - -domain=cluster.local
      - name: skydns
        image: gcr.io/google_containers/skydns:2015-03-11-001
        resources:
          limits:
            cpu: 100m
            memory: 50Mi
        args:
        # command = "/skydns"
        - -machines=http://localhost:4001
        - -addr=0.0.0.0:53
        - -domain=cluster.local.
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 30
          timeoutSeconds: 5
      - name: healthz
        image: gcr.io/google_containers/exechealthz:1.0
        resources:
          limits:
            cpu: 10m
            memory: 20Mi
        args:
        - -cmd=nslookup kubernetes.default.svc.cluster.local localhost >/dev/null
        - -port=8080
        ports:
        - containerPort: 8080
          protocol: TCP
      volumes:
      - name: etcd-storage
        emptyDir: {}
      dnsPolicy: Default  # Don't use cluster DNS.

`

func main() {

	//Codec := v1.Codec
	//	codec := runtime.YAMLDecoder(Codec)
	//内部具体使用的是这个包 https://github.com/ghodss/yaml
	//感觉还是比较好用的
	out, err := yaml.ToJSON([]byte(data))
	if err != nil {
		panic(err)
	}

	//resp, err := codec.Decode()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", string(out))

}
