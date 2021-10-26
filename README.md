# argocd-vs-fluxcd-application

An application that persists a random JSON object in the Redis, then reads and writes it to the STDOUT

### How to start

```shell
$ docker run redis
```

```shell
$ docker run \
  -e REDIS_URL=172.17.0.2:6379 \
  swr.eu-de.otc.t-systems.com/images/argocd-vs-fluxcd-application:latest
```

### Description

A short description of the application

- Go as a programming language
- Redis as a dependency
- [**GitHub Actions**](https://docs.github.com/en/actions) for the CI part of the CI/CD pipeline
- [**OTC Cloud Container Engine**](https://open-telekom-cloud.com/en/products-services/cloud-container-engine)

[Git Repository](https://github.com/iits-consulting/argocd-vs-fluxcd-infrastructure) with Infrastructure and CD
automation

### Purposes

**OTC Event:** [_A battle of GitOps tools: Argo CD vs FluxCD_](https://community.open-telekom-cloud.com/community?id=community_event&sys_id=8a84320fb7763450d15aa7b16b8c0222)
