package main

import (
	"github.com/sirupsen/logrus"
	veleroplugin "github.com/vmware-tanzu/velero/pkg/plugin/framework"
)

func main() {
	veleroplugin.NewServer().
		RegisterObjectStore("community.openstack.org/openstack", newSwiftObjectStore).
		RegisterVolumeSnapshotter("community.openstack.org/openstack", newCinderBlockStore).
		Serve()
}

func newSwiftObjectStore(logger logrus.FieldLogger) (interface{}, error) {
	return NewObjectStore(logger), nil
}

func newCinderBlockStore(logger logrus.FieldLogger) (interface{}, error) {
	return NewBlockStore(logger), nil
}
