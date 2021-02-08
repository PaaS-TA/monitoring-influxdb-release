## paas-ta-monitoring-influxdb-release

### Create PaaS-TA Monitoring InfluxDB Release   
  - Download the latest PaaS-TA Monitoring InfluxDB Release    
    ```   
    $ git clone https://github.com/PaaS-TA/paas-ta-monitoring-influxdb-release.git    
    $ cd paas-ta-monitoring-influxdb-release   
    ```   
  - Download & Copy "source files" into the src directory    
    ```   
    ## download source files    
    $ wget -O src.zip https://45.248.73.44/index.php/s/dPeHkRFer5zmQd5/download
    
    ## unzip download source files    
    $ unzip src.zip    

    ## final src directory
    src
      ├── bootstrapper
      │   └── src
      │       └── bootstrapper
      ├── chronograf
      │   └── chronograf-1.4.3.0_linux_amd64.tar.gz
      ├── golang
      │   └── go1.9.2.linux-amd64.tar.gz
      ├── influxdb
      │   ├── influxdb-1.8.4_linux_amd64.tar.gz
      │   └── src
      │       ├── collectd.org
      │       ├── github.com
      │       ├── go.uber.org
      │       └── golang.org
      └── pidutils.sh
    
    ```  
  - Create PaaS-TA Portal UI Release   
    ```   
    ## <VERSION> :: release version (e.g. 1.0.0)   
    ## <RELEASE_TARBALL_PATH> :: release file path (e.g. /home/ubuntu/workspace/monitoring-influxdb-release-<VERSION>.tgz)    
    $ bosh -e <bosh_name> create-release --name=monitoring-influxdb-release --sha2 --version=<VERSION> --tarball=<RELEASE_TARBALL_PATH> --force   
    ```    
### Deployment
- https://github.com/PaaS-TA/monitoring-deployment   

