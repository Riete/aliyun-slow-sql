### docker build
``` docker build . -t <image>:<tag> ```

### or pull 
``` docker pull riet/aliyun-slow-sql ```

### run
```
docker run \ 
  -d \ 
  --name aliyun-slow-sql \
  riet/aliyun-slow-sql \
  --access.key.id=<aliyun ak>
  --access.key.secret=<aliyun ak sk>
  --region.id=<region id>
  --robot.url=https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxx
  --secret=SECxxxxxxxxxx
  --instance.ids=rds-id1,rds-id2,rds-id3
```
