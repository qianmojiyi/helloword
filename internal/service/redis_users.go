package service

import "helloword/internal/model"

//RedisUser 添加
func (s *Service) RedisUser() (users *model.Users, err error) {
	data := new(model.Users)
	users, err = s.dao.RedisUser(data)
	if err != nil {
		return
	}
	return users, nil
}

//RedisAdd redis业务逻辑
func (s *Service) RedisAdd() (users string, err error) {
	data := new(model.Users)
	users, err = s.dao.RedisAdd(data)
	if err != nil {
		return
	}
	return users, nil
}

//RedisDel
func (s *Service) RedisDel() (users string, err error) {
	data := new(model.Users)
	users, err = s.dao.RedisDel(data)
	if err != nil {
		return
	}
	return users, nil
}

//RedisGet
func (s *Service) RedisGet() (users string, err error) {
	data := new(model.Users)
	users, err = s.dao.RedisGet(data)
	if err != nil {
		return
	}
	return users, nil
}

//NewRedisGet
func (s *Service) NewRedisGet() (users *model.Users, err error) {
	data := new(model.Users)
	users, err = s.dao.NewRedisGet(data)
	if err != nil {
		return
	}
	return users, nil
}
