## Deployments

Deployments have been automated in the form of Makefile targets which call the docker CLI for building/pushing and the Microsoft Azure CLI for creating a "web application" in Azure.

### Prerequisites

The host machine must have the following:

* Docker CLI (docker) - must also be logged into docker acct
* Azure CLI (az) - must also be logged into azure acct

The azure account must have a database in THE SAME resource group as the "web application"

#### This is what a full deployment should look like:

```
08:29 $ make deploy
GOOS=linux GOARCH=amd64 go build -a -o nwhacks_api_linux 	# build a golang binary for linux
docker build -t nwhacks_api:latest . 				# build docker image from this dir
Sending build context to Docker daemon  8.143MB
Step 1/5 : FROM alpine:latest
latest: Pulling from library/alpine
cd784148e348: Pull complete
Digest: sha256:46e71df1e5191ab8b8034c5189e325258ec44ea739bba1e5645cff83c9048ff1
Status: Downloaded newer image for alpine:latest
 ---> 3f53bb00af94
Step 2/5 : RUN apk add --update bash curl && rm -rf /var/cache/apk/*
 ---> Running in 202a28fbe930
fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/community/x86_64/APKINDEX.tar.gz
(1/10) Installing ncurses-terminfo-base (6.1_p20180818-r1)
(2/10) Installing ncurses-terminfo (6.1_p20180818-r1)
(3/10) Installing ncurses-libs (6.1_p20180818-r1)
(4/10) Installing readline (7.0.003-r0)
(5/10) Installing bash (4.4.19-r1)
Executing bash-4.4.19-r1.post-install
(6/10) Installing ca-certificates (20171114-r3)
(7/10) Installing nghttp2-libs (1.32.0-r0)
(8/10) Installing libssh2 (1.8.0-r3)
(9/10) Installing libcurl (7.61.1-r1)
(10/10) Installing curl (7.61.1-r1)
Executing busybox-1.28.4-r2.trigger
Executing ca-certificates-20171114-r3.trigger
OK: 15 MiB in 23 packages
Removing intermediate container 202a28fbe930
 ---> 1535a8620936
Step 3/5 : ADD nwhacks_api_linux /bin/nwhacks_api_linux
 ---> e1925d637512
Step 4/5 : EXPOSE 80
 ---> Running in a1b02fe168af
Removing intermediate container a1b02fe168af
 ---> 2e50b21c9d74
Step 5/5 : CMD ["/bin/nwhacks_api_linux"]
 ---> Running in 1bab119a06b6
Removing intermediate container 1bab119a06b6
 ---> 4da798eb7995
Successfully built 4da798eb7995
Successfully tagged nwhacks_api:latest
docker tag nwhacks_api:latest adrianosela/nwhacks_api:laptop_push # tag image with the docker hub repo name
docker push adrianosela/nwhacks_api:laptop_push			# push image to docker hub
The push refers to repository [docker.io/adrianosela/nwhacks_api]
62df9b94f910: Pushed
62843a35da11: Pushed
7bff100f35cb: Layer already exists
laptop_push: digest: sha256:8b2c57f40cce412f4321d379119e488ca7ec632bd0e0f623cc0c8db1a7dfb5d2 size: 950
az webapp delete --name ezpill --resource-group NWHacks2019 	# kill current azure deployment
az appservice plan create --name nwhackssp --resource-group NWHacks2019 --sku S1 --is-linux # service plan in same resource group as DB
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
  "location": "westus",
  "maximumNumberOfWorkers": 10,
  "name": "nwhackssp",
  "numberOfSites": 0,
  "perSiteScaling": false,
  "provisioningState": null,
  "reserved": true,
  "resourceGroup": "NWHacks2019",
  "sku": {
    "capabilities": null,
    "capacity": 1,
    "family": null,
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
az webapp create --resource-group NWHacks2019 --plan nwhackssp --name ezpill --deployment-container-image-name adrianosela/nwhacks_api:laptop_push # deploy to azure
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
  "ftpPublishingUrl": "ftp://waws-prod-bay-081.ftp.azurewebsites.windows.net/site/wwwroot",
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
  "lastModifiedTimeUtc": "2019-01-28T16:33:16.063333",
  "location": "West US",
  "maxNumberOfWorkers": null,
  "name": "ezpill",
  "outboundIpAddresses": "13.64.73.110,40.118.133.8,40.118.169.141,40.118.253.162,13.64.147.140",
  "possibleOutboundIpAddresses": "13.64.73.110,40.118.133.8,40.118.169.141,40.118.253.162,13.64.147.140,52.160.85.217,13.93.238.69",
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