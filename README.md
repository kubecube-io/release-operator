# Release-Operator

release-operator do customize operations for specified project

## Config

```json
[
  {
    "name": "KubeCube",
    "repo": "https://github.com/kubecube-io/KubeCube.git",
    "branch": "v1.0.2",
    "exec": "make build"
  },
  {
    "name": "kubecube-audit",
    "repo": "https://github.com/kubecube-io/kubecube-audit.git",
    "branch": "v1.0.2",
    "exec": "make build"
  },
  {
    "name": "kubecube-webconsole",
    "repo": "https://github.com/kubecube-io/kubecube-webconsole.git",
    "branch": "v1.0.2",
    "exec": "make build"
  },
  {
    "name": "kubecube-front",
    "repo": "https://github.com/kubecube-io/kubecube-front.git",
    "branch": "v1.0.2",
    "exec": "make build"
  }
]
```

## Usage

execute release operate

```bash
make run
```

clean local repos

```bash
make clean up
```