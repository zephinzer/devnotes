<!-- .slide: data-background="./res/cloud-outline.jpg" -->
<!-- .slide: data-background-color="white" -->
### ~~server~~

------

- joseph (@zephinzer)
- my careers future
- dev/devops

---

### ~~server~~?

------

![./res/there-is-no-cloud.png](./res/there-is-no-cloud.png)
#### serverless !== no servers

------

backend-as-a-service

<img src="./res/firebase.png" style="height:100px;border:none;">

functions-as-a-service

<img src="./res/azure-functions.png" style="height:100px;border:none;">
<img src="./res/aws-lambda.png" style="height:100px;border:none;">
<img src="./res/cloud-functions.png" style="height:100px;border:none;">


------

~~backend-as-a-service~~

**functions-as-a-service**

---

### where we came from

------

##### monolithic
*deploy entire product*  

------

<img src="./res/architecture-monolith.png" style="height: 500px" />

------

##### service oriented
*deploy applications*  

------

<img src="./res/architecture-service-oriented.png" style="height: 500px" />

------

##### microservices
*deploy self-contained services*  

------

<img src="./res/architecture-microservices.png" style="height: 500px" />

------

##### serverless
*deploy functions*

------

<img src="./res/architecture-serverless.png" style="height: 500px" />

---

### the good

- really small deployments
- language/framework independent
- usage-based scaling

------

##### really small deployments

- lower risk of breaking things

------

##### language/framework independent

- no lock in
- easily disposable code
- onboard based on capability

------

##### scale based on usage

- lower costs
- pay only for what is needed
- `t2.medium`
  - 0.0584/hour
  - 1.4016/day
  - 42.048/month

---

### the bad

- 

------

### use cases

- web hooks
- scheduled tasks
- data extract/transform/load
- 