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


## CREATE application
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


```
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


