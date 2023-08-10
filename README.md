# OCS_job


# db command
sqlite3 wordcards.db 


# 代码优化
1. 要求用的govalidator， 否则直接在tag里面写validator就行了
2. 管理员可以撤销JWT，模拟用户可能需要强制退出的情况。这个部分用了全局变量，只是为了方便，正常项目肯定用缓存
3. 处理多个用户同时更新单词卡的情况，可以考虑用锁，但是我更偏向于将更新请求放入消息队列中，然后由后台的异步任务来处理更新操作。
4. 大量读取问题我会将单词卡的数据放入缓存中，并设置过期时间。缓存没有情况才去数据库查找，这样可以让访问量大的单词从缓存获取。


# 启动方式
1. make up
2. make shell
3. make run

# 测试
1. make shell
2. cd handlers
3. go test 
