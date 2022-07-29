package service

import (
	"fmt"
	"helloword/internal/model"
)

//新增user业务逻辑
func (s *Service) InsertUser() (id int64, err error) {
	id, err = s.dao.AddUser(nil)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//func (s *Service) NewInsertUser() (users string, err error) {
//	data := new(model.Users)
//	users, err = s.dao.NewAddUser(data)
//	if err != nil {
//		return
//	}
//	return users, nil
//}

func (s *Service) SearchUser() (users map[string]string, err error) {
	users, err = s.dao.SearchUser(nil)
	if err != nil {
		return users, nil
	}
	fmt.Println("users:", users)
	return users, nil
}

func (s *Service) SearchStructUser() (users *model.Users, err error) {
	data := new(model.Users) //分配了内存空间
	//data.Age = 18
	//data.Name = "aaa"
	users, err = s.dao.SearchStructUser(data)
	if err != nil {
		return users, nil
	}
	fmt.Println("users:", users)
	return users, nil
}

func (s *Service) SearchStruct() (users []*model.Users, err error) {
	users, err = s.dao.SearchStruct()
	if err != nil {
		return users, nil
	}
	return users, nil
}

func (s *Service) UpdateUser() (id int64, err error) {
	id, err = s.dao.UpdateUser(nil)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *Service) DeleteUser() (id int64, err error) {
	id, err = s.dao.DeleteUser(nil)
	if err != nil {
		return 0, err
	}
	return id, nil
}
