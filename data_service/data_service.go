package data_service

import (
	"context"
	"log"

	pb "./service_protos"
)

type DataService struct{}

func (d *DataService) GetByLine(ctx context.Context, req pb.GetDataRequest, resp pb.DataResponse) error {
	log.Print("Received DataService.GetByLine")
}
