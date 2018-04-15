package docker_cli

import (
	"github.com/docker/docker/api/types"
	"fmt"
	"os"
)

/**
	Build an image from a tar archive with a Dockerfile in it.
 */
func BuildImage(contextPath, imageName, version string) (interface{}, error){
	cli, err := getCli()
	if nil != err {
		return nil, err
	}
	contextReader, err := os.Open(contextPath)
	defer contextReader.Close()
	if nil != err {
		return nil, err
	}
	imageBuildResponse, err := cli.ImageBuild(getContext(), contextReader, types.ImageBuildOptions{
		Tags: []string{imageName+":"+version},
		NoCache: true,
		SuppressOutput: true,
		Remove: true,
		ForceRemove: true,
	})
	if nil != err {
		return nil, err
	}
	resBuff := make([]byte, 1024)
	imageBuildResponse.Body.Read(resBuff)
	defer imageBuildResponse.Body.Close()
	return string(resBuff), nil
}

func ListImage() ([]types.ImageSummary, error) {
	cli, err := getCli()
	if nil != err {
		return nil, err
	}
	images, err := cli.ImageList(getContext(), types.ImageListOptions{})
	if nil != err {
		return nil, err
	}
	return images, nil
}

func InspectImage(imageID string) (interface{}, error) {
	cli, err := getCli()
	if nil != err {
		return nil, err
	}
	imageType, image, err := cli.ImageInspectWithRaw(getContext(), imageID)
	if nil != err {
		return nil, err
	}
	fmt.Println(string(image))
	return imageType, nil
}
