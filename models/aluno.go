package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF  int    `json:"cpf"`
	RG   string `json:"rg"`
}
