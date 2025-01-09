/*
/api/admin/users		GET		List all users
/api/admin/users/:id	GET		Get a single user by ID
/api/admin/users		POST	Create a new user
/api/admin/users/:id	PUT		Update an existing user
/api/admin/users/:id	DELETE	Delete a user

		userGroup.GET("/", handlers.GetUsers)
		userGroup.GET("/:id", handlers.GetUserByID)
		userGroup.GET("/", handlers.CreateUser)
		userGroup.GET("/:id", handlers.UpdateUser)
		userGroup.GET("/:id", handlers.DeleteUser)
*/

package handlers
