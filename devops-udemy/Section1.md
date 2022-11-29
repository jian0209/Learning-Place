## Start up IntelliJ in linux

```
cd /opt
ls
### copy intellij name
cd intellij name/bin
./idea.sh
```

---

## Connect with ubuntu linux

```
ssh jian@192.168.64.5
```

---

## Process

SDLC

```
Requirement
> Planning
> > Designing
> > > Development
> > > > Testing
> > > > > Deploy & Maintain
```

---

## Continuous Integration

### Version Control System

- Ex: Github

  The place that everyone push and pull the code

- Artifact of the flow

  Build > Tested > Evaluated

### Process

```
Code
> Fetch
> > Build
> > > Test
> > > > Notify
> > > > > Feedback
```

### Build Tools

```
Maven, Ant, Gradle
MSBuild, Visual Build
IBM Urban Code
Make
Grunt
```

### Software Repository

```
Sonatype Nexus
Jfrog Artifactory
Archiva
Cloudsmith Package
Grunt
```

### CI Tools

```
Jenkins
CircleCi
TeamCity
Bamboo CI
Cruise Control
```

### Conclusion

- Code <------------
- Build & Test         |
- Code Analysis     |
- Artifact Repo ----- (done then going on)
- Deploy Automation <-
- DB / Sec Changes     |
- Software Testing------ (done then going on)
- Schedule Prod Deployment

## Continuous Delivery

### Deployment

```
Server Provisioning
Dependencies
Conf Changes
Network
Artifact Deploy
```

### Tools

```
Ansible, Pupper, Chef
Terraform, Cformation
Jenkins, Octopus Deploy
Helm Charts
Code Deploy
```

### Test Automation

```
Functional
Load
Performance
DB
Security
```

### Conclusion

- Code
- Code Build
- Code Test
- Code Analysis
- Delivery
- DB/Sec Changes
- Software Testing
- Schedule Prod Deployment

