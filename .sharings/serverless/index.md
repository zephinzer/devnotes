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

### considerations for serverless

------

##### features

really small deployments

natural separation of concerns

language/runtime independent

scale/deploy based on need

------

##### benefits

low risk deployments

onboard based on capability

ease of contributions

lower costs

------

##### risks

requires structuring



------

##### use cases

##### event driven solutions

data extract/transform/load

payment gateways

notification bots

business rule engines

---

### development

------

##### sync vs async

------

##### hot functions

---

### testing

------

#### contract testing

---

### current tools

------

#### cloud solutions
- [AWS Lambda](https://aws.amazon.com/lambda/)
- [Google Cloud Functions](https://cloud.google.com/functions/)
- [Azure Functions](https://azure.microsoft.com/en-us/services/functions/)
- [WebTask](https://webtask.io/)

------

#### hosted solutions
- [kubeless](https://github.com/kubeless/kubeless)
- [OpenWhisk](https://github.com/apache/incubator-openwhisk)
- [FnProject](https://github.com/fnproject/fn)

------

#### frameworks
- [Serverless](https://github.com/serverless/serverless)

---

is it time for an example?

[yes it is](./example.md)
