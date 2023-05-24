package adder

import (
	"math/rand"
	"time"

	"github.com/dahuang-purestorage/grpc-streaming-sandbox/pkg/api"
	"github.com/sirupsen/logrus"
)

// GRPCServer struct
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add method for calculate X + Y
func (s *GRPCServer) GetNumber(req *api.GetRequest, stream api.Adder_GetNumberServer) error {
	iter := 0
	for {
		iter += 1

		result := rand.Int31()
		logrus.Infof("sending result %v with iteration %v", result, iter)
		response := api.GetResponse{
			RandomNumber: result,
		}
		if err := stream.Send(&response); err != nil {
			return err
		}

		// if result%61 == 0 {
		// 	logrus.Infof("number is divisble by 13.. done")
		// 	break
		// }

		sleepRand := rand.Intn(10)
		logrus.Infof("Sleeping for %v", sleepRand)
		time.Sleep(time.Duration(sleepRand) * time.Second)

	}
	return nil
}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {

}
