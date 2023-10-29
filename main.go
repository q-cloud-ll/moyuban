package main

import "fmt"

// Observer 定义观察者接口
type Observer interface {
	Update(message string)
}

// Subject 定义主题接口
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(message string)
}

// MessageNotifier 具体主题实现
type MessageNotifier struct {
	observers []Observer
}

func (m *MessageNotifier) Register(observer Observer) {
	m.observers = append(m.observers, observer)
}

func (m *MessageNotifier) Unregister(observer Observer) {
	for i, o := range m.observers {
		if o == observer {
			m.observers = append(m.observers[:i], m.observers[i+1:]...)
			break
		}
	}
}

func (m *MessageNotifier) Notify(message string) {
	for _, observer := range m.observers {
		observer.Update(message)
	}
}

// User 具体观察者实现
type User struct {
	name string
}

func (u *User) Update(message string) {
	fmt.Printf("[%s] Received message: %s\n", u.name, message)
}

// 示例用法
func main() {
	notifier := &MessageNotifier{}

	// 创建观察者
	user1 := &User{name: "User1"}
	user2 := &User{name: "User2"}

	// 注册观察者
	notifier.Register(user1)
	notifier.Register(user2)

	// 发送通知
	notifier.Notify("Hello, World!")

	// 取消注册观察者
	notifier.Unregister(user2)

	// 再次发送通知
	notifier.Notify("Goodbye!")

}
