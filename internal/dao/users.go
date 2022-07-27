package dao

import (
	"context"
	"fmt"
	model "helloword/internal/model"
	"log"
)

// AddUser mysql添加用户
func (d *dao) AddUser(data *model.Users) (int64, error) {

	data = new(model.Users)
	data.Name = "kelly"
	//data.Age = 18
	//使用db.exec手动插入数据
	res, err := d.db.Exec(context.TODO(), "insert into users (name, age) values (?, ?)", data.Name, 20)
	if err != nil {
		return 0, fmt.Errorf("新增姓名 %s 出错，错误为 %v", data.Name, err)
	}
	id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("返回 id 出错，错误为 %v", err)
	}
	return id, nil

}

//SearchUser map查询一条数据
func (d *dao) SearchUser(data *model.Users) (map[string]string, error) {
	//查询
	data = new(model.Users)
	data.Name = "tom"
	var uid, name, age, ctime, mtime string
	err := d.db.QueryRow(context.TODO(), "select * from users where name= ?", data.Name).Scan(&uid, &name, &age, &ctime, &mtime)
	if err != nil {
		fmt.Printf("返回 %s", err)
	}

	fmt.Println("err:", err)
	//fmt.Printf("返回：name:%s, age:%s, ctime:%s,mtime:%s", name, age, ctime, mtime)
	s2 := map[string]string{"uid:": uid, "name": name, "age": age, "ctime": ctime, "mtime": mtime}
	return s2, nil
}

//SearchStructUser 结构体查询
func (d *dao) SearchStructUser(data *model.Users) (*model.Users, error) {
	data.Name = "zhangkunling"
	fmt.Println(data.Age, data.Name)

	fmt.Println("-----", data) //接收service下的方法
	err := d.db.QueryRow(context.TODO(), "select * from users where name= ?", data.Name).Scan(&data.Uid, &data.Name, &data.Age, &data.Ctime, &data.Mtime)
	if err != nil {
		fmt.Printf("返回 %s", err)
		return nil, err
	}

	fmt.Println("data:", data)

	return data, nil
}

//SearchStruct 多条查询，返回的是一条数据
func (d *dao) SearchStruct(data *model.Users) (*model.Users, error) {
	//data.Name = "zhangkunling"
	//data.Age = 18
	rows, err := d.db.Query(context.TODO(), "select * from users")
	if err != nil {
		fmt.Printf("返回 %s", err)
		log.Fatalln(err)
		//return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&data.Uid, &data.Name, &data.Age, &data.Ctime, &data.Mtime)
		if err != nil {
			log.Fatal("err", err)
		}
		log.Println("log:", data.Uid, data.Name, data.Age, data.Ctime, data.Mtime)
		fmt.Println("rowlog:", data)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("data:", data)
	return data, nil
}

//UpdateUser 修改数据
func (d *dao) UpdateUser(data *model.Users) (int64, error) {
	data = new(model.Users)
	data.Name = "kun"
	res, err := d.db.Exec(context.TODO(), "update users set name = ? where uid = ? ", data.Name, 9)
	if err != nil {
		fmt.Println("update failed err:", err)
	} else {
		fmt.Println("update success!")
	}
	id, err := res.LastInsertId()
	return id, nil
}

//DeleteUser 删除数据
func (d *dao) DeleteUser(data *model.Users) (int64, error) {
	data = new(model.Users)
	data.Name = "zhangkunling"
	res, err := d.db.Exec(context.TODO(), "delete from users where name = ? and uid=8", data.Name)
	if err != nil {
		fmt.Println("delete err:", err)
	} else {
		fmt.Println("delete success")
	}

	id, err := res.LastInsertId()
	return id, nil
}