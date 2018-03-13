 ```bash
 curl -H "Private-Token: xxxxxxxxxxxxxxxxxxxx" https://gitlab.hostname.com/api/v4/groups/4/projects | jq '.[].ssh_url_to_repo' -r | xargs -I@ git clone @
 ```