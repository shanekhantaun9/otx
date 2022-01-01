# otx

[![GitHub stars](https://badgen.net/github/stars/0xsheinn/otx)](https://gitHub.com/0xsheinn/otx/stargazers/)
[![GitHub forks](https://img.shields.io/github/forks/0xsheinn/otx)](https://github.com/0xsheinn/otx/network/)
[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/cloudposse.svg?style=social&label=Follow%20%40l33t_Enough)](https://twitter.com/l33t_En0ugh)


## Description
- This tool is base on AlienVault Open Threat Exchange (OTX)? and this tool can help you to extract all the urls endpoints which can be vulnerable or leak sensitive informations etc.

## Installation
```
go install github.com/0xsheinn/otx@latest
```
```
git clone https://github.com/0xsheinn/otx.git
cd otx
go build main.go
sudo mv otx /usr/bin/otx
```

## Usage
```
cat live-domain.txt | otx 
cat live-domain.txt | otx -c 50
echo "https://facebook.com" | otx -c 50
```

<img src=example.png>


### Reference 
- https://otx.alienvault.com/


#### Support
[![Donate with Bitcoin](https://en.cryptobadges.io/badge/small/3ECwMfc88iN9DtjHkoYADCMoKeURQpoym1)](https://en.cryptobadges.io/donate/3ECwMfc88iN9DtjHkoYADCMoKeURQpoym1)
