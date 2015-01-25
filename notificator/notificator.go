package notificator

import (
	"github.com/gophergala/not_golang_experts/model"
	"github.com/mailgun/mailgun-go"
	"log"
	"os"
)

func SendPageUpdatedNotification(u *model.User, url string) {
	gopher_env := os.Getenv("GOPHER_ENV")

	if gopher_env == "production" {
		go func() {
			mg := mailgun.NewMailgun("nts.mailgun.org", "key-6muwgm3md06odh43loir2bqoa4dws086", "")

			m := mg.NewMessage(
				"Oscar Swanros <notif@gostalker.com>", // From
				"Update!", // Subject
				"The page "+url+" has been updated. Check it out!", // Plain-text body
				"Oscar Swanros <"+u.Email+">",                      // Recipients (vararg list)
			)

			_, _, err := mg.Send(m)

			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}
