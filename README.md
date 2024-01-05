# az-tools
Azure Cloud Shell Tools : switch Subscription and Aks

## Config file
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
        - name: aks31
        - name: aks32
```