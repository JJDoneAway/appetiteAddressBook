# addressbook frontend
----------------------
This is a small vue.js example application which integrates all enterprise needs:

1. Communicate via REST with an API
2. Has routs
3. Has components
3. Has general services exposed to extra files
4. Uses the composition (`setup`) feature to cluster the code
5. Uses SIAM OAuth / OPIC / JWT / PKCE for authentication and authorization 
6. Uses Run time environment injection to deliver the code to the cloud


## Project setup
----------------
```
pnpm install
```

### Compiles and hot-reloads for development
```
pnpm run dev 
```
open http://localhost:8081

### Build and test for production
```
npm run build:ci
docker build -t frontend .
docker run --rm -p 8081:8080 --name frontend -e ODJ_DEP_ADDRESS_BOOK_API_URI=https://address-book-demo-app-aka-pareto-app.api.prod.sit.sys.odj.cloud  frontend 
```
open http://localhost:8081

### Build and run on ODJ
* see an up and running version: https://address-book-frontend.prod.sit.svc.odj.cloud/
* Have a look into the pipeline configuration [odj-azure-pipeline.yml](odj-azure-pipeline.yml)
* Have a look into the Sonarqube configuration [sonar-project.properties](sonar-project.properties)
* Always check that you don't have any serious reported issues from Snyk and Sonarqube 
* To manage the injected ODJ-runtime dependencies I used https://github.com/JJDoneAway/env_injector (JS applications aren't able to manage runtime dependencies)
* Mind that you enabled Snyk and Sonarqube in your Dev/Ops dependencies

## Tips to learn and understand Vue.JS / OAuth / OpenID / JWT
-------------------------------------------------------------
Beside a lot of reading I can definitely recommend the following Pluralsight courses:
* Learning Vue JS the funny way: https://app.pluralsight.com/library/courses/vuejs-fundamentals/table-of-contents
* Understand OAuth: https://app.pluralsight.com/library/courses/oauth-2-getting-started/table-of-contents
* Understand JWT: https://app.pluralsight.com/library/courses/jwt-fundamentals/table-of-contents

## Enable Security
------------------
### general hints
* As this is a SPA (single page application) I implemented the whole OAuth dance in the client. This enables you to deliver the application even via a CDN with no server backend. This approach takes the load from the backend and uses the resources of the client
* The implemented OAuth flow is ["Authorization Code Grand Type"](https://oauth.net/2/grant-types/authorization-code/) with [PKCE](https://oauth.net/2/pkce/) this is state of the art [good and short description](https://auth0.com/docs/get-started/authentication-and-authorization-flow/authorization-code-flow-with-proof-key-for-code-exchange-pkce)
* I used [oidc-client-ts](https://github.com/authts/oidc-client-ts) as implementation of the client side part of the dance  
* read for the callback https://damienbod.com/2019/01/29/securing-a-vue-js-app-using-openid-connect-code-flow-with-pkce-and-identityserver4/
* Don't forget to check (via your browsers developer tools) that the code challenge is generated and send to the OAuth tenant and that it is crypted (code_challenge and code_very must not be the same).     
* Mind that SAIM forces to use the  client secret. So it is not secret any more. But this is no problem as we use PKCE

### Order SIAM application
* The SIAM client can be ordered via IT4YOU service catalog (search for "SIAM OAuth/OIDC (OpenID Connect) Client on SIAM IDP"). find a good description about what to do https://itdoc.schwarz/pages/viewpage.action?pageId=1717032895#expand-HowToArtikel
* The SIAM application (To create new roles) can be ordered via IT4YOU service catalog ("search for "SIAM Initial Application Connection (Consulting and Application Container Creation)")

### Order SIAM Role management
To be able to manage roles in SIAM order "SIAM Initial Application Connection (Consulting and Application Container Creation)" via IT4YOU self service

We created the following roles to illustrate the role management

address-book-enterprise-example-xx-sit-change
address-book-enterprise-example-xx-sit-admin



## Where to find what in the code
---------------------------------
* Routing [index.js](./src/router/index.js)
* Check if user has logged in [App.vue](./src/App.vue)
* Code and configuration of the OAuth dance [auth.service.js](./src/components/services/auth.service.js)
* The Address Editor component [AddressEditor](./src/components/AddressEditor.vue)
* The main page which uses the composition API [AddressBook](./src/components/AddressBook.vue) 
* The environment page [Environment](./src/components/CheckEnvironment.vue)
* The API calls [api package](./src/components/api/)

