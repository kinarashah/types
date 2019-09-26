module github.com/rancher/types

go 1.12

replace (
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	github.com/rancher/norman => github.com/kinarashah/norman v0.0.0-20190925222837-edac411b7652
	github.com/rancher/wrangler => github.com/kinarashah/wrangler v0.2.1-0.20190925215808-0f46c68e1e13
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)

require (
	github.com/coreos/prometheus-operator v0.33.0
	github.com/knative/pkg v0.0.0-20190817231834-12ee58e32cc8
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rancher/norman v0.0.0-20190821234528-20a936b685b0
	github.com/sirupsen/logrus v1.4.2
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v2.0.0-alpha.0.0.20181121191925-a47917edff34+incompatible
	k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a
	k8s.io/kube-aggregator v0.0.0-20190805183716-8439689952da
)
