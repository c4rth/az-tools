# az-tools
Azure Cloud Shell Tools : switch active Subscription and get credentials for an Aks

## Config file
location: $HOME/.az-tools/config.yaml
```yaml
subscriptions:
- name: subscription1
  resource-groups:
    - name: resource-group-1
      aks:
        - name: aks11
        - name: aks12
    - name: resource-group-2
      aks:
        - name: aks21
        - name: aks22
- name: subscription2
  resource-groups:
    - name: resource-group-3
      aks:
        - name: aks31
        - name: aks32
    - name: resource-group-4
      aks:
        - name: aks41
        - name: aks42
```
Will show
```
╔Az tools═════════════════════════╗
║ subscriptions                   ║
║ ├──subscription1                ║
║ │  ├──resource-group-1 / aks11  ║
║ │  ├──resource-group-1 / aks12  ║
║ │  ├──resource-group-2 / aks21  ║
║ │  └──resource-group-2 / aks22  ║
║ └──subscription2                ║
║    ├──resource-group-3 / aks31  ║
║    ├──resource-group-3 / aks32  ║
║    ├──resource-group-4 / aks41  ║
║    └──resource-group-4 / aks42  ║
╚═════════════════════════════════╝
```