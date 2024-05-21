package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JJDoneAway/addressbook/middleware"
	"github.com/JJDoneAway/addressbook/models"
	"github.com/gin-gonic/gin"
)

const ADDRESSES = "addresses"
const ID = "addresses/:id"

func AddAddressRouts(router *gin.Engine) {
	router.GET(ADDRESSES, doGetAll)
	router.POST(ADDRESSES, doPOST)
	router.DELETE(ADDRESSES, doDeleteAll)
	router.GET(ID, doGet)
	router.PUT(ID, doPut)
	router.DELETE(ID, doDelete)
	router.POST(ADDRESSES+"/default", doReset)
}

// @Summary      Pagination over all addresses
// @Description  Will paginate and filter the list of all currently known addresses
// @ID           getAllAddresses
// @Tags         addresses
// @Param        page  		query 	int 	false	"Page number starting from 0"		minimum(0)
// @Param        size  		query 	int 	false	"Max elements per page"				minimum(1)
// @Param		 sort		query	string	false	"Sort result by attribute asc or desc. E.g.: 'lastName' will order by last name asc. '-firstName,-lastName' will order by first name and last name desc"
// @Param		 filters	query	string	false	"Filter by value. E.g.: [[&quot;firstName&quot;,&quot;like&quot;,&quot;jo&quot;],[&quot;OR&quot;],[&quot;lastName&quot;,&quot;like&quot;,&quot;jo&quot;]] will filer all addresses with an 'jo' in the first or last name"
// @Produce      json
// @Success      200  		{object}   models.Page
// @Router       /addresses [get]
func doGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, models.PaginateAddresses(c.Request))
}

// @Summary      Add a new addresses
// @Description  Will add a new addresses entity to the storage. The new created addresses will be returned. Don't add the Id to the addresses parameter
// @ID           addAddress
// @Tags         addresses
// @Accept       json
// @Produce      json
// @Param        addresses body models.AddressChange true "The new addresses without ID"
// @Success      201  {object}  models.Address
// @Failure      400  {string}  string "ID must be zero, Unparsable JSON body"
// @Router       /addresses [post]
func doPOST(c *gin.Context) {
	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if address.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must not be set"})
		return
	}

	address.InsertAddress()

	c.JSON(http.StatusCreated, address)
}

// @Summary      Delete all addresses
// @Description  Will delete all addresses. It is not common to have an endpoint like this but we illustrate how to manage security for this kind of operations
// @ID           deleteAllAddresses
// @Tags         addresses
// @Produce      json
// @Success      204
// @Router       /addresses [delete]
func doDeleteAll(c *gin.Context) {
	if ok := models.DeleteAllAddresses(); ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ok})
		return
	}
	if ret, ok := models.GetAllAddresses(); ok != nil || len(ret) > 0 {
		log.Printf("It was not possible to delete all elements: %v", ok)
		c.JSON(http.StatusInternalServerError, "Something went wrong")
	} else {
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}

// @Summary      Get one address
// @Description  Get a address with the provided ID
// @ID           getOneAddress
// @Tags         addresses
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Success      200  {object}  models.Address
// @Failure      404  {string}  string "Unknown ID"
// @Router       /addresses/{id} [get]
func doGet(c *gin.Context) {
	ID := getID(c)
	address := (&models.Address{ID: ID}).GetAddressByID()
	if address == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unknown ID"})
		return
	}
	c.JSON(http.StatusOK, address)
}

// @Summary      Update an existing address
// @Description  Will update an existing address which is identified via its ID
// @ID           updateOneAddress
// @Tags         addresses
// @Accept       json
// @Produce      json
// @Param        id path integer true "ID of the address"
// @Param        user body models.AddressChange true "The new address without ID"
// @Success      200  {object}  models.Address
// @Failure      400  {string}  string "Unparsable JSON body"
// @Failure      404  {string}  string "Unknown ID"
// @Router       /addresses/{id} [put]
func doPut(c *gin.Context) {
	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	address.ID = getID(c)

	err := address.UpdateAddress()
	if err == models.ErrUnknownID {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("The address with the id '%d' is unknown. Maybe you mean a POST request", address.ID)})
		return
	}

	c.JSON(http.StatusOK, address)
}

// @Summary      Delete one address
// @Description  Delete a address with the provided ID
// @ID           deleteOneAddress
// @Tags         addresses
// @Produce      json
// @Param        id path integer true "ID of the address"
// @Success      204
// @Failure      400  {string}  string "Unknown ID"
// @Router       /addresses/{id} [delete]
func doDelete(c *gin.Context) {

	ID := getID(c)

	if err := (&models.Address{ID: ID}).DeleteAddressByID(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The address with the id '%d' is unknown.", ID)})
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)

}

// @Summary      Reset address book
// @Description  Will insert all initial dummy addresses to restore the initial state
// @ID           resetDB
// @Tags         addresses
// @Produce      json
// @Success      204
// @Router       /addresses/default [post]
func doReset(c *gin.Context) {
	middleware.AddDummies()
	c.Writer.WriteHeader(http.StatusNoContent)
}

////////////////
// Helpers    //
////////////////

// Helper function as I don't know how to do this in gin
func getID(c *gin.Context) uint {
	ret, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0
	}
	return uint(ret)
}
