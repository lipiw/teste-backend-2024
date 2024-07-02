package consumers

import (
    "ms-go/app/services/products"
    "ms-go/app/models"

    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/segmentio/kafka-go"
)

func RunConsumer() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:     []string{"kafka:29092", "localhost:9092", "localhost:9093", "localhost:9094"},
        Topic:       "rails-to-go",
        Partition:   0,
        MinBytes:    10e3,
        MaxBytes:    10e6,
        GroupID:     "ms-rails-consumer-group",
        StartOffset: kafka.FirstOffset,
    })

    for {
        m, err := r.FetchMessage(context.Background())
        if err != nil {
            log.Fatalf("Erro ao ler mensagem: %v", err)
        }

        fmt.Printf("Mensagem em offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

        var product models.Product
        if err := json.Unmarshal(m.Value, &product); err != nil {
            log.Printf("Erro ao converter dados JSON: %v", err)
            continue
        }

        createdProduct, err := products.Create(product, false)
        if err != nil {
            log.Printf("Erro ao inserir no banco de dados: %v", err)
            continue
        }

        log.Printf("Produto criado: %v", createdProduct)

        if err := r.CommitMessages(context.Background(), m); err != nil {
            log.Printf("Erro ao fazer commit do offset: %v", err)
        }
    }

    if err := r.Close(); err != nil {
        log.Fatalf("Erro ao fechar o leitor: %v", err)
    }
}