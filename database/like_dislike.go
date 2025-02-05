package database

import (
	"fmt"
)

// Like adds a like for a post or comment and removes any existing dislike for the same post or comment.
func LikePost(userID int, postID string) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer a rollback in case of failure
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Remove any existing dislike for the same post or comment
	dislikeQuery := `DELETE FROM dislikes WHERE user_id = ? AND post_id = ?`
	_, err = tx.Exec(dislikeQuery, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to remove dislike: %w", err)
	}

	// Insert the like
	likeQuery := `INSERT INTO likes (user_id, post_id) VALUES (?, ?)`
	_, err = tx.Exec(likeQuery, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to insert post like: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func LikeComment(userID int, commentID string) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer a rollback in case of failure
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Remove any existing dislike for the same post or comment
	dislikeQuery := `DELETE FROM dislikes WHERE user_id = ? AND comment_id = ?`
	_, err = tx.Exec(dislikeQuery, userID, commentID)
	if err != nil {
		return fmt.Errorf("failed to remove dislike: %w", err)
	}

	// Insert the like
	likeQuery := `INSERT INTO likes (user_id, comment_id) VALUES (?, ?)`
	_, err = tx.Exec(likeQuery, userID, commentID)
	if err != nil {
		return fmt.Errorf("failed to insert comment like: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Dislike adds a dislike for a post or comment and removes any existing like for the same post or comment.
func DislikePost(userID int, postID string) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer a rollback in case of failure
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Remove any existing like for the same post or comment
	likeQuery := `DELETE FROM likes WHERE user_id = ? AND post_id = ?`
	_, err = tx.Exec(likeQuery, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to remove like: %w", err)
	}

	// Insert the dislike
	dislikeQuery := `INSERT INTO dislikes (user_id, post_id) VALUES (?, ?)`
	_, err = tx.Exec(dislikeQuery, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to insert post dislike: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func DislikeComment(userID int, commentID string) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer a rollback in case of failure
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Remove any existing like for the same post or comment
	likeQuery := `DELETE FROM likes WHERE user_id = ? AND comment_id = ?`
	_, err = tx.Exec(likeQuery, userID, commentID)
	if err != nil {
		return fmt.Errorf("failed to remove like: %w", err)
	}

	// Insert the dislike
	dislikeQuery := `INSERT INTO dislikes (user_id, comment_id) VALUES (?, ?)`
	_, err = tx.Exec(dislikeQuery, userID, commentID)
	if err != nil {
		return fmt.Errorf("failed to insert post dislike: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
