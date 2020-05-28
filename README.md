# tfplan-converter
tf plan converter to json for terraform v11

to convert terraform plan v11 from `terraform.InstanceDiff` to json.
```
$ ./tf-converter --tfplan ./test/tfplan-2
{
    "aws_s3_bucket.b1": {
        "destroy": true,
        "destroy_tainted": false
    },
    "aws_s3_bucket.b2": {
        "acceleration_status": "",
        "acl": "private",
        "arn": "",
        "bucket": "tf-plan-converter-test",
        "bucket_domain_name": "",
        "bucket_regional_domain_name": "",
        "destroy": false,
        "destroy_tainted": false,
        "force_destroy": "false",
        "hosted_zone_id": "",
        "id": "",
        "region": "",
        "request_payer": "",
        "tags.%": "2",
        "tags.Environment": "Dev",
        "tags.Name": "My bucket",
        "versioning.#": "",
        "website_domain": "",
        "website_endpoint": ""
    },
    "destroy": false
}
```

# distribution and cross compile 
for linux:
```
env GOOS=linux GOARCH=arm64 go build . 

```

for MacOs:
```
env GOOS=darwin GOARCH=amd64 go build . 

```

for windows:
```
env GOOS=windows GOARCH=amd64 go build .
```
for more operating systems check here:
https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

# Annotations
Thanks to [palantir's tfjson project](https://github.com/palantir/tfjson)