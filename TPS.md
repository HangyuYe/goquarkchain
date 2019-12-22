# TPS Competition Questionnaire

*Please replace the square brackets and the text in it with your answers*

**Number of CPUs**
```
[
  {
    "Cluster 0":
    "Independent Master" = true
    "Salve Servers" = "64 Core * 42"
  },
  {
    "Cluster 1":
    "Independent Master" = false
    "Slaves Servers" = "64 Core * 32"
  },
  {
    "Cluster 2":
    "Independent Master" = false
    "Slaves Servers" = "64 Core * 16"
  }
]
```

**Memory (GB)**
```
[
    256 GB --> 64 Core
]
```
**Storage (GB)**

[ SSD 50 GB ]

**Network**

[ 1 Gbps LAN ]

**Machine Type (Optional)**

[ 腾讯云：标准型SA1 + 标准型S3 ]

**Command Lines for Running Cluster**
```
go run deployer_config.go
```
For more details please see deployer_config.json

**Peak TPS**

[ **Cluster 0: 1121785.00** ]
[ **Cluster 1: 1000023.00** ]

**Video URL**

[Youtube](https://youtu.be/_PDzU0GXjlw)   

**Output From `stats` Tool**


**Cluster Configurations**

[cluster_config_template](https://github.com/HangyuYe/goquarkchain/blob/master/tests/loadtest/deployer/cluster_config_template.json)  
[deployer_config](https://github.com/HangyuYe/goquarkchain/blob/master/tests/loadtest/deployer/deployConfig.json)  

**Additional Comment**
