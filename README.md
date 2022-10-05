# ChangeSet Markdown Converter

> Convert AWS CloudFormation ChangeSet JSON to a MarkDown table.

``` sh
$ aws cloudformation describe-change-set --change-set-name arn:aws:cloudformation:us-east-1:12334343:changeSet/fooo-qeb056uck4q-lvp65fhtmi/f26ee5d8-fcc3-4e80-b96a-c89d1219b01b > changeset.json
$ changesetmd changeset.json > changeset.md
```

* This tool also accepts the JSON format available in the "JSON Changes" tab.
