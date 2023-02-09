### edit config.json or create new
``` 
{
    "rds": [
        {
            "AccessKeyId":"xxx",
            "AccessKeySecret": "xxx",
            "RegionId": "xxx",
            "InstanceIds": [
                "xxx",
                "yyy"
            ],
            "Excluded": [
                "xxx",
                "yyy"
            ]
        }
    ],
    "polardb": [
        {
            "AccessKeyId":"xxx",
            "AccessKeySecret": "xxx",
            "RegionId": "xxx",
            "InstanceIds": [
                "xxx",
                "yyy"
            ],
            "Excluded": [
                "xxx",
                "yyy"
            ]
        }
    ]
}
```

### run 
```
./aliyun-slow-sql -config /path/to/file
```
