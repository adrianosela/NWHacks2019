## Steps to Deploy Docker App on Azure

**Login to azure**

```
$ az login
```

**Create a resource group**

```
$ az group create --name NWHacks2019 --location "West US"
```

```
$ az group list
[
  {
    "id": "/subscriptions/6ef3c11b-f061-4a57-bb11-20c2ea006720/resourceGroups/NWHacks2019",
    "location": "westus",
    "managedBy": null,
    "name": "NWHacks2019",
    "properties": {
      "provisioningState": "Succeeded"
    },
    "tags": null
  }
]
```

**Create a Linux based app service plan --> will use a Linux worker to host the Docker app**

```
$ az appservice plan create --name ServicePlan9952f74b-9762 --resource-group NWHacks2019 --sku S1 --is-linux
```

**Create a webapp**

```
$ az webapp create --resource-group NWHacks2019 --plan ServicePlan9952f74b-9762  --name refill --deployment-container-image-name adrianosela/nwhacks_api:laptop_push
```