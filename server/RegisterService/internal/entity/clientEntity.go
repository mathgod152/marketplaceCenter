package entity

import "errors"

type Client struct{
	ID string
	ClintName string
	ClientLastName string
	Username string
	Password string
	ClientPrincipalAdress string
}

func NewClient(id, clintName, ClientLastName, username, password, clientPrincipalAdress string) (*Client, error){
	client := &Client{
		ID: id,
		ClintName: clintName,
		ClientLastName: ClientLastName,
		Username: username,
		Password: password,
		ClientPrincipalAdress: clientPrincipalAdress,
	}
	if err:= Client.ValidateClient(client); err != nil{
		return nil, err
	}
	
	return client, nil
}

func (c *Client)ValidateClient() error{
	if Client.ID == ""{
		return errors.New("Id invalido")
	}
}