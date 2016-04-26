
## Creating a deployment for a release

```
$ enaml generate-jobs https://bosh.io/d/github.com/concourse/concourse?v=1.1.0
Could not find release in local cache. Downloading now.
317624091/317624091
completed generating release job structs for  https://bosh.io/d/github.com/concourse/concourse?v=1.1.0

$ ls *
concourse.go      concourse_test.go suite_test.go

enaml-gen:
atc          baggageclaim blackbox     groundcrew   postgresql   tsa
```

## pull it into your deployment

```golang


package concourse

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/atc"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse/enaml-gen/tsa"
)
```
