Next, in no particular order:

- [ ] expose threescale API to custom controller
- [ ] check how to deploy controller on master (what ADAC rights do we need `admin-cluster`)
- [ ] filter to select only annotated DC
- [ ] e2e example, in link:/example[/example] define one microservice exposing one CRUD API with a swagger definition
- [ ] introduce link:https://github.com/golang/dep[Dep] Gopkg.toml for clean dep
- [ ] makefile, build: run test
- [ ] makefile, dev-deployment on minishift: docker build, kedge temple, push to minishift. One central repo with link:https://github.com/xcoulon/fabric8-minishift[Xavier's minishift repo]
- [ ] makefile, ops-deployment on OS: the real stuff with secret, CI/CD integration
- [ ] explore operator (required OS v3.9+)