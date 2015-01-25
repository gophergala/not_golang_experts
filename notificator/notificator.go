package notificator

import (
	"github.com/mailgun/mailgun-go"
	"log"
	"strings"
	"os"
)

func SendPageUpdatedNotificationToUsers(emails []string, url string) {
	gopher_env := os.Getenv("GOPHER_ENV")

	if gopher_env == "production" {
		go func() {
			mg := mailgun.NewMailgun("nts.mailgun.org", "key-6muwgm3md06odh43loir2bqoa4dws086", "")

			m := mg.NewMessage(
				"Oscar Swanros <notif@gostalker.com>", // From
				"Update!", // Subject
				"The page "+url+" has been updated. Check it out!", // Plain-text body
				strings.Join(emails, ", "),
			)

			_, _, err := mg.Send(m)

			if err != nil {
				log.Fatal(err)
			}else{
				log.Println("Emails Sent!")
			}
		}()
	}
}
