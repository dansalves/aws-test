package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// função que recebe as tags de uma instancia e retorna o valor atribuido a tag Name
func getInstanceName(tags []types.Tag) string {
	for _, tag := range tags {
		if *tag.Key == "Name" {
			return *tag.Value
		}
	}
	return "N/A"
}

// Função para pegar o IP (trata nil)
func getIP(ip *string) string {
	if ip == nil {
		return "Sem IP"
	}
	return *ip
}

func buildInstance(instance types.Instance, reservation *types.Reservation) Ec2Instance {
	i_instance := Ec2Instance{
		AccountID: aws.ToString(reservation.OwnerId), // Usando helper aws.ToString
		Name:      getInstanceName(instance.Tags),
		ID:        aws.ToString(instance.InstanceId),
		Type:      string(instance.InstanceType),
		PublicIP:  getIP(instance.PublicIpAddress),
		PrivateIP: getIP(instance.PrivateIpAddress),
		AZ:        aws.ToString(instance.Placement.AvailabilityZone),
	}
	return i_instance
}

// struct para armazenar os dados nescessarios de EC2
type Ec2Instance struct {
	AccountID string
	Name      string
	ID        string
	Type      string
	PublicIP  string
	PrivateIP string
	AZ        string
}

// struct para armazenar os dados nescessarios de EC2
type EbsVolume struct {
}

func main() {
	// carrega configuraçoes para variavel cfg
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Erro ao carregar configurações: ", err)
	}

	// cria um cliente para utilizar recursos de ec2
	client := ec2.NewFromConfig(cfg)

	// carrega os dados de instancias para a variavel reusltEC2
	resultEC2, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatal("Erro ao listar instâncias: ", err)
	}

	// carrega os dados de volumes para a variavel resultEBS
	/*resultEBS, err := client.DescribeVolumes(context.TODO(), &ec2.DescribeVolumesInput{})
	if err != nil {
		log.Fatal("Erro ao listar instâncias: ", err)
		}*/

	// cria slices de struct para armazenamento dos dados, para instancias e volumes e variaveis auxiliares
	instancesStc := []Ec2Instance{}
	//volumesStc := []EbsVolume{}

	var i_instance Ec2Instance
	//var i_volume EbsVolume

	// itera sobre a raviarel de resultEC2 trazendo as instancias
	for _, reservation := range resultEC2.Reservations {

		// atribuindo cada valor da instancia a variavel para dar o append nas slices
		for _, instance := range reservation.Instances {

			/*i_instance.AccountID = *reservation.OwnerId
			i_instance.Name = getInstanceName(instance.Tags)
			i_instance.ID = *instance.InstanceId
			i_instance.Type = string(instance.InstanceType)
			i_instance.PublicIP = getIP(instance.PublicIpAddress)
			i_instance.PrivateIP = getIP(instance.PrivateIpAddress)
			i_instance.AZ = *instance.Placement.AvailabilityZone*/

			i_instance = buildInstance(instance, &reservation)
			// append de i_instance para a struct instancesStc
			instancesStc = append(instancesStc, i_instance)
		}
	}

	// itera sobre a variavel de resultEBS trazendo os volumes
	//for _, volume := range resultEBS.Volumes {

	//}

	for _, instance := range instancesStc {
		fmt.Println(instance)
	}
}
