# az-tools
Azure Cloud Shell Tools : switch active Subscription and get credentials for an cluster

## Build and test

```sh
go test ./...
go build -trimpath -ldflags='-s -w' -o bin/k8switch ./cmd/k8switch
```

## Config file
location: $HOME/.az-tools/config.yaml
```yaml
subscriptions:
- name: subscription1
  resource-groups:
    - name: resource-group-1
      cluster:
        - name: cluster11
        - name: cluster12
    - name: resource-group-2
      cluster:
        - name: cluster21
        - name: cluster22
- name: subscription2
  resource-groups:
    - name: resource-group-3
      cluster:
        - name: cluster31
        - name: cluster32
    - name: resource-group-4
      cluster:
        - name: cluster41
        - name: cluster42
```
Will show
```

├──rancher
│  ├──rancher1
│  ├──rancher2
│  └──rancher3
├──azure
│  ├──subscription1                  
│  ├──resource-group-1 / cluster11
│  ├──resource-group-1 / cluster12
│  ├──resource-group-2 / cluster21
│  └──resource-group-2 / cluster22
│  └──subscription2                  
      ├──resource-group-3 / cluster31
      ├──resource-group-3 / cluster32
      ├──resource-group-4 / cluster41
      └──resource-group-4 / cluster42

```