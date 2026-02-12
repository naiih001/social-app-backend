package controllers

import "github.com/gin-gonic/gin"

// GET api/v1/users/:id
func GetUserProfile(c *gin.Context) {}

// GET api/v1/users/username/:username
func GetUserResolveByUsername(c *gin.Context) {}

// PATCH api/v1/users/me
func UpdateUserProfile(c *gin.Context) {}

// DELETE api/v1/users/me
func DeactivateUserProfile(c *gin.Context) {}

// GET api/v1/users/m
func GetUserFollowers(c *gin.Context) {}

// GET api/v1/users/m
func GetUserFollowing(c *gin.Context) {}

// GET api/v1/users/m
func GetUserPosts(c *gin.Context) {}

// GET api/v1/users/m
func Search(c *gin.Context) {}
