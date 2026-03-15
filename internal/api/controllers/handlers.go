package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers de usuarios

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, newUser)
}

func UpdateProfile(c *gin.Context) {
	// Handler para actualizar el perfil del usuario
}

func GetProfile(c *gin.Context) {
	// Handler para obtener el perfil del usuario
}

func ChangePassword(c *gin.Context) {
	// Handler para actualizar la contraseña
}

func DeleteUser(c *gin.Context) {
	// Handler para borrar usuario
}

// Handlers de login y JWT
func LoginUser(c *gin.Context) {
	//Handler para mandarle al usuario los JWT
}

func RefreshToken(c *gin.Context) {
	//Handler para mandarle al usuario los JWT
}

func LogoutUser(c *gin.Context) {
	//Handler para mandarle al usuario los JWT
}

// Handlers de resources
func SearchAndAddResource(c *gin.Context) {
	// Handler para crear una serie de manga/anime
}

func GetResources(c *gin.Context) {
	// Handler para obtener una serie de manga/anime
}

func GetResourcesByGenre(c *gin.Context) {
	// Handler para obtener una serie de manga/anime
}

func GetMostPopular(c *gin.Context) {
	// Handler para obtener una serie de manga/anime
}

func GetTopRated(c *gin.Context) {
	// Handler para obtener una serie de manga/anime
}

func UpdateResource(c *gin.Context) {
	// Handler para actualizar una serie de manga/anime
}

func DeleteResource(c *gin.Context) {
	// Handler para borrar recurso
}

// Handlers para listas
func CreateList(c *gin.Context) {
	// Handler para crear una lista de manga/anime
}

func GetUserLists(c *gin.Context) {
	// Handler para obtener la lista de un usuario
}

func AddToListByListID(c *gin.Context) {
	// Handler para agregar anime/manga a la lista
}

func UpdateList(c *gin.Context) {
	// Handler para actualizar una lista de manga/anime
}

func UpdateListEntryByListIDAndEntryID(c *gin.Context) {
	// Handler para agregar estado/progreso/etc
}

func DeleteList(c *gin.Context) {
	// Handler para borrar una lista
}

func DeleteListEntryByListIDAndEntryID(c *gin.Context) {
	// Handler para eliminar list entry
}

// Handlers para puntuaciones
func CreateScore(c *gin.Context) {
	// Handler para crear una puntuación de un manga/anime
}

func GetScore(c *gin.Context) {
	// Handler para obtener una puntuación de un manga/anime
}

func UpdateScore(c *gin.Context) {
	// Handler para actualizar una puntuación de un manga/anime
}

func DeleteScore(c *gin.Context) {
	// Handler para borrar una puntuación de un manga/anime
}

// Handlers para reseñas
func CreateReview(c *gin.Context) {
	// Handler para crear una reseña de un manga/anime
}

func GetReview(c *gin.Context) {
	// Handler para obtener una reseña de un manga/anime
}

func UpdateReview(c *gin.Context) {
	// Handler para actualizar una reseña de un manga/anime
}

func DeleteReview(c *gin.Context) {
	// Handler para borrar una reseña de un manga/anime
}
