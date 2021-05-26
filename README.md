# go_blog_system

### admin
####restApi:
    1. login
    2. add_user
    3. remove_user
    2. add_blog
    3. update_blog
    4. remove_blog
    5. get_blog
####struct:
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