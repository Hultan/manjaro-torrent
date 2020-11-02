package notify_user

import (
	"fmt"
	"github.com/hultan/manjaro-torrent/internal/manjaro"
	"github.com/keybase/go-notifier"
	"os"
)

type NotifyUser struct {

}

func New() *NotifyUser {
	return new(NotifyUser)
}

func (n *NotifyUser) NotifyUserIfNeeded(manjaroNew, manjaroOld *manjaro.Manjaro) {
	for name, oldVersion := range manjaroOld.Distributions {
		newVersion, ok := manjaroNew.Distributions[name]
		if ok == false {
			// We found a new edition
			notify, err := notifier.NewNotifier()
			if err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("Failed to create notifier! %s", err.Error()))
				break
			}
			notify.DeliverNotification(notifier.Notification{
				Title:   fmt.Sprintf("Manjaro %s has been added!", newVersion.Name),
				Message: fmt.Sprintf("Add torrent for %s.\n\nNew version : %s", newVersion.Name, newVersion.Version),
			})
		}
		if newVersion.Version != oldVersion.Version {
			// We found an updated edition
			notify, err := notifier.NewNotifier()
			if err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("Failed to create notifier! %s", err.Error()))
				break
			}
			notify.DeliverNotification(notifier.Notification{
				Title:   fmt.Sprintf("Manjaro %s has been updated!", newVersion.Name),
				Message: fmt.Sprintf("Update torrent for %s.\n\nOld version : %s\nNew version : %s", newVersion.Name, newVersion.Version, oldVersion.Version),
			})
		}
	}
}
