# hackthecloud-workshop-uhv5
Demonstrate a REST API to show how to start hosting software using a fully self-hosted, open-source stack.

# Stack
- Coolify
- Golang
- PostgreSQL
- Docker (as part of Coolify)
- GitOps

# Routes
### Create Task
`POST /tasks`
```code
curl -X POST -H "Content-Type:application/json" -d '{"title":"Test","description":"Test task"}' https://app_url/tasks
```

### View added tasks
`GET /tasks`

```code
curl https://your-app/tasks
```


### Remove tasks
- To be added during workshop
`DELETE /tasks/:id`
```code
curl -X DELETE -H "Content-Type:application/json" https://app_url/tasks/<task_id>
```
