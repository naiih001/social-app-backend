package controllers

import "github.com/gin-gonic/gin"

// GET api/v1/admin/users
// Get all users (for admin)
func GetUsersByAdminPrivilege(c *gin.Context) {}

// DELETE api/v1/admin/posts/:id
// Delete a post (for admin)
func DeletePostByAdminPrivilege(c *gin.Context) {}

// POST api/v1/admin/ban/:id
// Ban a user (for admin)
func BanUserByAdminPrivilege(c *gin.Context) {}
