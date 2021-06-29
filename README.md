# Go Project Template

## Strings to replace

github.com/freeo/`go-project-template`
github.com/freeo/go-project-template/`submodule`

## Project structure

```
./
	/cmd:
		- main.go
		- go.mod (requires submodule)
	/submodule:
		- go.mod
	Dockerfile
	README.md
```

TODO: go.mod in root.folder necessary? VSCode support for multiple-module
projects is in early stage and therefore currently not best practice. A common
root go.mod was found in other projects.

## Dockerfile

Includes multi-stage build with Alpine golang builder and Alpine runtime image.
