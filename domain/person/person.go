package person

import (
	"encoding/json"
	"fmt"
	"github.com/maicongavino/api-crud-go/domain"
	"io/ioutil"
	"os"
)

type Service struct {
	dbFilePath string
	peope      domain.People
}

func NewService(dbFilePath string) (Service, error) {
	//verificar se o arquivo exite
	_, err := os.Stat(dbFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			//se não existir, crio arquivo vazio
			err = createEmptyFile(dbFilePath)
			if err != nil {
				return Service{}, err
			}
		}
		return Service{
			dbFilePath: dbFilePath,
			peope:      domain.People{},
		}, nil
	}

	//se existir, leio o arquivo e atualizo a variavel people do serviço com as pessoas do arquivo
	jsonFile, err := os.Open(dbFilePath)
	if err != nil {
		return Service{}, fmt.Errorf("Error trying to open file that  contains all people: %s", err.Error())
	}
	jsonFileContentByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Service{}, fmt.Errorf("Error trying to read file: %s", err.Error())
	}

	var allPeople domain.People
	json.Unmarshal(jsonFileContentByte, &allPeople)

	return Service{
		dbFilePath: dbFilePath,
		peope:      domain.People{},
	}, nil

}

//se não existir, crio arquivo vazio
func createEmptyFile(dbFilePhat string) error {
	var people domain.People = domain.People{
		Peole: []domain.Person{},
	}
	peopleJSON, err := json.Marshal(people)
	if err != nil {
		return fmt.Errorf("Error trying to encode people as JSON: %s", err.Error())
	}
	err = ioutil.WriteFile(dbFilePhat, peopleJSON, 0755)
	if err != nil {
		return fmt.Errorf("Error trying to write to file. Error: %s", err.Error())
	}
	return nil
}

//Metodo de Criação de Pessoas
func (s Service) Create(person domain.Person) error {
	//verificar se a pessoa existe
	//se já existe então retorno err
	if s.existsPerson(person) {
		return fmt.Errorf("Error trying to create person. There is a person with this ID already registerend")
	}

	//adiciono a pessoa no slice de pessoas
	s.peope.Peole = append(s.peope.Peole, person)
	//salvo o arquivo
	err := s.salveFile()
	if err != nil {
		return fmt.Errorf("Error trying save file in method Create. Error: %s", err.Error())
	}
	return nil
}

//Metodo de verificação de pessoas
func (s Service) existsPerson(person domain.Person) bool {
	for _, currentPerson := range s.peope.Peole {
		if currentPerson.ID == person.ID {
			return true
		}

	}
	return false
}

//Metodo de salvar
func (s Service) salveFile() error {
	allPeopleJSON, err := json.Marshal(s.peope)
	if err != nil {
		return fmt.Errorf("Error trying to encode people as json: %s", err.Error())
	}
	return ioutil.WriteFile(s.dbFilePath, allPeopleJSON, 0755)
}
