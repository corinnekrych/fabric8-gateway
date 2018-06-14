# 3scale API

All API documentation is available at [YOUR_PORTAL_URL/p/admin/api_docs](https://redhatopenshiftio-admin.3scale.net/p/admin/api_docs)

## CREATE service
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

## Import definition via 3scale-cli
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

## CREATE application_plan

```
curl -v  -X POST "https://redhatopenshiftio-admin.3scale.net/admin/api/services/2555417758668/application_plans.xml" -d 'access_token=XXX&name=otherbasic2&system_name=otherbasic2&state_event=publish'
```

```
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

```
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


## CREATE application
```
curl -v  -X POST "https://corinne-admin.3scale.net/admin/api/accounts/2445582058137/applications.xml" -d 'access_token=XXX&plan_id=2357355951501&name=corinneapp&description=description'
```

403

```
x-runtime: 0.196748
date: Thu, 14 Jun 2018 14:31:02 GMT
content-encoding: gzip
x-content-type-options: nosniff
x-frame-options: DENY
connection: keep-alive
content-type: application/xml; charset=utf-8
cache-control: no-cache
strict-transport-security: max-age=15552000
vary: Accept-Encoding
content-length: 122
x-xss-protection: 1; mode=block
x-request-id: 784faef0-8e16-4eb6-ab05-34142495183f
x-served-by: mt05
```


