## proposed structure (maybe put this is a self contained repo for all deployments to pull from)
### These standard components can be thought of as the following

- we typically would have re-usable or common components across different deployments (Networks, Stemcells, Releases, etc)
- these should probably exist in their own package or repo so that they can be leveraged by all deployments
- these also benefit from being centrally maintained, and centrally distributable(imported into deployments)
- these are basically factories for re-usable or common elements of deployments
- these are used to compose specific components in a standard context
- this is just a proposed convention not a hard requirement

