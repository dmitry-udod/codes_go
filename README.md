# Open data catalog of the Ukraine
Source: https://data.gov.ua

### Instalation

### Dokku deployment

##### Remote Server 
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
dokku proxy:ports-remove da.org.ua https:443:5000
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