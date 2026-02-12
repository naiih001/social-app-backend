package controllers

import "github.com/gin-gonic/gin"

// GET api/v1/timeline/home
func GetFeed(c *gin.Context) {}

// GET api/v1/timeline/user/:id
func GetUserFeed(c *gin.Context) {}

// GET api/v1/timeline/trending
func GetTrendingFeed(c *gin.Context) {}

// GET api/v1/timeline/explore
func GetExploreFeed(c *gin.Context) {}
