# whatami

> What Am I? - point to a project and see what tools, languages and dependencies
> it has.

## About

### What?

Tool that recursively walks a directory and based on the files present derives:

- Tools: package managers, compilers, linters, etc.
- Languages: markup, configuration and programming languages
- Dependencies: software dependencies from package managers configuration files

For each of the above a version _may_ be available in which case it's included.

### Why?

My main use case is to use the information outputted to create dynamic CI/CD
pipelines. Based on the output, I know what tools I need to have in place for a
particular project and can even infer what commands I can run to perform certain
tasks like building or testing the project.

### How?

For each file or directory found while walking, a set of pre-defined _adapters_
is executed to derive tools, languages and dependencies.

If more than one adapter finds the same tools, languages or dependencies, they're
de-duplicated based on the version. Wins the first tool, language or dependency
that for which a version was found (regardless if subsequent findings also have
a version).

## Installing

```console
$ go install github.com/crqra/whatami/cmd/whatami@latest
```

## Usage

```
$ whatami -h
usage: whatami [-h] [-i=<PATH> ...] [directory]
```

### Flags

| Flag        | Description                           |
| ----------- | ------------------------------------- |
| `-h`        | Show usage                            |
| `-i <PATH>` | Path patterns to ignore. Accepts many |

## Example

In a standard TypeScript project:

```bash
$ ls
index.ts node_modules/ package-lock.json package.json

# Run whatami
$ whatami -i node_modules
```

<details>
<summary>See output</summary>

```json
{
  "tools": {
    "node": {
      "version": "~15"
    },
    "npm": {
      "version": "~1.0.20"
    },
    "tsc": {
      "version": "^4.7.3"
    }
  },
  "dependencies": {
    "express": {
      "version": "^4.18.1",
      "type": "production"
    },
    "typescript": {
      "version": "^4.7.3",
      "type": "development"
    }
  },
  "languages": {
    "typescript": {
      "version": "^4.7.3"
    }
  }
}
```

</details>

## Adapters

An _adapter_ is an interface that implements the functionality to derive one or
more tools, languages or dependencies based on a given _file_.

```go
type Adapter interface {
	FindTools(f *file.File) ([]*Tool, error)
	FindDependencies(f *file.File) ([]*Dependency, error)
	FindLanguages(f *file.File) ([]*Language, error)
}
```

See the links in the list of supported adapters below for example implementations.

### Supported

- [Docker][docker-src]
  - Tools: `docker`
- [Go][go-src]
  - Tools: `go`
  - Languages: `go`
- [HCL][hcl-src]
  - Languages: `hcl`
- [Java][java-src]
  - Languages: `java`
- [JavaScript][javascript-src]
  - Languages: `javascript`
- [Maven][maven-src]
  - Tools: `mvn`
  - Languages: `java`
- [NPM][npm-src]
  - Tools: `node`, `npm`
- [Taskfile][taskfile-src]
  - Tools: `task`
- [Terraform][terraform-src]
  - Tools: `terraform`
  - Languages: `hcl`
- [TypeScript][typescript-src]
  - Tools: `tsc`
  - Languages: `typescript`
- [Yarn][yarn-src]
  - Tools: `yarn`

If the adapter you need is not on the list above, please open an [issue][issues]
or a [pull request][pulls]. Use the [adapter template][adapter-tpl] to get started.

## License

This project is released under the [MIT License](LICENSE).

[issues]: https://github.com/crqra/whatami/issues
[pulls]: https://github.com/crqra/whatami/pulls
[adapter-tpl]: https://github.com/crqra/whatami/tree/main/adapter/_template
[docker-src]: https://github.com/crqra/whatami/blob/main/adapter/docker/docker.go
[go-src]: https://github.com/crqra/whatami/blob/main/adapter/golang/golang.go
[hcl-src]: https://github.com/crqra/whatami/blob/main/adapter/hcl/hcl.go
[java-src]: https://github.com/crqra/whatami/blob/main/adapter/java/java.go
[javascript-src]: https://github.com/crqra/whatami/blob/main/adapter/javascript/javascript.go
[maven-src]: https://github.com/crqra/whatami/blob/main/adapter/maven/maven.go
[npm-src]: https://github.com/crqra/whatami/blob/main/adapter/npm/npm.go
[taskfile-src]: https://github.com/crqra/whatami/blob/main/adapter/taskfile/taskfile.go
[terraform-src]: https://github.com/crqra/whatami/blob/main/adapter/terraform/terraform.go
[typescript-src]: https://github.com/crqra/whatami/blob/main/adapter/typescript/typescript.go
[yarn-src]: https://github.com/crqra/whatami/blob/main/adapter/yarn/yarn.go
