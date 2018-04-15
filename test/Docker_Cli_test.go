package test

import (
	"testing"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"github.com/docker/docker/client"
	"fmt"
	"os"
	"../services/docker-cli"
)

func Test_buildImgae_service(t *testing.T) {
	var (
		contextPath = "/tmp/uploadImage/993CA2B6A5934991A42225D47EA7F4F6"
		imageFullName = "127.0.0.1/ascs/digital-summit"
		version = "1.1.0"
	)
	res, err := docker_cli.BuildImage(contextPath, imageFullName, version)
	if nil != err {
		t.Error(err)
		return
	}
	t.Log(res)
}

func tTest_buildImage(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"), client.WithVersion("1.36"))
	if nil != err {
		t.Error(err)
	}
	//tarPath := "/tmp/timeandspace-platform-2.0.10.tar.gz"
	var (
		contextPath = "/tmp/uploadImage/993CA2B6A5934991A42225D47EA7F4F6"
		imageFullName = "127.0.0.1/ascs/digital-summit"
		version = "1.1.0"
	)
	contextReader, err := os.Open(contextPath)
	defer contextReader.Close()
	if nil != err {
		t.Error(err)
	}
	imageBuildResponse, err := cli.ImageBuild(context.Background(), contextReader, types.ImageBuildOptions{
		Tags: []string{imageFullName + ":" + version},
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
	resBuff := make([]byte, 1024)
	imageBuildResponse.Body.Read(resBuff)
	defer imageBuildResponse.Body.Close()
	fmt.Println(string(resBuff))
}