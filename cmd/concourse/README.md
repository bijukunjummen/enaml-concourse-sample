There are 2 scripts to demo the ability to generate a valid concourse deployment for both bosh-lite and cloud-config enabled bosh

### Options ###
```
NAME:
   concourse generate - 

USAGE:
   concourse generate [command options] [arguments...]

DESCRIPTION:
   generate concourse yaml for given arguments

OPTIONS:
   --output-filename 						destination for output [$OUTPUT_FILENAME]
   --url 							url for concourse ui [$URL]
   --network-name 						name of network to deploy concourse on [$NETWORK_NAME]
   --bosh-deployment-name 					bosh deployment name [$BOSH_DEPLOYMENT_NAME]
   --bosh-stemcell-alias 					url for concourse ui [$BOSH_STEMCELL_ALIAS]
   --web-ips [--web-ips option --web-ips option]		array of IPs reserved for concourse web-ui [$WEB_IPS]
   --password 							password for concourse ui [$PASSWORD]
   --web-azs [--web-azs option --web-azs option]		array of AZs to deploy concourse web jobs to [$WEB_AZS]
   --database-azs [--database-azs option --database-azs option]	array of AZs to deploy concourse database jobs to [$DATABASE_AZS]
   --worker-azs [--worker-azs option --worker-azs option]	array of AZs to deploy concourse worker jobs to [$WORKER_AZS]
   --bosh-cloud-config 						true/false for generate cloudConfig compatible (default true) [$BOSH_CLOUD_CONFIG]
   --web-instances 						number of web instances (default 1) [$WEB_INSTANCES]
   --network-range 						network range to deploy concourse on - only applies in non-cloud config [$NETWORK_RANGE]
   --network-gateway 						network gateway for concourse - only applies in non-cloud config [$NETWORK_GATEWAY]
   --postgresql-db-pwd 						password for postgresql job [$POSTGRESQL_DB_PWD]
   --bosh-director-uuid 					bosh director uuid (bosh status --uuid) [$BOSH_DIRECTOR_UUID]
   --username 							username for concourse ui [$USERNAME]
```



### Bosh Lite ###
```
./concourse generate --output-filename concourse-bosh-lite.yml \
  --bosh-cloud-config false \
  --bosh-director-uuid 08356199-b2c3-41cc-8319-5bbb980994ab \
  --network-name concourse \
  --network-range 10.244.8.0/24 \
  --network-gateway 10.244.8.1 \
  --url http://10.244.8.2:8080 \
  --username concourse \
  --password concourse \
  --web-ips 10.244.8.2 \
  --web-instances 1
```

### Cloud Config Enabled ###
```
./concourse generate --output-filename concourse-cloudconfig.yml \
  --bosh-cloud-config true \
  --bosh-director-uuid 08356199-b2c3-41cc-8319-5bbb980994ab \
  --network-name private \
  --url http://my.concourse.ci \
  --username concourse \
  --password concourse \
  --web-instances 2 \
  --web-azs z1 \
  --worker-azs z1 \
  --database-azs z1 \
  --bosh-stemcell-alias trusty \
  --postgresql-db-pwd secret
```
