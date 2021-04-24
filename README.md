# Go Project
go init webservice

# DOCKER mysql-container
docker run -d --name mysql-container -p 3306:3306 -v /var/lib/mysql:/var/lib/mysql -e "MYSQL_ROOT_PASSWORD=mysqlpw" mysql

# MySql Go Driver
go get -u github.com/go-sql-driver/mysql

# Sql ip for using out WSL
172.21.93.208 mysql