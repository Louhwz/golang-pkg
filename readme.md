# Cloud Native Golang SDK

This is a simple golang sdk which mainly focus on consul, kubernetes and strings.

## Install

```
go get github.com/louhwz/pkg
```

## Usage
For example, to quickly get a configMap from kubernetes:
```
import 	"github.com/louhwz/pkg/loukubernetes"

func main() {
    clientset, err := loukubernetes.InClusterConn()
    if err != nil {
        panic(err)
    }
    cm, err := loukubernetes.GetConfigMap(clientset, common.Namespace, common.Name)
}
```

## Contributing

PRs accepted.

## License

MIT © Richard McRichface