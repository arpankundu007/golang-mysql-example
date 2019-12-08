# golang-mysql-example
Golang application which performs basic database queries on MySQL DB

GET("/mobile/mobiles/:id")
Fetches information of a `mobile` with `id`

GET("/mobile/all")
Fetches information of all mobiles in the database

POST("/mobile/add")
Creates a new mobile record

PUT("/mobile/update/:id)
Updates a row with new mobile data belonging to `id`

GET("/drop")
Drops the table. (Currently a server restart is required after a drop)

DELETE("/delete/{id}")
Delete record belonging to `id`
