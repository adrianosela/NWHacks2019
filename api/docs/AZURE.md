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
$ az appservice plan create --name nwhackssp --resource-group NWHacks2019 --sku S1 --is-linux
{
  "adminSiteName": null,
  "freeOfferExpirationTime": null,
  "geoRegion": "West US",
  "hostingEnvironmentProfile": null,
  "hyperV": false,
  "id": "/subscriptions/6ef3c11b-f061-4a57-bb11-20c2ea006720/resourceGroups/NWHacks2019/providers/Microsoft.Web/serverfarms/nwhackssp",
  "isSpot": false,
  "isXenon": false,
  "kind": "linux",
  "location": "West US",
  "maximumNumberOfWorkers": 10,
  "name": "nwhackssp",
  "numberOfSites": 0,
  "perSiteScaling": false,
  "provisioningState": "Succeeded",
  "reserved": true,
  "resourceGroup": "NWHacks2019",
  "sku": {
    "capabilities": null,
    "capacity": 1,
    "family": "S",
    "locations": null,
    "name": "S1",
    "size": "S1",
    "skuCapacity": null,
    "tier": "Standard"
  },
  "spotExpirationTime": null,
  "status": "Ready",
  "subscription": "6ef3c11b-f061-4a57-bb11-20c2ea006720",
  "tags": null,
  "targetWorkerCount": 0,
  "targetWorkerSizeId": 0,
  "type": "Microsoft.Web/serverfarms",
  "workerTierName": null
}
```

**Create a webapp**

```
$ az webapp create --resource-group NWHacks2019 --plan nwhackssp --name ezpill --deployment-container-image-name adrianosela/nwhacks_api:laptop_push
{
  "availabilityState": "Normal",
  "clientAffinityEnabled": true,
  "clientCertEnabled": false,
  "cloningInfo": null,
  "containerSize": 0,
  "dailyMemoryTimeQuota": 0,
  "defaultHostName": "ezpill.azurewebsites.net",
  "enabled": true,
  "enabledHostNames": [
    "ezpill.azurewebsites.net",
    "ezpill.scm.azurewebsites.net"
  ],
  "ftpPublishingUrl": "ftp://waws-prod-bay-109.ftp.azurewebsites.windows.net/site/wwwroot",
  "hostNameSslStates": [
    {
      "hostType": "Standard",
      "ipBasedSslResult": null,
      "ipBasedSslState": "NotConfigured",
      "name": "ezpill.azurewebsites.net",
      "sslState": "Disabled",
      "thumbprint": null,
      "toUpdate": null,
      "toUpdateIpBasedSsl": null,
      "virtualIp": null
    },
    {
      "hostType": "Repository",
      "ipBasedSslResult": null,
      "ipBasedSslState": "NotConfigured",
      "name": "ezpill.scm.azurewebsites.net",
      "sslState": "Disabled",
      "thumbprint": null,
      "toUpdate": null,
      "toUpdateIpBasedSsl": null,
      "virtualIp": null
    }
  ],
  "hostNames": [
    "ezpill.azurewebsites.net"
  ],
  "hostNamesDisabled": false,
  "hostingEnvironmentProfile": null,
  "httpsOnly": false,
  "hyperV": false,
  "id": "/subscriptions/6ef3c11b-f061-4a57-bb11-20c2ea006720/resourceGroups/NWHacks2019/providers/Microsoft.Web/sites/ezpill",
  "identity": null,
  "isDefaultContainer": null,
  "isXenon": false,
  "kind": "app,linux,container",
  "lastModifiedTimeUtc": "2019-01-27T03:05:13.456666",
  "location": "West US",
  "maxNumberOfWorkers": null,
  "name": "ezpill",
  "outboundIpAddresses": "104.42.78.153,40.118.225.64,104.42.151.111,104.42.252.192,104.42.254.180",
  "possibleOutboundIpAddresses": "104.42.78.153,40.118.225.64,104.42.151.111,104.42.252.192,104.42.254.180,40.118.227.167,40.118.224.22,40.118.226.153",
  "repositorySiteName": "ezpill",
  "reserved": true,
  "resourceGroup": "NWHacks2019",
  "scmSiteAlsoStopped": false,
  "serverFarmId": "/subscriptions/6ef3c11b-f061-4a57-bb11-20c2ea006720/resourceGroups/NWHacks2019/providers/Microsoft.Web/serverfarms/nwhackssp",
  "siteConfig": null,
  "slotSwapStatus": null,
  "state": "Running",
  "suspendedTill": null,
  "tags": null,
  "targetSwapSlot": null,
  "trafficManagerHostNames": null,
  "type": "Microsoft.Web/sites",
  "usageState": "Normal"
}
```