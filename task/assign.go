package task

import (
	"PharmaProject/config"
	"PharmaProject/conn"
	"PharmaProject/models"
	"PharmaProject/repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1/tasks"
)

func AddMedicine(pld string) error {
	fmt.Println("error")
	meds := []models.Medicine{}
	if err := json.Unmarshal([]byte(pld), &meds); err != nil {
		log.Println("Could not decode create notification task payload")
		return err
	}

	fmt.Println(meds)

	// db := conn.ConnectDB()
	// repo := repository.NewSQLNotificationRepository(db)
	result, err := repository.NewMedicineRepo().AddBulkMedicine(meds)
	if err != nil {
		// helper.CaptureError(err, map[string]string{"source": "task", "message": "failed to create notification"})
		// log.Println(err.Error())
		// helper.Status("create_assign_notification", c.AppName, http.StatusInternalServerError)
		return err
	}
	fmt.Println(result)
	return nil

}

func pushTask(name string, payload []byte) error {
	bk := config.Worker().BindingKey + "_assign"
	_, err := conn.DefaultAssignWorker().SendTask(&tasks.Signature{
		Name:         name,
		RoutingKey:   bk,
		RetryCount:   1,
		RetryTimeout: 10,
		Args: []tasks.Arg{
			{
				Type:  "json",
				Value: payload,
			},
		},
	})
	return err
}