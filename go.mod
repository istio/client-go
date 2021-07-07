module istio.io/client-go

go 1.16

require (
	istio.io/api v0.0.0-20210702170716-527b9df3805f
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
)

// FIXME: this is just here until https://github.com/istio/api/pull/1940 gets merged
replace istio.io/api => github.com/dgn/api v0.0.0-20210707145534-e8371b54d539
