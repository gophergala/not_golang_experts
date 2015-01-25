package notificator

import (
	"log"
	"github.com/gophergala/not_golang_experts/model"
	"github.com/mailgun/mailgun-go"
)

func SendPageUpdatedNotification(u *model.User, url string) {
	go func() {
		mg := mailgun.NewMailgun("nts.mailgun.org", "key-6muwgm3md06odh43loir2bqoa4dws086", "")

		m := mg.NewMessage(
				"Oscar Swanros <notif@gostalker.com>",
				"Update!",
				"The page " + url + " has been updated. Check it out!",
				"Oscar Swanros <" + u.Email + ">",
		)

		_, _, err := mg.Send(m)

		if err != nil {
			log.Fatal(err)
		}
	}()
}
