package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func checkDeadlines() {
	tasks, err := getTasks()

	if err != nil {
		fmt.Println("Something went wrong getting tasks:", err)
		return
	}

	now := time.Now()

	for _, task := range tasks {

		timeUntilDeadline := task.Deadline.Sub(now)

		if timeUntilDeadline <= 31*time.Minute && !task.ReminderSent {

			err := beeep.Notify(
				"Task Reminder",
				task.Name+" is due in 30 minutes!",
				"",
			)

			if err != nil {
				fmt.Println("Could not send notification:", err)
				continue
			}

			db.Model(&task).Update("reminder_sent", true)
		}

		if timeUntilDeadline <= 1*time.Minute && !task.DeadlineSent {

			err := beeep.Notify(
				"Task Reminder",
				task.Name+" is due !!",
				"",
			)

			if err != nil {
				fmt.Println("Could not send notification:", err)
				continue
			}

			db.Model(&task).Update("deadline_sent", true)
		}

		if timeUntilDeadline <= -20*time.Minute && !task.OverdueSent {

			err := beeep.Notify(
				"Task Reminder",
				task.Name+" is overdue !!",
				"",
			)

			if err != nil {
				fmt.Println("Could not send notification:", err)
				continue
			}

			db.Model(&task).Update("overdue_sent", true)
		}
	}
}
