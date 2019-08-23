# Open data catalog of the Ukraine
[Here](https://da.org.ua/) you can find structured and more readable data from https://data.gov.ua

### Installation
```shell 
git clone git@github.com:dmitry-udod/codes_go.git
go run main.go
```

### Import Data

Download and un-zip data archive

```shell 
wget https://data.gov.ua/dataset/b244f35a-e50a-4a80-b704-032c42ba8142/resource/b0476139-62f2-4ede-9d3b-884ad99afd08/download/15-ufop.zip
unzip 15-ufop.zip
go run main.go --import-legal-entity 15.1-EX_XML_EDR_UO.xml
go run main.go --import-fop 15.2-EX_XML_EDR_FOP.xml
go run main.go --import-terrorist zBlackListFull.xml
```

### Dokku deployment

##### Run on Remote Server 
- install Dokku http://dokku.viewdocs.io/dokku/getting-started/installation/

Create dokku app:

```shell
dokku apps:create da.org.ua

```

Map HTTP port:
```shell 
dokku  proxy:ports-set da.org.ua http:80:5000
```

Enable HTTPS:
```shell 
sudo dokku plugin:install https://github.com/dokku/dokku-letsencrypt.git
dokku config:set --global DOKKU_LETSENCRYPT_EMAIL=YOUR_EMAIL
dokku domains:enable da.org.ua
dokku domains:set da.org.ua da.org.ua
dokku letsencrypt da.org.ua
dokku proxy:ports-add da.org.ua https:443:5000
```

Integrate with ElasticSearch:
```shell 
sudo dokku plugin:install https://github.com/dokku/dokku-elasticsearch.git elasticsearch
export ELASTICSEARCH_IMAGE_VERSION="7.2.0"
echo 'vm.max_map_count=262144' | sudo tee -a /etc/sysctl.conf; sudo sysctl -p
dokku elasticsearch:create es
dokku elasticsearch:info es --status #should return `running`
dokku elasticsearch:link es da.org.ua
```

Edit `/var/lib/dokku/services/elasticsearch/es/config/elasticsearch.yml` and add in the end of file: 
```code
transport.host: localhost
```

```shell
dokku elasticsearch:restart es
```