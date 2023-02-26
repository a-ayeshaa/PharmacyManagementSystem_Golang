package task

import (
	"PharmaProject/internal/config"
	"PharmaProject/domain"
	"PharmaProject/internal/conn"
	"PharmaProject/medicine/repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1/tasks"
)

func AddMedicine(pld string) error {
	fmt.Println("error")
	meds := []domain.Medicine{}
	if err := json.Unmarshal([]byte(pld), &meds); err != nil {
		log.Println("Could not decode create notification task payload")
		return err
	}

	fmt.Println(meds)
	db := conn.ConnectDB()
	result, err := repository.New(db).AddBulkMedicine(meds)
	if err != nil {
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
