package database

import (
    "log"
	"forum/models"
	"forum/utils"
	"database/sql"
)

func GetUserByEmailOrUsername(email, username string) (models.User, error) {
	query := `SELECT id, username, email, password FROM users WHERE email = ? OR username = ? OR email = ? OR username = ?`
	var user models.User
	err := db.QueryRow(query, email, username, username, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func VerifyUser(email, password string) bool {
	// User email or username to get users' full credentials
	user, err := GetUserByEmailOrUsername(email, email)
	if err != nil {
		return false
	}

	// Compare provided password with the stored hashed password
	return utils.ComparePasswords(user.Password, password)
}

// GetUserbySessionID function
func GetUserbySessionID(UUID string) (models.User, error) {
    query := `SELECT id, username, email, bio, image, created_at FROM users WHERE session_id = ?`
    
    var user models.User
    var bio, image sql.NullString  // Use sql.NullString for nullable fields
    
    err := db.QueryRow(query, UUID).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &bio,      // Scan into NullString
        &image,    // Scan into NullString
        &user.CreatedAt,
    )
    
    if err != nil {
        log.Printf("Database error: %v\n", err)
        return models.User{}, err
    }

    // Convert NullString to string, using empty string if NULL
    if bio.Valid {
        user.Bio = bio.String
    }
    if image.Valid {
        user.Image = image.String
    }
    
    return user, nil
}

// Get all posts by a user
func GetUserPostsbyUserID(ID int)([]models.Post, error) {
	query := `SELECT user_id, title, content, media, created_at FROM posts WHERE user_id = ?`
    rows,err := db.Query(query,ID)
    if err != nil {
        log.Println("Error querying posts by user ID:", err)
        return nil, err
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        var media sql.NullString
        err:=rows.Scan( &post.UserID, &post.Title, &post.Content, &media, &post.CreatedAt)
        if err!=nil{
            log.Println("Error scanning post row:", err)
            return nil, err
        }
        if media.Valid{
            post.Media = media.String
        } else{
            post.Media = ""
        }
        posts=append(posts,post)
    }

    if err = rows.Err(); err != nil {
        log.Println("Error scanning posts:", err)
        return nil, err
    }    

    return posts,nil
}