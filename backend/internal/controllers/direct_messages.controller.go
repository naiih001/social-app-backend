package controllers

import "github.com/gin-gonic/gin"

// POST api/v1/messages
// Create a new direct message
func CreateDirectMessage(c *gin.Context) {}

// GET api/v1/messages/conversations
// Get all the posts
func GetConversations(c *gin.Context) {}

// GET api/v1/messages/conversations/:id
// Get a single, specific conversation
func GetSingleConversations(c *gin.Context) {}
