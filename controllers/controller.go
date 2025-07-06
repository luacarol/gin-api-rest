package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luacarol/api-go-gin/database"
	"github.com/luacarol/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	// Aqui você pode implementar a lógica para exibir uma saudação
	c.JSON(200, gin.H{
		"message": "Olá, seja bem-vindo " + nome + "!",
	})

}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorID(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	if err := database.DB.Where("id = ?", id).First(&aluno).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	if err := database.DB.Where("id = ?", id).First(&aluno).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	database.DB.Delete(&aluno)
	c.JSON(http.StatusOK, gin.H{"message": "Aluno deletado com sucesso"})
}

func EditaAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	if err := database.DB.Where("id = ?", id).First(&aluno).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var aluno models.Aluno
	if err := database.DB.Where("cpf = ?", cpf).First(&aluno).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
