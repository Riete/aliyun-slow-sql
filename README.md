### edit config.json or create new
``` 
{
    "tasks": [
        {
            "access_key_id":"access key",
            "access_key_secret": "access key secret",
            "region_id": "region id",
            "instances": [
                {
                    "instance_id": "instance-1",
                    "excluded_db": ["db-1", "db-22"],
                    "instance_type": "RDS or PolarDB"
                },
                {
                    "instance_id": "instance-2",
                    "excluded_db": ["db-1", "db-2"],
                    "instance_type": "RDS or PolarDB"
                }
            ],
            "ding_talk" : {
                "webhook": "https://xxxxx",
                "secret": "secret"
            }
        }
    ]
}
```

### run 
```
./aliyun-slow-sql -config /path/to/file
```
