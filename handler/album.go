package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GCU-Sharpic/sharpic-server/types/album"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAlbumList(c *gin.Context) {
	cookie, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	albumList, err := h.dbClient.FindAlbumListByUsername(cookie)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"albumList": albumList})
}

func (h *Handler) GetAlbum(c *gin.Context) {
	param := c.Param("albumId")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	album, err := h.dbClient.FindAlbumByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, album)
}

func (h *Handler) PostNewAlbum(c *gin.Context) {
	newAlbum, err := album.NewShouldBindJSON(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newAlbum.Username, err = c.Cookie("username")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": "error"})
		c.Abort()
		return
	}

	if err := h.dbClient.InsertNewAlbum(newAlbum); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "new album success"})
}
