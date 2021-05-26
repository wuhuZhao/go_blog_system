package blog_strcut

type blog interface {
	setBlogId(id int32)
	setTitle(title string)
	setTopicGroup(topicGroup string)
	setCreateTime(createTime string)
	setUpdateTime(updateTime string)
	setContent(content string)
	setAuthor(author string)
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

func (b *Blog) Add() {
	panic("implement me")
}

func (b *Blog) Remove() {
	panic("implement me")
}

func (b *Blog) Delete() {
	panic("implement me")
}

func (b *Blog) Get() {
	panic("implement me")
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
