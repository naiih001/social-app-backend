package controllers

import "github.com/gin-gonic/gin"

// POST /api/v1/posts/:id/like
// Like a post
func CreateLike(c *gin.Context) {}

// DELETE /api/v1/posts/:id/like
// Delete a like for a post
func DeleteLike(c *gin.Context) {}

// GET /api/v1/users/:id/lilkes
// Get number of likes on a post
func GetLikes(c *gin.Context) {}
