package test

import (
	"testing"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"github.com/docker/docker/client"
	"fmt"
	"os"
	//"io/ioutil"
)

func Test_buildImage(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"), client.WithVersion("1.36"))
	if nil != err {
		t.Error(err)
	}
	//tarPath := "/tmp/timeandspace-platform-2.0.10.tar.gz"
	tarPath := "/tmp/demo.tar.gz"
	tar, err := os.Open(tarPath)
	defer tar.Close()
	if nil != err {
		t.Error(err)
	}
	imageBuildResponse, err := cli.ImageBuild(context.Background(), tar, types.ImageBuildOptions{
		Tags: []string{"127.0.0.1/ascs/timeandspace-platform:2.0.10"},
		NoCache: true,
		SuppressOutput: true,
		Remove: true,
		ForceRemove: true,
		//PullParent: true,
		//Target: "/tmp/uploadImage/info",
	})
	if nil != err {
		t.Error(err)
	}
	fmt.Fprint(os.Stdout, imageBuildResponse.Body)
	defer imageBuildResponse.Body.Close()
}