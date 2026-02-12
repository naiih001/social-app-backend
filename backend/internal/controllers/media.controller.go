package controllers

import "github.com/gin-gonic/gin"

// POST api/v1/media/upload
// Upload a new media file (image, sound or video)
func UploadMedia(c *gin.Context) {}

// POST  api/v1/media/confirm
// Confirm a video has been uploaded
func ConfirmUpload(c *gin.Context) {}

// GET api/v1/media/:id
// Get a media file (image, sound or video)
func GetUpload(c *gin.Context) {}
