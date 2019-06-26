# mkcaasp-containerized

### The list of TO DO things prior to start working with this tool:
- setup your DockerHub account/repo (private repo for a docker image)
- put the variable RegCode in any of your /utils/*.go files ( e.g. 'const RegCode = <your SUSE CaaSP Registration Code>')
- make sure you properly installed concourse and fly
- make sure you have access to your concourse instance + a working token
- make sure you set up a proper pipeline.yml in concourse
- make sure you set your openstack.json file (/caasp-openstack-terraform/openstack.json) with all your proper credentials
- make sure you set the constant ClusterTempl in /utils/data.go with all proper settings related to your openstack Env