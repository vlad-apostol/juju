// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Code generated by github.com/juju/juju/generate/filetoconst. DO NOT EDIT.

package cloud

const fallbackPublicCloudInfo = `# DO NOT EDIT, will be overwritten, use "juju update-public-clouds" to refresh.
clouds:
  aws:
    type: ec2
    description: Amazon Web Services
    auth-types: [ access-key ]
    regions:
      us-east-1:
        endpoint: https://ec2.us-east-1.amazonaws.com
      us-east-2:
        endpoint: https://ec2.us-east-2.amazonaws.com
      us-west-1:
        endpoint: https://ec2.us-west-1.amazonaws.com
      us-west-2:
        endpoint: https://ec2.us-west-2.amazonaws.com
      ca-central-1:
        endpoint: https://ec2.ca-central-1.amazonaws.com
      eu-west-1:
        endpoint: https://ec2.eu-west-1.amazonaws.com
      eu-west-2:
        endpoint: https://ec2.eu-west-2.amazonaws.com
      eu-west-3:
        endpoint: https://ec2.eu-west-3.amazonaws.com
      eu-central-1:
        endpoint: https://ec2.eu-central-1.amazonaws.com
      eu-central-2:
        endpoint: https://ec2.eu-central-2.amazonaws.com
      eu-north-1:
        endpoint: https://ec2.eu-north-1.amazonaws.com
      eu-south-1:
        endpoint: https://ec2.eu-south-1.amazonaws.com
      eu-south-2:
        endpoint: https://ec2.eu-south-2.amazonaws.com
      af-south-1:
        endpoint: https://ec2.af-south-1.amazonaws.com
      ap-east-1:
        endpoint: https://ec2.ap-east-1.amazonaws.com
      ap-south-1:
        endpoint: https://ec2.ap-south-1.amazonaws.com
      ap-south-2:
        endpoint: https://ec2.ap-south-2.amazonaws.com
      ap-southeast-1:
        endpoint: https://ec2.ap-southeast-1.amazonaws.com
      ap-southeast-2:
        endpoint: https://ec2.ap-southeast-2.amazonaws.com
      ap-southeast-3:
        endpoint: https://ec2.ap-southeast-3.amazonaws.com
      ap-southeast-4:
        endpoint: https://ec2.ap-southeast-4.amazonaws.com
      ap-northeast-1:
        endpoint: https://ec2.ap-northeast-1.amazonaws.com
      ap-northeast-2:
        endpoint: https://ec2.ap-northeast-2.amazonaws.com
      ap-northeast-3:
        endpoint: https://ec2.ap-northeast-3.amazonaws.com
      me-south-1:
        endpoint: https://ec2.me-south-1.amazonaws.com
      me-central-1:
        endpoint: https://ec2.me-central-1.amazonaws.com
      sa-east-1:
        endpoint: https://ec2.sa-east-1.amazonaws.com
  aws-china:
    type: ec2
    description: Amazon China
    auth-types: [ access-key ]
    regions:
      cn-north-1:
        endpoint: https://ec2.cn-north-1.amazonaws.com.cn
      cn-northwest-1:
        endpoint: https://ec2.cn-northwest-1.amazonaws.com.cn
  aws-gov:
    type: ec2
    description: Amazon (USA Government)
    auth-types: [ access-key ]
    regions:
      us-gov-west-1:
        endpoint: https://ec2.us-gov-west-1.amazonaws.com
      us-gov-east-1:
        endpoint: https://ec2.us-gov-east-1.amazonaws.com
  google:
    type: gce
    description: Google Cloud Platform
    auth-types: [ jsonfile, oauth2 ]
    regions:
      us-east1:
        endpoint: https://www.googleapis.com
      us-east4:
        endpoint: https://www.googleapis.com
      us-central1:
        endpoint: https://www.googleapis.com
      us-west1:
        endpoint: https://www.googleapis.com
      us-west2:
        endpoint: https://www.googleapis.com
      us-west3:
        endpoint: https://www.googleapis.com
      us-west4:
        endpoint: https://www.googleapis.com
      asia-east1:
        endpoint: https://www.googleapis.com
      asia-east2:
        endpoint: https://www.googleapis.com
      asia-northeast1:
        endpoint: https://www.googleapis.com
      asia-northeast2:
        endpoint: https://www.googleapis.com
      asia-northeast3:
        endpoint: https://www.googleapis.com
      asia-south1:
        endpoint: https://www.googleapis.com
      asia-southeast1:
        endpoint: https://www.googleapis.com
      asia-southeast2:
        endpoint: https://www.googleapis.com
      australia-southeast1:
        endpoint: https://www.googleapis.com
      europe-central2:
        endpoint: https://www.googleapis.com
      europe-north1:
        endpoint: https://www.googleapis.com
      europe-west1:
        endpoint: https://www.googleapis.com
      europe-west2:
        endpoint: https://www.googleapis.com
      europe-west3:
        endpoint: https://www.googleapis.com
      europe-west4:
        endpoint: https://www.googleapis.com
      europe-west6:
        endpoint: https://www.googleapis.com
      northamerica-northeast1:
        endpoint: https://www.googleapis.com
      southamerica-east1:
        endpoint: https://www.googleapis.com
  azure:
    # Note: the storage endpoint definitions below are no longer used by
    # recent Juju versions, and are retained for compatibility only.
    type: azure
    description: Microsoft Azure
    auth-types: [ interactive, service-principal-secret, managed-identity ]
    regions:
      centralus:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      eastus:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      eastus2:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      northcentralus:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      southcentralus:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      westcentralus:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      westus:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      westus2:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      westus3:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      northeurope:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      westeurope:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      eastasia:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      southeastasia:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      japaneast:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      japanwest:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      brazilsouth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      australiacentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      australiacentral2:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      australiaeast:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      australiasoutheast:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      centralindia:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      southindia:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      westindia:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      canadacentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      canadaeast:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      uksouth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      ukwest:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      koreacentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      koreasouth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      francecentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      francesouth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      southafricanorth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      southafricawest:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      germanynorth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      germanywestcentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      uaecentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      uaenorth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      norwayeast:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      norwaywest:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      switzerlandnorth:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      switzerlandwest:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      swedencentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      qatarcentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://login.microsoftonline.com
      polandcentral:
        endpoint: https://management.azure.com
        storage-endpoint: https://core.windows.net
        identity-endpoint: https://graph.windows.net
  azure-china:
    type: azure
    description: Microsoft Azure China
    auth-types: [ interactive, service-principal-secret ]
    regions:
      chinaeast:
        endpoint: https://management.chinacloudapi.cn
        storage-endpoint: https://core.chinacloudapi.cn
        identity-endpoint: https://graph.chinacloudapi.cn
      chinaeast2:
        endpoint: https://management.chinacloudapi.cn
        storage-endpoint: https://core.chinacloudapi.cn
        identity-endpoint: https://graph.chinacloudapi.cn
      chinanorth:
        endpoint: https://management.chinacloudapi.cn
        storage-endpoint: https://core.chinacloudapi.cn
        identity-endpoint: https://graph.chinacloudapi.cn
      chinanorth2:
        endpoint: https://management.chinacloudapi.cn
        storage-endpoint: https://core.chinacloudapi.cn
        identity-endpoint: https://graph.chinacloudapi.cn
  oracle:
    type: oci
    description: Oracle Cloud Infrastructure
    auth-types: [ httpsig ]
    regions:
      af-johannesburg-1:
        endpoint: https://iaas.af-johannesburg-1.oraclecloud.com
      ap-chiyoda-1:
        endpoint: https://iaas.ap-chiyoda-1.oraclecloud.com
      ap-chuncheon-1:
        endpoint: https://iaas.ap-chuncheon-1.oraclecloud.com
      ap-dcc-canberra-1:
        endpoint: https://iaas.ap-dcc-canberra-1.oraclecloud.com
      ap-hyderabad-1:
        endpoint: https://iaas.ap-hyderabad-1.oraclecloud.com
      ap-ibaraki-1:
        endpoint: https://iaas.ap-ibaraki-1.oraclecloud.com
      ap-melbourne-1:
        endpoint: https://iaas.ap-melbourne-1.oraclecloud.com
      ap-mumbai-1:
        endpoint: https://iaas.ap-mumbai-1.oraclecloud.com
      ap-osaka-1:
        endpoint: https://iaas.ap-osaka-1.oraclecloud.com
      ap-seoul-1:
        endpoint: https://iaas.ap-seoul-1.oraclecloud.com
      ap-singapore-1:
        endpoint: https://iaas.ap-singapore-1.oraclecloud.com
      ap-sydney-1:
        endpoint: https://iaas.ap-sydney-1.oraclecloud.com
      ap-tokyo-1:
        endpoint: https://iaas.ap-tokyo-1.oraclecloud.com
      ca-montreal-1:
        endpoint: https://iaas.ca-montreal-1.oraclecloud.com
      ca-toronto-1:
        endpoint: https://iaas.ca-toronto-1.oraclecloud.com
      eu-amsterdam-1:
        endpoint: https://iaas.eu-amsterdam-1.oraclecloud.com
      eu-dcc-dublin-1:
        endpoint: https://iaas.eu-dcc-dublin-1.oraclecloud.com
      eu-dcc-dublin-2:
        endpoint: https://iaas.eu-dcc-dublin-2.oraclecloud.com
      eu-dcc-milan-1:
        endpoint: https://iaas.eu-dcc-milan-1.oraclecloud.com
      eu-dcc-milan-2:
        endpoint: https://iaas.eu-dcc-milan-2.oraclecloud.com
      eu-dcc-rating-1:
        endpoint: https://iaas.eu-dcc-rating-1.oraclecloud.com
      eu-dcc-rating-2:
        endpoint: https://iaas.eu-dcc-rating-2.oraclecloud.com
      eu-frankfurt-1:
        endpoint: https://iaas.eu-frankfurt-1.oraclecloud.com
      eu-frankfurt-2:
        endpoint: https://iaas.eu-frankfurt-2.oraclecloud.com
      eu-jovanovac-1:
        endpoint: https://iaas.eu-jovanovac-1.oraclecloud.com
      eu-madrid-1:
        endpoint: https://iaas.eu-madrid-1.oraclecloud.com
      eu-madrid-2:
        endpoint: https://iaas.eu-madrid-2.oraclecloud.com
      eu-marseille-1:
        endpoint: https://iaas.eu-marseille-1.oraclecloud.com
      eu-milan-1:
        endpoint: https://iaas.eu-milan-1.oraclecloud.com
      eu-paris-1:
        endpoint: https://iaas.eu-paris-1.oraclecloud.com
      eu-stockholm-1:
        endpoint: https://iaas.eu-stockholm-1.oraclecloud.com
      eu-zurich-1:
        endpoint: https://iaas.eu-zurich-1.oraclecloud.com
      il-jerusalem-1:
        endpoint: https://iaas.il-jerusalem-1.oraclecloud.com
      me-abudhabi-1:
        endpoint: https://iaas.me-abudhabi-1.oraclecloud.com
      me-dcc-muscat-1:
        endpoint: https://iaas.me-dcc-muscat-1.oraclecloud.com
      me-dubai-1:
        endpoint: https://iaas.me-dubai-1.oraclecloud.com
      me-jeddah-1:
        endpoint: https://iaas.me-jeddah-1.oraclecloud.com
      mx-monterrey-1:
        endpoint: https://iaas.mx-monterrey-1.oraclecloud.com
      mx-queretaro-1:
        endpoint: https://iaas.mx-queretaro-1.oraclecloud.com
      sa-santiago-1:
        endpoint: https://iaas.sa-santiago-1.oraclecloud.com
      sa-saopaulo-1:
        endpoint: https://iaas.sa-saopaulo-1.oraclecloud.com
      sa-vinhedo-1:
        endpoint: https://iaas.sa-vinhedo-1.oraclecloud.com
      uk-cardiff-1:
        endpoint: https://iaas.uk-cardiff-1.oraclecloud.com
      uk-london-1:
        endpoint: https://iaas.uk-london-1.oraclecloud.com
      us-ashburn-1:
        endpoint: https://iaas.us-ashburn-1.oraclecloud.com
      us-chicago-1:
        endpoint: https://iaas.us-chicago-1.oraclecloud.com
      us-langley-1:
        endpoint: https://iaas.us-langley-1.oraclecloud.com
      us-luke-1:
        endpoint: https://iaas.us-luke-1.oraclecloud.com
      us-phoenix-1:
        endpoint: https://iaas.us-phoenix-1.oraclecloud.com
      us-sanjose-1:
        endpoint: https://iaas.us-sanjose-1.oraclecloud.com
`
