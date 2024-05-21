# Pareto App - Address Book

>
> ***The Pareto App provides within minutes an full blown up and running tech stack for enterprise applications within SIT*** 
>

-------------------------

## Why in few words
* We believe that most of our enterprise applications have more or less the same tech stack (60% - 80%)
* We realize that it takes currently up to some weeks to setup this tech stack from scratch
* We know that there are many technical requirements to have a production ready product at SIT
* We face the situation that many of our colleagues have deep expert knowledge in their domain and ok'ish experience in developing code (T or 𝝅 shape)

## What in few words
* The Pareto App will be installed with all its components via ODJ within minutes.
* Afterwards it can be checked out by the development team so that they can immediately start with the development of their use cases.
* The coding of the Pareto App is as easy and as transparent as possible, so that it can be adopted quick from normal skilled developers (no magic code principle).

## How in few words
the Pareto App consists out of a frontend, a backend and a data base

### Frontend Components
> https://dev.azure.com/schwarzit-wiking/schwarzit.johannes-meets-stackit/_git/johannes-meets-stackit.monorepo?path=/address-book-frontend
* VueJS Single Page Application
* SIAM login with OAUth - PKCE flow 
* Router
* Deep Links (login safe) to dedicated addresses via id
* Components (composition)
* Tables with pagination, filtering and sorting
* Forms to edit with validation
* Store / State management (Pinia)
* REST communication with the backend
* Runtime environment injection for ODJ deployments

### Backend
> https://dev.azure.com/schwarzit-wiking/schwarzit.johannes-meets-stackit/_git/johannes-meets-stackit.monorepo?path=/address-book-backend
* Go coded
* JWT validation / user and role Management
* REST layer (GIN)
* OpenAPI exposure (SWAG)
* MyAPI deployment
* Entity relation management (GORM & paginate)
* Monitoring (gin-metrics)
* 4+1 exposure (magic monitoring)

### Deployment
> https://console.odj.cloud/prods/show?product=johannes-meets-stackit
* ODJ
* Azure DevOps
* STACKIT advanced mode
* PostgreSQL
* MyAPI
* magic monitoring