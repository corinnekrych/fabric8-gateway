# 3scale API

All API documentation is available at [YOUR_PORTAL_URL/p/admin/api_docs](https://redhatopenshiftio-admin.3scale.net/p/admin/api_docs)

## Step 1: CREATE service
```
curl -v  -X POST "https://redhatopenshiftio-admin.3scale.net/admin/api/services.xml" -d 'access_token=XXX&name=fabric8-toggles-test&deployment_option=hosted&backend_version=1&system_name=fabric8-toggles-test'
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<service>
  <id>2555417758668</id>
  <account_id>2445582058137</account_id>
  <name>fabric8-toggles-test</name>
  <state>incomplete</state>
  <system_name>fabric8-toggles-test</system_name>
  <backend_version>1</backend_version>
  <end_user_registration_required>true</end_user_registration_required>
  <metrics>
    <metric>
      <id>2555418108730</id>
      <name>hits</name>
      <system_name>hits</system_name>
      <friendly_name>Hits</friendly_name>
      <service_id>2555417758668</service_id>
      <description>Number of API hits</description>
      <unit>hit</unit>
    </metric>
  </metrics>
</service>
```

## LIST all service (optional)
```
curl -v  -X GET "https://redhatopenshiftio-admin.3scale.net/admin/api/services.xml?access_token=XXX"
```
List all the services you're able to see depending of your admin rights.

## Step 2: Import definition via 3scale-cli
```
3scale-cli import swagger -f ./swagger/swagger.yaml -s 2555417758668
```

According to [2scale-cli readme](), when you import an API definition, the following actions will be performed in the background:
* Create a new service (unless you specify one) ID for API calls is 2555417758668 and system name is fabric8-toggles-test
* Create methods in the 'Definition' section
* Attach newly created methods to the 'Hits' metric
* Create mapping rules and show them under API > Integration

TODO: retranscrire those cli step into API calls

## READ service (optional)
```
curl -v  -X GET "https://redhatopenshiftio-admin.3scale.net/admin/api/services/2555417758668.xml?access_token=XXX"
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<service>
  <id>2555417758668</id>
  <account_id>2445582058137</account_id>
  <name>fabric8-toggles-test</name>
  <state>incomplete</state>
  <system_name>fabric8-toggles-test</system_name>
  <backend_version>1</backend_version>
  <end_user_registration_required>true</end_user_registration_required>
  <metrics>
    <metric>
      <id>2555418108730</id>
      <name>hits</name>
      <system_name>hits</system_name>
      <friendly_name>Hits</friendly_name>
      <service_id>2555417758668</service_id>
      <description>Number of API hits</description>
      <unit>hit</unit>
    </metric>
    <method>
      <id>2555418108738</id>
      <name>api_features_GET</name>
      <system_name>api_features_GET</system_name>
      <friendly_name>api_features_GET</friendly_name>
      <service_id>2555417758668</service_id>
      <description/>
      <metric_id>2555418108730</metric_id>
    </method>
    <method>
      <id>2555418108739</id>
      <name>api_features__featureName__GET</name>
      <system_name>api_features__featureName__GET</system_name>
      <friendly_name>api_features__featureName__GET</friendly_name>
      <service_id>2555417758668</service_id>
      <description/>
      <metric_id>2555418108730</metric_id>
    </method>
    <method>
      <id>2555418108740</id>
      <name>api_status_GET</name>
      <system_name>api_status_GET</system_name>
      <friendly_name>api_status_GET</friendly_name>
      <service_id>2555417758668</service_id>
      <description/>
      <metric_id>2555418108730</metric_id>
    </method>
  </metrics>
```

## Step 3: CREATE application_plan

```
curl -v  -X POST "https://redhatopenshiftio-admin.3scale.net/admin/api/services/2555417758668/application_plans.xml" -d 'access_token=XXX&name=otherbasic2&system_name=otherbasic2&state_event=publish'
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<plan custom="false" default="false">
  <id>2357355951501</id>
  <name>otherbasic2</name>
  <type>application_plan</type>
  <state>published</state>
  <service_id>2555417758668</service_id>
  <end_user_required>false</end_user_required>
  <setup_fee>0.0</setup_fee>
  <cost_per_month>0.0</cost_per_month>
  <trial_period_days>0</trial_period_days>
  <cancellation_period>0</cancellation_period>
</plan>
```

## LIST application_plan per service (optional)

```
curl -v  -X GET "https://redhatopenshiftio-admin.3scale.net/admin/api/services/2555417758668/application_plans.xml?access_token=XXX"
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<plans>
  <plan custom="false" default="false">
    <id>2357355951443</id>
    <name>Basic</name>
    <type>application_plan</type>
    <state>published</state>
    <service_id>2555417758668</service_id>
    <end_user_required>false</end_user_required>
    <setup_fee>0.0</setup_fee>
    <cost_per_month>0.0</cost_per_month>
    <trial_period_days/>
    <cancellation_period>0</cancellation_period>
  </plan>
</plans>
```


## Step 4: CREATE application
```
curl -v  -X POST "https://redhatopenshiftio-admin.3scale.net/admin/api/accounts/2445582556479/applications.xml" -d 'access_token=XXX&plan_id=2357355951501&name=appf8test&description=appf8test'
```

> NOTE: __do not used the account_id returned by the create service. The developer_id should be known (param of the script)__.
>The account_id returned from the service create is the one that the service belongs to (The tenant)
> - Service belongs to a Tenant (account), for all the services created under the same domain it will always return the same `account_id`
> - Developers (account) can subscribe to a Service by creating an application.
> In the UI you can go to the developer tab and that’s where you can pick the account_id
> For automatic scripting you need to know for which developer you want to create an application
> You can create one Developer if you do not have one from the API endpoint `Signup Express (Account create)` and use the `id` as the `account_id`


```xml
<?xml version="1.0" encoding="UTF-8"?>
<application>
  <id>1409617773838</id>
  <created_at>2018-06-15T07:10:49Z</created_at>
  <updated_at>2018-06-15T07:10:49Z</updated_at>
  <state>live</state>
  <user_account_id>2445582556479</user_account_id>
  <first_traffic_at/>
  <first_daily_traffic_at/>
  <end_user_required>false</end_user_required>
  <service_id>2555417758668</service_id>
  <user_key>USER_KEY</user_key>
  <provider_verification_key>PROVIDER_KY</provider_verification_key>
  <plan custom="false" default="false">
    <id>2357355951501</id>
    <name>otherbasic2</name>
    <type>application_plan</type>
    <state>published</state>
    <service_id>2555417758668</service_id>
    <end_user_required>false</end_user_required>
    <setup_fee>0.0</setup_fee>
    <cost_per_month>0.0</cost_per_month>
    <trial_period_days/>
    <cancellation_period>0</cancellation_period>
  </plan>
  <name>appf8test</name>
  <description>appf8test</description>
  <extra_fields></extra_fields>
</application>
```

## Step 5: UPDATE proxy
Update the proxy with private URL etc...
```
curl -v  -X PATCH "https://redhatopenshiftio-admin.3scale.net/admin/api/services/2555417758668/proxy.xml" -d 'access_token=XXX&api_backend=https%3A%2F%2FPRIVATE_URL.com%3A443'
```

```xml
<proxy>
  <service_id>2555417758668</service_id>
  <endpoint>https://fabric8-toggles-test-2445582058137.production.gw.apicast.io:443</endpoint>
  <api_backend>https://YYY.com:443</api_backend>
  <credentials_location>headers</credentials_location>
  <auth_app_key>app_key</auth_app_key>
  <auth_app_id>app_id</auth_app_id>
  <auth_user_key>user-key</auth_user_key>
  <error_auth_failed>Authentication failed</error_auth_failed>
  <error_auth_missing>Authentication parameters missing</error_auth_missing>
  <error_status_auth_failed>403</error_status_auth_failed>
  <error_headers_auth_failed>text/plain; charset=us-ascii</error_headers_auth_failed>
  <error_status_auth_missing>403</error_status_auth_missing>
  <error_headers_auth_missing>text/plain; charset=us-ascii</error_headers_auth_missing>
  <error_no_match>No Mapping Rule matched</error_no_match>
  <error_status_no_match>404</error_status_no_match>
  <error_headers_no_match>text/plain; charset=us-ascii</error_headers_no_match>
  <secret_token>Shared_secret_sent_from_proxy_to_API_backend_8c5c0bcfe22158a8</secret_token>
  <hostname_rewrite/>
  <sandbox_endpoint>https://fabric8-toggles-test-2445582058137.staging.gw.apicast.io:443</sandbox_endpoint>
  <api_test_path>/api/features</api_test_path>
  <api_test_success>false</api_test_success>
  <policies_config>[{"name"=&gt;"apicast", "version"=&gt;"builtin", "configuration"=&gt;{}, "enabled"=&gt;true}]</policies_config>
  <created_at>2018-06-14 08:11:45 UTC</created_at>
  <updated_at>2018-06-14 09:38:56 UTC</updated_at>
  <lock_version>8</lock_version>
</proxy>
```

## READ proxy (optional)
```
curl -v  -X GET "https://redhatopenshiftio-admin.3scale.net/admin/api/services/2555417758668/proxy.xml?access_token=XXX"
```

response idem than UPDATE proxy.

## Step 6: Proxy Config promote
All the step we did applied to staging environment.
Promote v5 to production for ex:

```
curl -v  -X POST "https://redhatopenshiftio-admin.3scale.net/admin/api/services/2555417758668/proxy/configs/sandbox/5/promote.json" -d 'access_token=XXX&to=production'
```

```xml
{
 "proxy_config": {
  "id": 40183,
  "version": 5,
  "environment": "production",
  "content": {
   "id": 2555417758668,
   "account_id": 2445582058137,
   "name": "fabric8-toggles-test",
   "oneline_description": null,
   "description": null,
   "txt_api": null,
   "txt_support": null,
   "txt_features": null,
   "created_at": "2018-06-14T08:11:44Z",
   "updated_at": "2018-06-14T09:38:56Z",
   "logo_file_name": null,
   "logo_content_type": null,
   "logo_file_size": null,
   "state": "incomplete",
   "intentions_required": false,
   "draft_name": "",
   "infobar": null,
   "terms": null,
   "display_provider_keys": false,
   "tech_support_email": null,
   "admin_support_email": null,
   "credit_card_support_email": null,
   "buyers_manage_apps": true,
   "buyers_manage_keys": true,
   "custom_keys_enabled": true,
   "buyer_plan_change_permission": "request",
   "buyer_can_select_plan": false,
   "notification_settings": null,
   "default_application_plan_id": null,
   "default_service_plan_id": null,
   "default_end_user_plan_id": null,
   "end_user_registration_required": true,
   "tenant_id": 2445582058137,
   "system_name": "fabric8-toggles-test",
   "backend_version": "1",
   "mandatory_app_key": true,
   "buyer_key_regenerate_enabled": true,
   "support_email": "devtools-sre@redhat.com",
   "referrer_filters_required": false,
   "deployment_option": "hosted",
   "proxiable?": true,
   "backend_authentication_type": "service_token",
   "backend_authentication_value": "XXX",
   "proxy": {
    "id": 104922,
    "tenant_id": 2445582058137,
    "service_id": 2555417758668,
    "endpoint": "https://fabric8-toggles-test-2445582058137.production.gw.apicast.io:443",
    "deployed_at": null,
    "api_backend": "https://PRIVATE_URL.com:443",
    "auth_app_key": "app_key",
    "auth_app_id": "app_id",
    "auth_user_key": "user-key",
    "credentials_location": "headers",
    "error_auth_failed": "Authentication failed",
    "error_auth_missing": "Authentication parameters missing",
    "created_at": "2018-06-14T08:11:45Z",
    "updated_at": "2018-06-14T09:38:56Z",
    "error_status_auth_failed": 403,
    "error_headers_auth_failed": "text/plain; charset=us-ascii",
    "error_status_auth_missing": 403,
    "error_headers_auth_missing": "text/plain; charset=us-ascii",
    "error_no_match": "No Mapping Rule matched",
    "error_status_no_match": 404,
    "error_headers_no_match": "text/plain; charset=us-ascii",
    "secret_token": "Shared_secret_sent_from_proxy_to_API_backend_8c5c0bcfe22158a8",
    "hostname_rewrite": "",
    "oauth_login_url": null,
    "sandbox_endpoint": "https://fabric8-toggles-test-2445582058137.staging.gw.apicast.io:443",
    "api_test_path": "/api/features",
    "api_test_success": false,
    "apicast_configuration_driven": true,
    "oidc_issuer_endpoint": null,
    "lock_version": 8,
    "authentication_method": "1",
    "hostname_rewrite_for_sandbox": "PRIVATE_URL.com",
    "endpoint_port": 443,
    "valid?": true,
    "service_backend_version": "1",
    "hosts": [
     "fabric8-toggles-test-2445582058137.production.gw.apicast.io",
     "fabric8-toggles-test-2445582058137.staging.gw.apicast.io"
    ],
    "backend": {
     "endpoint": "https://su1.3scale.net",
     "host": "su1.3scale.net"
    },
    "policy_chain": [
     {
      "name": "apicast",
      "version": "builtin",
      "configuration": {}
     }
    ],
    "proxy_rules": [
     {
      "id": 288145,
      "proxy_id": 104922,
      "http_method": "GET",
      "pattern": "/",
      "metric_id": 2555418108730,
      "metric_system_name": "hits",
      "delta": 1,
      "tenant_id": 2445582058137,
      "created_at": "2018-06-14T08:11:45Z",
      "updated_at": "2018-06-14T08:11:45Z",
      "redirect_url": null,
      "parameters": [],
      "querystring_parameters": {}
     },
     {
      "id": 288155,
      "proxy_id": 104922,
      "http_method": "GET",
      "pattern": "/api/features",
      "metric_id": 2555418108738,
      "metric_system_name": "api_features_GET",
      "delta": 1,
      "tenant_id": 2445582058137,
      "created_at": "2018-06-14T09:21:05Z",
      "updated_at": "2018-06-14T09:21:05Z",
      "redirect_url": null,
      "parameters": [],
      "querystring_parameters": {}
     },
     {
      "id": 288156,
      "proxy_id": 104922,
      "http_method": "GET",
      "pattern": "/api/features/{featureName}",
      "metric_id": 2555418108739,
      "metric_system_name": "api_features__featureName__GET",
      "delta": 1,
      "tenant_id": 2445582058137,
      "created_at": "2018-06-14T09:21:07Z",
      "updated_at": "2018-06-14T09:21:07Z",
      "redirect_url": null,
      "parameters": [
       "featureName"
      ],
      "querystring_parameters": {}
     },
     {
      "id": 288157,
      "proxy_id": 104922,
      "http_method": "GET",
      "pattern": "/api/status",
      "metric_id": 2555418108740,
      "metric_system_name": "api_status_GET",
      "delta": 1,
      "tenant_id": 2445582058137,
      "created_at": "2018-06-14T09:21:09Z",
      "updated_at": "2018-06-14T09:21:09Z",
      "redirect_url": null,
      "parameters": [],
      "querystring_parameters": {}
     }
    ]
   }
  }
 }
}
```
