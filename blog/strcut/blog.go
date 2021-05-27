package blog_strcut

import (
	"database/sql"
	"fmt"
)

type blog interface {
	setBlogId(id int32)
	setTitle(title string)
	setTopicGroup(topicGroup string)
	setCreateTime(createTime string)
	setUpdateTime(updateTime string)
	setContent(content string)
	setAuthor(author string)
}
type blogOperation interface {
	Add(db *sql.DB) (res interface{}, err error)
	Remove(db *sql.DB, id int32) (res interface{}, err error)
	Update(db *sql.DB) (res interface{}, err error)
	Get(db *sql.DB, id int32) (res interface{}, err error)
}
type Blog struct {
	blogId     int32  `json:"blog_id"`
	title      string `json:"title"`
	topicGroup string `json:"topic_group"`
	createTime string `json:"create_time"`
	updateTime string `json:"update_time"`
	content    string `json:"content"`
	author     string `json:"author"`
}

func (b *Blog) setBlogId(id int32) {
	b.blogId = id
}

func (b *Blog) setTitle(title string) {
	b.title = title
}

func (b *Blog) setTopicGroup(topicGroup string) {
	b.topicGroup = topicGroup
}

func (b *Blog) setCreateTime(createTime string) {
	b.createTime = createTime
}

func (b *Blog) setUpdateTime(updateTime string) {
	b.updateTime = updateTime
}

func (b *Blog) setContent(content string) {
	b.content = content
}

func (b *Blog) setAuthor(author string) {
	b.author = author
}

func (b *Blog) Add(db *sql.DB) (res interface{}, err error) {
	sqlStr := "INSERT INTO blog(title,topicGroup,createTime,updateTime,content,author) VALUES(?,?,?,?,?,?)"
	result, err := db.Exec(sqlStr, b.title, b.topicGroup, b.createTime, b.updateTime, b.content, b.author)
	if err != nil {
		fmt.Printf("insert faile,err:%v\n", err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert failed,err:%v\n", err)
	}
	return id, nil
}

func (b *Blog) Remove(db *sql.DB, id int32) (res interface{}, err error) {
	sqlStr := "DELETE FROM blog WHERE blogId = ?"
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete data failed,err:%v\n", err)
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed,err:%v\n", err)
		return nil, err
	}
	return result, nil
}

func (b *Blog) Update(db *sql.DB) (res interface{}, err error) {
	sqlStr := "UPDATE blog SET title=?,topicGroup=?,createTime=?,updateTime=?,content=?,author=? WHERE blogId=?"
	result, err := db.Exec(sqlStr, b.title, b.topicGroup, b.createTime, b.updateTime, b.content, b.author, b.blogId)
	if err != nil {
		fmt.Printf("update data failed,err:%v\n", err)
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsaffected failed,err:%v\n", err)
		return nil, err
	}
	return n, nil
}

func (b *Blog) Get(db *sql.DB, id int32) (res interface{}, err error) {
	sqlStr := "SELECT blogId,title,topicGroup,createTime,updateTime,content,author FROM blog WHERE blogId=?"
	err = db.QueryRow(sqlStr, id).Scan(b.blogId, b.title, b.topicGroup, b.createTime, b.updateTime, b.content, b.author)
	if err != nil {
		fmt.Printf("get data failed,err:%v\n", err)
		return nil, err
	}
	return b, nil
}

func NewBlog(id int32, title, group, createTime, updateTime, content, author string) *Blog {
	res := Blog{}
	res.setAuthor(author)
	res.setBlogId(id)
	res.setContent(content)
	res.setCreateTime(createTime)
	res.setUpdateTime(updateTime)
	res.setTitle(title)
	res.setTopicGroup(group)
	return &res
}
