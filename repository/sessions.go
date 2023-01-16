package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	// Select target token and delete from listSessions
	// TODO: answer here
	list := model.Session{}

	fmt.Sprintln(list)
	for i, session1 := range listSessions {
		if (session1.Token != tokenTarget) {
			continue
		}else{
			listSessions[i] = list
		}
	}

	jsonData, err := json.Marshal(listSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	data, Alert := u.ReadSessions()

	data = append(data, session)
	database, Alert := json.Marshal(data)

	db.NewJsonDB().Save("sessions", database)
	if (Alert != nil){
		Alert := model.ErrorResponse{
			Error: "Data Error",
		}
		fmt.Sprintln(Alert)
		return nil
	}else{
		
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	session := model.Session{}

	data, Alert := u.TokenExist(token)

	if (Alert != nil ){
		return session, nil // TODO: replace this
	}else{

	}
	if (!u.TokenExpired(data)){
		return data, nil // TODO: replace this
	}else{
		return session, fmt.Errorf("Token is Expired!") // TODO: replace this
	}
	// return model.Session{}, nil  TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
