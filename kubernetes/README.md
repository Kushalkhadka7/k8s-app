# Deploy to kubernetes

## Deploy applications

- Deploy admin app ([Admin](#Admin))
- Deploy auth app ([Auth](#Auth))
- Deploy manager app ([Manager](#Manager))

After all these apps are deployed, the apps will be accessibel only if the port is forwarded, service is exposed to the outside world.

To access the application from outside `2` methods can be implemented

- [NginxIngress](#NginxIngress)
- [Istio](#Istio) [Along with istio ingress]

## Nginx Ingress
