package docker

import (
	"strings"

	"github.com/crqra/whatami/adapter"
	"github.com/crqra/whatami/file"
)

const (
	docker        = "docker"
	dockerfile    = "Dockerfile"
	dockerCompose = "docker-compose"
)

type DockerAdapter struct{}

func (a DockerAdapter) FindLanguages(*file.File) ([]*adapter.Language, error) {
	return nil, nil
}

func (a DockerAdapter) FindTools(f *file.File) ([]*adapter.Tool, error) {
	tools := []*adapter.Tool{}

	if strings.HasPrefix(f.Name(), dockerfile) {
		tools = append(tools, &adapter.Tool{Name: docker})
	}

	if strings.Replace(f.Name(), f.Ext(), "", 1) == dockerCompose {
		tools = append(tools, &adapter.Tool{Name: dockerCompose})
	}

	return tools, nil
}

func (a DockerAdapter) FindDependencies(*file.File) ([]*adapter.Dependency, error) {
	return nil, nil
}
