package db

import (
    "context"
    "votacao-paredao-bbb/core/models"
    "votacao-paredao-bbb/core/ports"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/bson"
)

type VotoMongoRepository struct {
    Client *mongo.Client
}

func (r *VotoMongoRepository) SalvarVoto(voto models.Voto) error {
    collection := r.Client.Database("votacao_bbb").Collection("votos")
    _, err := collection.InsertOne(context.Background(), voto)
    return err
}

func (r *VotoMongoRepository) ObterVotos() ([]models.Voto, error) {
    var votos []models.Voto
    collection := r.Client.Database("votacao_bbb").Collection("votos")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }

    err = cursor.All(context.Background(), &votos)
    if err != nil {
        return nil, err
    }
    return votos, nil
}

func NovoVotoMongoRepository(client *mongo.Client) *VotoMongoRepository {
    return &VotoMongoRepository{Client: client}
}

func ConectarMongoDB(uri string) (*mongo.Client, error) {
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }
    err = client.Connect(context.Background())
    return client, err
}