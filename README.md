# Go Project
```bash
go init webservice
```

# DOCKER mysql-container
```bash
docker run -d --name mysql-container -p 3306:3306 -v /var/lib/mysql:/var/lib/mysql -e "MYSQL_ROOT_PASSWORD=mysqlpw" mysql
```

# MySql Go Driver
```bash
go get -u github.com/go-sql-driver/mysql
```

# Sql ip for using out WSL
172.21.93.208 mysql

```js
let ws = new WebSocket("ws://localhost:5000/websocket")
ws.send(JSON.stringify({data: "test message from browser", type: "test"}))
```