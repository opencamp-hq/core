package notify

import (
	"fmt"
	"os"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/opencamp-hq/core/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailNotifier struct {
	log       log15.Logger
	apiKey    string
	fromName  string
	fromEmail string
}

func (n *EmailNotifier) Notify(toName, toEmail string, cg models.Campground, startDate, endDate string, sites models.Campsites) error {
	from := mail.NewEmail(n.fromName, n.fromEmail)
	subject := "Good news! Your campground is available"
	to := mail.NewEmail(toName, toEmail)

	content := fmt.Sprintf(`
	'%s' has sites available from %s to %s!

// HERE
// in the case of multiple sites being available at the same time, want to support that...
// so yeah need to do some sort of strings.Join or iterator
// ... does go templating have anything built in?
//  .. can we attach a template language to this?


	Sites:
	%s
	
	To reserve: <link goes here>`, campgroundName, startDate, endDate, " - Site "+strings.Join(sites, "\n - Site "))

	plainTextContent := content
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(message)
	if err != nil {
		return err
	}

	n.log.Debug("Email sent", "status", resp.StatusCode, "to", to)
	return nil
}
