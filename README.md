# go_blog_system

### admin
###restApi:
  1. login
  2. add_user
  3. remove_user
  4. add_blog
  5. update_blog
  6. remove_blog
  7. get_blog
###struct:
```js
user = {
    username: '',
    password: '',
    auth: 'root_user' //now only root_user
}
blog = {
    blog_id: '',
    title: '',
    topic_group: '',
    create_time: '',
    update_time: '',
    content: '',
    author: ''
}
topic = {
    topic_id:'',
    topic_group: '',
}
```