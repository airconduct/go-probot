package gitlab_test

import (
	"testing"
	// "github.com/xanzy/go-gitlab"
)

func TestGitlab(t *testing.T) {
	// gitlab.NewClient()
	// gitlab.ParseWebhook()
	// gitlab.WebhookEventType()
}

// sudo docker run --detach \
//   --hostname gitlab.example.com \
//   --publish 443:443 --publish 80:80 --publish 422:22 \
//   --name gitlab \
//   --shm-size 256m \
//   gitlab/gitlab-ee:latest
//   --volume $GITLAB_HOME/config:/etc/gitlab \
//   --volume $GITLAB_HOME/logs:/var/log/gitlab \
//   --volume $GITLAB_HOME/data:/var/opt/gitlab \
