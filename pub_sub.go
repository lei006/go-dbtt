package godbtt

import "github.com/tidwall/buntdb"

type Subscriber struct {
	ID       string
	Channels []string // 存储订阅的频道路径
}

func (dbtt *DBTT) Subscribe(db *buntdb.DB, subscriberID string, channel string) error {
	return db.Update(func(tx *buntdb.Tx) error {
		// 检查订阅者是否存在，如果不存在则创建
		var sub Subscriber
		if _, err := tx.Get("subscriber:"+subscriberID, &sub); err != nil {
			sub = Subscriber{ID: subscriberID}
		}
		// 添加新的频道到订阅者列表
		sub.Channels = append(sub.Channels, channel)
		// 将更新后的订阅者信息存回数据库
		_, _, err := tx.Set("subscriber:"+subscriberID, sub, nil)
		return err
	})
}

func (dbtt *DBTT) Unsubscribe(db *buntdb.DB, subscriberID string, channel string) error {
	return db.Update(func(tx *buntdb.Tx) error {
		// 获取订阅者信息
		var sub Subscriber
		if _, err := tx.Get("subscriber:"+subscriberID, &sub); err != nil {
			return err
		}
		// 从订阅者列表中移除频道
		for i, c := range sub.Channels {
			if c == channel {
				sub.Channels = append(sub.Channels[:i], sub.Channels[i+1:]...)
				break
			}
		}
		// 更新订阅者信息
		_, _, err := tx.Set("subscriber:"+subscriberID, sub, nil)
		return err

	})
}

func (dbtt *DBTT) Publish(db *buntdb.DB, channel string, data interface{}) error {

	return db.View(func(tx *buntdb.Tx) error {
		// 遍历所有订阅者
		iter := tx.Ascend("subscriber:", func(key, value string) bool {
			var sub Subscriber
			if _, err := tx.Get(key, &sub); err != nil {
				return true
			}
			// 检查订阅者是否订阅了这个频道
			for _, c := range sub.Channels {
				if c == channel {
					// 发送数据给订阅者
					// 注意：此处需要实现具体的发送逻辑
				}
			}
			return true
		})
		return iter
	})
}
