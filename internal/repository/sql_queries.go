package repository

const (
	queryCreateUser = `
		INSERT INTO users(
			name,
			email,
			password,
			role,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	queryGetUserByEmail = `
		SELECT 
		    id,
		    name,
		    email,
		    password,
		    role,
		    created_at,
		    updated_at
		FROM users
		WHERE email = $1`

	queryGetUserByID = `
		SELECT
    		id,
    		name,
    		email,
    		password,
    		role,
    		created_at,
    		updated_at
		FROM users
		WHERE id = $1`

	queryCreatePost = `
		INSERT INTO posts(
			title,
			text,
		    created_at,
		    updated_at,
		    user_id
		) VALUES ($1, $2, $3, $4, $5)
	  	RETURNING id`

	queryGetPostByID = `
		SELECT 
		    id,
		    title,
		    text,
		    created_at,
		    updated_at,
		    user_id
		FROM posts
		WHERE id = $1`

	queryGetPostByTitle = `
		SELECT 
		    id,
		    title,
		    text,
		    created_at,
		    updated_at,
		    user_id
		FROM posts
		WHERE title = $1`

	queryGetPostByUserID = `
		SELECT 
		    id,
		    title,
		    text,
		    created_at,
		    updated_at,
		    user_id
		FROM posts
		WHERE user_id = $1`

	queryUpdatePost = `
		UPDATE posts
		SET title = $1,
		    text = $2,
		    updated_at = $3
		WHERE id = $4`

	queryCreateImage = `
		INSERT INTO images(
			url,
		   	post_id
		) VALUES ($1, $2)
	  	RETURNING id`

	queryGetImageByPostID = `
		SELECT url
		FROM images
		WHERE post_id = $1`
)
