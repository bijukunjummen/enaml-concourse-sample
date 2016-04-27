package concourse

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/atc"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/baggageclaim"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/garden"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/groundcrew"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/postgresql"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/tsa"
)

const (
	concourseReleaseName string = "concourse"
	gardenReleaseName    string = "garden-linux"
)

//Deployment -
type Deployment struct {
	enaml.Deployment
	manifest          *enaml.DeploymentManifest
	ConcourseURL      string
	ConcourseUserName string
	ConcoursePassword string
	DirectorUUID      string
	NetworkName       string
	WebIPs            []string
	WebInstances      int
	CloudConfig       bool
	WebAZs            []string
	DatabaseAZs       []string
	WorkerAZs         []string
	DeploymentName    string
	NetworkRange      string
	NetworkGateway    string
	StemcellAlias     string
	PostgresPassword  string
}

//NewDeployment -
func NewDeployment() (d Deployment) {
	d = Deployment{}
	d.manifest = new(enaml.DeploymentManifest)
	d.CloudConfig = false
	d.WebInstances = 1
	d.DeploymentName = "concourse"
	d.PostgresPassword = "dummy-postgres-password"
	return
}

//Initialize -
func (d *Deployment) Initialize() (err error) {
	var resourcePoolName = ""

	//TODO Add validations to provide feedback on invalid property configuration
	/*if !d.isStrongPass(d.ConcoursePassword) {
		err = fmt.Errorf("Sorry. The given password is too weak")
	}*/

	d.manifest.SetName(d.DeploymentName)
	d.manifest.SetDirectorUUID(d.DirectorUUID)
	d.manifest.AddReleaseByName(concourseReleaseName)
	d.manifest.AddReleaseByName(gardenReleaseName)

	if d.CloudConfig {
		d.manifest.AddStemcellByName("ubuntu-trusty", d.StemcellAlias)
	} else {
		resourcePool := d.CreateResourcePool(d.NetworkName)
		d.manifest.AddResourcePool(resourcePool)
		resourcePoolName = resourcePool.Name

		compilation := d.CreateCompilation(d.NetworkName)
		d.manifest.SetCompilation(compilation)

		deploymentNetwork := d.CreateManualDeploymentNetwork(d.NetworkName, d.NetworkRange, d.NetworkGateway, d.WebIPs)
		d.manifest.AddNetwork(deploymentNetwork)
	}

	update := d.CreateUpdate()
	d.manifest.SetUpdate(update)

	web := d.CreateWebInstanceGroup(resourcePoolName)
	d.manifest.AddInstanceGroup(web)

	db := d.CreateDatabaseInstanceGroup(resourcePoolName)
	d.manifest.AddInstanceGroup(db)

	worker := d.CreateWorkerInstanceGroup(resourcePoolName)
	d.manifest.AddInstanceGroup(worker)

	return
}

//CreateWebInstanceGroup -
func (d *Deployment) CreateWebInstanceGroup(resourcePoolName string) (web *enaml.InstanceGroup) {
	web = &enaml.InstanceGroup{
		Name:         "web",
		Instances:    d.WebInstances,
		ResourcePool: resourcePoolName,
		VMType:       "web",
		AZs:          d.WebAZs,
		Stemcell:     d.StemcellAlias,
	}
	web.AddNetwork(enaml.Network{
		Name:      d.NetworkName,
		StaticIPs: d.WebIPs,
	})
	atc := enaml.NewInstanceJob("atc", concourseReleaseName, atc.Atc{
		ExternalUrl:        d.ConcourseURL,
		BasicAuthUsername:  d.ConcourseUserName,
		BasicAuthPassword:  d.ConcoursePassword,
		PostgresqlDatabase: "atc",
	})
	tsa := enaml.NewInstanceJob("tsa", concourseReleaseName, tsa.Tsa{})
	web.AddJob(atc)
	web.AddJob(tsa)
	return
}

//CreateDatabaseInstanceGroup -
func (d *Deployment) CreateDatabaseInstanceGroup(resourcePoolName string) (db *enaml.InstanceGroup) {
	db = &enaml.InstanceGroup{
		Name:           "db",
		Instances:      1,
		ResourcePool:   resourcePoolName,
		PersistentDisk: 10240,
		VMType:         "database",
		AZs:            d.DatabaseAZs,
		Stemcell:       d.StemcellAlias,
	}
	db.AddNetwork(enaml.Network{
		Name: d.NetworkName,
	})
	dbs := make([]DBName, 1)
	dbs[0] = DBName{
		Name:     "atc",
		Role:     "atc",
		Password: d.PostgresPassword,
	}
	db.AddJob(enaml.NewInstanceJob("postgresql", concourseReleaseName, postgresql.Postgresql{
		Databases: dbs,
	}))
	return
}

//CreateWorkerInstanceGroup -
func (d *Deployment) CreateWorkerInstanceGroup(resourcePoolName string) (worker *enaml.InstanceGroup) {
	worker = &enaml.InstanceGroup{
		Name:         "worker",
		Instances:    1,
		ResourcePool: resourcePoolName,
		VMType:       "worker",
		AZs:          d.WorkerAZs,
		Stemcell:     d.StemcellAlias,
	}

	worker.AddNetwork(enaml.Network{
		Name: d.NetworkName,
	})
	worker.AddJob(enaml.NewInstanceJob("groundcrew", concourseReleaseName, groundcrew.Groundcrew{}))
	worker.AddJob(enaml.NewInstanceJob("baggageclaim", concourseReleaseName, baggageclaim.Baggageclaim{}))
	worker.AddJob(enaml.NewInstanceJob("garden", gardenReleaseName, Garden{
		garden.Garden{
			ListenAddress:   "0.0.0.0:7777",
			ListenNetwork:   "tcp",
			AllowHostAccess: true,
		},
	}))
	return
}

//CreateManualDeploymentNetwork -
func (d *Deployment) CreateManualDeploymentNetwork(networkName, networkRange, networkGateway string, webIPs []string) (network enaml.ManualNetwork) {
	network = enaml.ManualNetwork{
		Name: networkName,
		Type: "manual",
	}
	subnets := make([]enaml.Subnet, 1)
	subnet := enaml.Subnet{
		Range:   networkRange,
		Gateway: networkGateway,
		Static:  webIPs,
	}
	subnets[0] = subnet
	network.Subnets = subnets

	return
}

//CreateUpdate -
func (d *Deployment) CreateUpdate() (update enaml.Update) {
	update = enaml.Update{
		Canaries:        1,
		MaxInFlight:     3,
		Serial:          false,
		CanaryWatchTime: "1000-60000",
		UpdateWatchTime: "1000-60000",
	}

	return
}

//CreateResourcePool -
func (d *Deployment) CreateResourcePool(networkName string) (resourcePool enaml.ResourcePool) {
	const resourcePoolName = "concourse"
	resourcePool = enaml.ResourcePool{
		Name:    resourcePoolName,
		Network: networkName,
		Stemcell: enaml.Stemcell{
			Name:    "bosh-warden-boshlite-ubuntu-trusty-go_agent",
			Version: "latest",
		},
	}

	return
}

//CreateCompilation -
func (d *Deployment) CreateCompilation(networkName string) (compilation enaml.Compilation) {
	compilation = enaml.Compilation{
		Network: networkName,
		Workers: 3,
	}

	return
}

func (d Deployment) isStrongPass(pass string) (ok bool) {
	ok = false
	if len(pass) > 8 {
		ok = true
	}
	return
}

func insureHAInstanceCount(instances int) int {
	if instances < 2 {
		instances = 2
	}
	return instances
}

//GetDeployment -
func (d Deployment) GetDeployment() enaml.DeploymentManifest {
	return *d.manifest
}
