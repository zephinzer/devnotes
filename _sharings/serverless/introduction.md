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

<span style="font-size:1.5em">monolith</span> > services > <span style="font-size: 0.5em">services</span> > <span style="font-size:1.5em">f( *x* )</span>

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

<!-- .slide: data-align="left" -->

### the good

- really small deployments
- language/framework independent
- usage-based scaling

------

- lower risk of breaking things
- easily disposable code - no legacy
- onboard based on capability
- ease of outsourcing
- lower costs
- `t2.medium` - 42.048/month
- let's see the [Lambda Pricing Calculator](https://s3.amazonaws.com/lambda-tools/pricing-calculator.html)

---

### the bad

- difficult to debug
- complex integration tests
- relatively immature technologies
- lack of resources

---

### the useful

- web hooks
- scheduled tasks
- data extract/transform/load
- file uploads & pre/post-processing
- event triggered actions

------

- content delivery
- payment processing
- file antivirus scans
- data ingress/sync jobs
- notification systems

---

[methodology and tools](./methodology-and-tools.md)