package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	Cpf  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos []Aluno
