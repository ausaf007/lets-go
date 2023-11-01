
This is the planned directory structure for the generated backend server:

```
/my-awesome-project
	/api
		# business-logic
		/handlers
			user.go
		/middlewares
			logging.go
		/routes
	/cmd
		# starter
		/main.go
	/db
		# db-related-everything
		/migrations
			schema.sql
        /datastore
            # data models related everything
            # contains all the generated files by sqlc
            db.go
            models.go
            user.sql.go
		/queries
			user.sql
		sqlc.yaml
	/docs
	/pkg
		/configs
		/utils
	.env.sample
	go.mod
	go.sum
```

Note: In this representation of directory structure, directory name starts with `\`, 
files are written with just the filename and format, 
and the directory structure is organized as per indentation. Also lines starting with # indicates a comment.
