# ROLE BASE ACCESS CONTROL

## Description

### Role

- Role applies to a name space
- Rules that defiles what are allowed and what are not allowed to perform.
- Components of Role:

  - Api Groups

    - Groups holds resources.
    - eg. Pods are in core api group.
    - eg. Deployment and replica set falls under app api group.

  - Resource
    - Action are performed on resources.
    - Pods, services etc.
  - Verbs
    - Verbs are the actual action that are performed on the cluster resources.
    - eg. get, delete, update, watch

- After the role creation we need to bind that role to a particular user.
- We can create a group while creating certificates and bind the role to it, so any user in that group will have that role.(helps in case of many users).

### Cluster Role

- Cluster role is applicable to full cluster

## Setps

- Create certificate

  ```
  openssl genrsa -out <user_name>.key 2048
  ```

  ```
  eg.
  -----BEGIN RSA PRIVATE KEY-----
  MIIEowIBAAKCAQEAlVcdUjBcrOH5nNU2OVeefVZ22PiIY0N9PJhYPg7p5byOoxcS
  Lhgde65gpMm82+N6CSxFKVS83CPwjfouN0bg/YxAuozFRbHaLb96p5vU65fVcQfl
  kxiHcdNzjsbWa7pRWeMbe4Op23IhqCXZDGHpvyOo4+zMKC2nrHanif7IGS1P5HqS
  InMRIfcsMzBeugwd/+w0xeBWHIuKN9F7m+6YzLR1AsupgEA536ty3hiiQMfHK2a1
  IAI9PJNIZhwnLoQfBJjkpAjjLKvt0UKklN4yqlh2vyUuNPphWXpZOJ4QusoywEx8
  QLHbSUnu8sEf82DJ1M9Wpv6mHKv80KVD0uxHBQIDAQABAoIBACmDf+ka/IgqgSLx
  uhMSokIhSQRjgjZlymiPyOfnaWmtktEPVsRte7Sn8sGcOt29TG1EWy10GtPvQa2W
  AEbpcx3IYS36d9Q57naNJeF11ihvmqvDEDRluzEPIK5t46kqr/aWxUr+fLc98NE3
  RDDG/x1dIlPJva7lkhjFGymnI6SroTZVGxVdEUWwAn9k/UrUDNS8D/9TFbraTbZO
  jG/7zXBXQQIa0SkdKGrBFia78gl0JIo0aenUgVFckmb5BJAiwYxeyr53ZYrPyKge
  hIIpMOvOPz2Xb8QmfRq4tzBaSF/cC1F1kqgvwm1y6XY/YnBqWnsuAjUvWtKG/QAa
  NBZ7yzkCgYEAxJdbT5nWPxo1btAAKVmZBp78Qp3+K9EvofZdsKNBIlGoH7X3FxPW
  if3sz2+oWeYVNeChIItVj/6sJhT7tA1E/dtawTR8xRWrMo2Qa/La+l/rqYCm5+Qj
  i33Yw6TMkSv7hlIPrDOMLVqdEYC/GutEyAEqQ5DHl6jGDajr5P+G4SMCgYEAwnhX
  eeuQnUSyTtzNRcOU9iYL9JTT5tE/EIUt8ZtllKmUK1mJLdcJA/+gVFlb89qRi5w0
  OXbEEj/ti4uBzwrF3WoaJkjfyIKRDaq8wF2yH9qx5RtTVzZOLP3dToPrOzweh15K
  PqwgRloR+H5HCJqMcF3akLH5Vv2x8sPo3RZ0PbcCgYAsSmFTDLidHN1+6rOunTYB
  rN3ucEPsQUO2/JddDVrGTZlQE6HofBqwmhm5HlHSdyi+K66Bc4LCX0EPj1ip1vd9
  LedDcfzONjSMviIUf/FfyS6B0K9JvTdY3PByG3XPGxaey2wmgrKJIUoTwGrNuMWV
  PdX0xPCCtY5xU2ZaAEu1+wKBgA0neLtNfNW2fN0qwFPUu/G+dxaPLpoxUzO53na2
  TTO2GMip4MaSAhEcLGsoNDkVcrzfzwSIDREy8815nUk08FdewvAOA90bP9aN0wCs
  Fc1jzJ6zr3p5uOYAMgUKlVHHakU70ofUusYnd4m4ePOTCKbdIlihQT1qIaxLJig2
  SZyXAoGBALxYOjCof+aiVpP5/qA6yJwzjIotXLTXr6FPSjQbcaLXRicGadYPy2nN
  9DlrdfcqeUC4V0P+UjwXaCQaGIKR6QFdSEhhUPJ69zIj5fn71L/ewE6CP63PZ9ZZ
  /r1oGlNcl/eSaKcWF5GuE0vJ7YEOCLktjC9HzJkX+IACVb9SZVeT
  -----END RSA PRIVATE KEY-----
  ```

- Create certificate singing request.
  ```
  openssl req -new -key <user_name>.key -out <user_name>.csr -subj "/CN=<user_name>/O=<namespace>
  ```
- Now we need `ca.crt` and `ca.key` file to sign the certificate for the user.
- These files can be found in kubernets master node under `/etc/kubernetes/pki/` dir.
- Login to the `kubernetes master node` and navigate to the dir and copy the files.

  ```
  scp root@172.16.16.100:/etc/kubernetes/pki/ca.{crt,key}`.
  ```

- Sign the certificate.
  ```
  openssl x509 -req -in <user_name>.csr -CA ca.crt -CAkey =CAcredateserial -out <user_name>.crt -days <expiration date on days>
  ```
- Now we need to create the kube.config file.

<!-- ```

kubectl --kubeconfig john.kubeconfig config set-cluster kubernetes --server https://172.16.16.100:6443 --certificate-authority=ca.crt

```

- Aftet this we need to add user to this kubeconfig file.

```

kubectl --kubeconfig john.kubeconfig config set-credentials john --client-certificate ./john.crt --client-key ./john.key

```

- Now we need to set the context.

```

kubectl --kubeconfig john.kubeconfig config set-context john-k8s --cluster kubernetes --namespace default --user john
``

```

edit the john.kubeconfig file and add context there

copy the config file to .kube/config -->

<!-- Another way -->

- Make a copy of the original config file and update it accordingly.
- Encode the <user_name>.crt file using base64

  ```
  cat <user_name>.crt | base64 -w0
  ```

- Copy the encoded value and paste in the `client-certificate-data` field in the config file.
- Repeat above steps for the `<user_name>.key` file also.

  ```
    cat john.crt | base64 -w0
  ```

- Copy the encoded value and paste in the `client-key-data` field in the config file.

- Now we need to create a role.

  ```

  kubectl create role <user_name>-default --verb=get,list --resource=pods --namespace <namespace_name>

  Or

  we can crete a role.yaml file and deploy it.

  ```

- We need to attach the role to user john, it id done by role binding.

  ```

  kubectl create rolebinding john-finance-rolebinding --role=<user_name>-default --user=<user_name>   --namespace <namespace_name>

  Or

  we can create it by using deployment file.

  ```

- GIve this config file the user.
- User can paste the config file in `~/.kube/config` and access the cluster resources as defined in the role.
