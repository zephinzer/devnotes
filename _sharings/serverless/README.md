# Technical Sharing on Serverless Architecture
Visual aids and content for Agile Consulting & Engineering technical sharing session on 27th April 2018.

## Get Started
### Clone This Repository
Clone this repository, navigate to this directory and install the dependencies:

```sh
npm i
```

You can start the slides by running:

```sh
npm start
```

## Quick Commands
### Start Fn Server
```sh
fn start
```


### Start Fn UI

```sh
docker run --rm -it --link fnserver:api -p 4000:4000 -e "FN_API_URL=http://api:8080" fnproject/ui
```

## Licensing
These slides are licensed under [Creative Commons Attribution-ShareAlike 3.0 Singapore](https://creativecommons.org/licenses/by-sa/3.0/sg/). Full legal speak is [available here](https://creativecommons.org/licenses/by-sa/3.0/sg/legalcode).