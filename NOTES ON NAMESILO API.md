# Notes on namesilo's API

## request format

https://www.namesilo.com/api/OPERATION?version=VERSION&type=TYPE&key=YOURAPIKEY

required fields:
- OPERATION: To be replaced by the name of the specific operation you would like to execute.
- VERSION: To be replaced by the API version you would like to use. The current version is "1".
- TYPE: To be replaced by the format you would like to receive returned. The only current option is "xml".
- YOURAPIKEY: To be replaced by your unique API key. Visit the API Manager page within your account for details.

## response format

All processed requests return HTTP Status Code 200.  Must check code in XML for success/failure/error.

```xml
<namesilo>
    <request>
        <operation>OPERATION</operation> 
        <ip>YOUR IP</ip> 
    </request>
    <reply>
        <code>RESPONSE CODE</code> 
        <detail>RESPONSE DETAIL</detail>            
    </reply>
</namesilo>
```

## registerDomain operation
e.g. `/api/registerDomain?version=1&type=xml&key=12345&domain=namesilo.com&years=2&private=1&auto_renew=1`

additional required fields:
- domain
- years

various optional fields.

## getDomainInfo operation
e.g. `/api/getDomainInfo?version=1&type=xml&key=12345&domain=namesilo.com`

additional required fields:
- domain

## listDomains operation
e.g. `/api/getDomainInfo?version=1&type=xml&key=12345&domain=namesilo.com`

optional fields:
- portfolio

