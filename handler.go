package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func readNoteHandler(c *gin.Context) {
	notes, err := database.GetAllNotes()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	fmt.Println(notes)
	c.JSON(http.StatusOK, notes)
}

func insertNoteHandler(c *gin.Context) {
	if err := database.SaveNote(c.Param("value")); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	readNoteHandler(c)
}

func deleteNoteHandler(c *gin.Context) {
	if err := database.DeleteNote(c.Param("value")); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	readNodeHandler(c)
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, database.GetHealthStatus())
}

func whoAmIHandler(c *gin.Context) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	addresses, err := getAllAddresses(ifaces)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": appVersion,
	})
}
