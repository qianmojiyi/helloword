package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"go-common/library/cache/redis"
	"go-common/library/log"
	"helloword/internal/model"
)

//RedisUser 添加mysql同时加redis
func (d *dao) RedisUser(data *model.Users) (users *model.Users, err error) {
	data = new(model.Users)
	data.Name = "kun"
	res, err := d.db.Exec(context.TODO(), "insert into users (name, age) values (?, ?)", data.Name, 30)
	if err != nil {
		log.Error("db.exec(%s) error:", res, err)
	} else {
		id, err := res.LastInsertId() //获取数据库UID
		err = d.db.QueryRow(context.TODO(), "select * from users where name= ? and uid= ?", data.Name, id).Scan(&data.Uid, &data.Name, &data.Age, &data.Ctime, &data.Mtime)
		if err != nil {
			fmt.Printf("返回 %s", err)
			return nil, err
		} else {
			fmt.Println("获取到mysql数据：")
			bytes1, _ := json.Marshal(&data) //data结构体转换为切片，获取 {"uid":23,"name":"kun","age":30,"ctime":"2022-07-29T20:26:58+08:00","mtime":"2022-07-29T20:26:58+08:00"}
			fmt.Println("data:", data)
			fmt.Println("string(bytes1):", string(bytes1))
			result, err := redis.String(d.redis.Do(context.TODO(), "SET", data.Name, string(bytes1))) //将[]byte转换成string，并进行添加到redis中，key为uid。value是user
			if err != nil && err != redis.ErrNoReply {
				log.Error("d.redis.Do error(%v)", err)
				fmt.Println("result:", result)
				//return
			}
		}
	}
	return data, nil

}

//RedisAdd 添加数据到redis
func (d *dao) RedisAdd(data *model.Users) (users string, err error) {
	data.Name = "tom"
	data.Uid = 9
	_, err = d.redis.Do(context.TODO(), "SET", data.Uid, data.Name)
	//_, err = d.redis.Do(context.TODO(), "DEL", data.Uid, data.Name)

	// 错误处理
	if err != nil && err != redis.ErrNoReply {
		log.Error("d.redis.Do error(%v)", err)
		return
	}
	fmt.Println("data-uid:", data.Uid, data.Name)
	new1, err := redis.Int(d.redis.Do(context.TODO(), "EXISTS", data.Uid))
	if new1 != 0 {
		result, err := redis.String(d.redis.Do(context.TODO(), "GET", data.Uid))
		return result, err
	} else {
		println("error key不存在:", new1)
	}
	//result, err := redis.String(d.redis.Do(context.TODO(), "GET", data.Uid)) //将[]byte转换成string，并进行添加到redis中，key为uid。value是user
	//if err != nil && err != redis.ErrNoReply {
	//	log.Error("d.redis.Do error(%v)", err)
	//	return
	//}
	//d.redis.Do(context.TODO(), "EXPIRE", data.Uid, 20) //设置过期时间，到期自动清空value

	return
	//return result, nil

}

//RedisDel 删除数据
func (d *dao) RedisDel(data *model.Users) (users string, err error) {
	data.Name = "tom"
	data.Uid = 9

	new1, err := redis.Int(d.redis.Do(context.TODO(), "EXISTS", data.Uid))
	if new1 != 0 {
		//result, err := redis.String(d.redis.Do(context.TODO(), "GET", data.Uid))
		_, err = d.redis.Do(context.TODO(), "DEL", data.Uid) //, data.Name)
		fmt.Printf("uid: %d 删除成功", data.Uid)
		return users, nil
	} else {
		println(data.Uid, " value不存在,不需要删除")
	}
	//d.redis.Do(context.TODO(), "EXPIRE", data.Uid, 20) //设置过期时间，到期自动清空value

	return users, nil
	//return result, nil
}

//RedisGet 查询数据
func (d *dao) RedisGet(data *model.Users) (users string, err error) {
	//data.Name = "tom"
	data.Uid = 9

	fmt.Println("data-uid:", data.Uid, data.Name)
	new1, err := redis.Int(d.redis.Do(context.TODO(), "EXISTS", data.Uid))
	if new1 != 0 {
		result, err := redis.String(d.redis.Do(context.TODO(), "GET", data.Uid))
		return result, err
	} else {
		println("error key不存在:", new1)
	}

	//d.redis.Do(context.TODO(), "EXPIRE", data.Uid, 20) //设置过期时间，到期自动清空value

	return

}

//NewRedisGet 查询redis数据，没有从mysql中查找并添加
func (d *dao) NewRedisGet(data *model.Users) (users *model.Users, err error) {
	data.Name = "zkl"
	data.Uid = 6
	//redis数据是否存在
	new1, err := redis.Int(d.redis.Do(context.TODO(), "EXISTS", data.Uid))
	if new1 != 0 {
		result, err := redis.String(d.redis.Do(context.TODO(), "GET", data.Uid))
		fmt.Println("redis存在数据，result:", result)
		if err != nil && err != redis.ErrNoReply {
			log.Error("d.redis.Do error(%v)", err)
			return nil, err
		}
		json.Unmarshal([]byte(result), &data) //字符串变成data结构体
		return data, nil
	} else {
		println("error key不存在,去mysql查询:", new1)
		//mysql数据库查询
		//res, err := d.db.Exec(context.TODO(), "insert info users (name, age) values (?, ?)", data.Name, 25)
		err = d.db.QueryRow(context.TODO(), "select * from users where name= ?", data.Name).Scan(&data.Uid, &data.Name, &data.Age, &data.Ctime, &data.Mtime)
		if err != nil {
			log.Error("error:", err)
			//fmt.Printf("返回 %s", err)
		}
		bytes1, _ := json.Marshal(&data)                                                    //date结构体变成json字符串，bytes返回 {"uid":6,"name":"zkl","age":20,"ctime":"2022-07-23T18:46:52+08:00","mtime":"2022-07-27T14:39:49+08:00"}
		_, err := redis.String(d.redis.Do(context.TODO(), "SET", data.Uid, string(bytes1))) //将[]byte转换成string，并进行添加到redis中，key为uid。value是user等信息, _ = ok

		if err != nil && err != redis.ErrNoReply {
			log.Error("d.redis.Do error(%v)", err)
		}
		//d.redis.Do(context.TODO(), "EXPIRE", data.Uid, 20) //设置过期时间，到期自动清空value
	}
	return data, nil
}

//DeleteUser 删除数据,
func (d *dao) DeleteUser(data *model.Users) (int64, error) {
	data = new(model.Users)
	data.Name = "kelly"
	data.Uid = 11
	res, err := d.db.Exec(context.TODO(), "select * from users where name = ? and uid= ? ", data.Name, data.Uid)
	if err != nil {
		fmt.Println("aaa")
		return 0, err
	}
	res, err = d.db.Exec(context.TODO(), "delete from users where name = ? and uid=?", data.Name, data.Uid)
	if err != nil {
		fmt.Println("delete err:", err)
	} else {
		fmt.Println("delete success")
		//删除redis中的数据
		new1, _ := redis.Int(d.redis.Do(context.TODO(), "EXISTS", data.Uid))
		if new1 != 0 {
			fmt.Println("redis存在数据，准备删除 :", data.Uid)
			result, err := redis.Int(d.redis.Do(context.TODO(), "DEL", data.Uid))
			if err != nil {
				log.Error("删除数据失败：", result)
			}
		}
	}
	id, err := res.LastInsertId()
	return id, nil
}
