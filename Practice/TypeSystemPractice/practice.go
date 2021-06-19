package main

import "fmt"

func main() {
	// 定义一个新的类型

	// 默认值为新类型里属性对应的默认值
	var bill user
	fmt.Println(bill)

	//类型创建并初始化
	newUser := user{
		name:       "smith",
		email:      "smith@example.com",
		ext:        0,
		privileged: true,
	}
	fmt.Println(newUser)

	//类型创建并初始化的另外一种方式,需要注意值得顺序
	newUser2 := user{"smith", "smith@example.com", 0, false}
	fmt.Println(newUser2)

	// 类型嵌套
	type admin struct {
		persion user
		level   string
	}

	newUser3 := admin{
		persion: user{
			name:       "smith",
			email:      "smith@example.com",
			ext:        0,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println(newUser3)

	//给user添加新的行为
	fmt.Printf("person, addrss is %p \n", &newUser)
	newUser.changeName("smith1")
	// fmt.Println(newUser)

	newUser.changeEmail("smith1@example.com") //(&newUser).changeEmail("smith1@example.com"),两者是等价的
	// fmt.Println(newUser)

	fmt.Printf("person, addrss is %p \n", &newUser)

	// (&newUser).changeEmail("smith2@example.com")
	// fmt.Println(newUser)
	fmt.Println("----------")

	newUser4 := &user{"smith", "smith@example.com", 0, false}
	fmt.Printf("person, addrss is %p \n", newUser4)
	newUser4.changeName("123")
	// fmt.Println(newUser4)
	newUser4.changeEmail("test@test")
	fmt.Printf("person, addrss is %p \n", newUser4)

}

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func (u user) changeName(name string) {
	fmt.Printf("internal, addrss is %p \n", &u)
	u.name = name
}

func (u *user) changeEmail(email string) {
	fmt.Printf("internal, addrss is %p \n", u)
	u.email = email
}
