package seed

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"github.com/ternakkode/packform-backend/internal/data"
	"github.com/ternakkode/packform-backend/pkg"
	"github.com/ternakkode/packform-backend/pkg/bunclient"
	"github.com/ternakkode/packform-backend/pkg/csvreader"
)

func (s *Seed) InitialSeeder(ctx context.Context) error {
	companies, err := csvreader.ReadCSVFile("static/seed/Test task - Postgres - customer_companies.csv")
	if err != nil {
		return err
	}

	companyForDB := make([]data.Company, 0, len(companies)-1)
	for _, line := range companies {
		if len(line) != 2 {
			continue
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}

		companyForDB = append(companyForDB, data.Company{
			ID:          int64(id),
			CompanyName: line[1],
		})
	}

	customers, err := csvreader.ReadCSVFile("static/seed/Test task - Postgres - customers.csv")
	if err != nil {
		return err
	}

	customerForDB := make([]data.Customer, 0, len(customers)-1)
	for _, line := range customers {
		if len(line) != 6 {
			continue
		}

		userID := line[0]
		login := line[1]
		pwd := pkg.HashString(line[2])
		name := line[3]
		companyID, err := strconv.Atoi(line[4])
		if err != nil {
			return err
		}

		var creditCards []string
		if err = json.Unmarshal([]byte(line[5]), &creditCards); err != nil {
			return err
		}

		customerForDB = append(customerForDB, data.Customer{
			UserID:      userID,
			Login:       login,
			Password:    pwd,
			Name:        name,
			CompanyID:   int64(companyID),
			CreditCards: creditCards,
		})
	}

	orders, err := csvreader.ReadCSVFile("static/seed/Test task - Postgres - orders.csv")
	if err != nil {
		return err
	}

	orderForDB := make([]data.Order, 0, len(orders)-1)
	for _, line := range orders {
		if len(line) != 4 {
			continue
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}

		createdAt, err := time.Parse(time.RFC3339, line[1])
		if err != nil {
			return err
		}

		orderForDB = append(orderForDB, data.Order{
			ID:             int64(id),
			CreatedAt:      createdAt,
			OrderName:      line[2],
			CustomerUserID: line[3],
		})
	}

	ordersItems, err := csvreader.ReadCSVFile("static/seed/Test task - Postgres - order_items.csv")
	if err != nil {
		return err
	}

	orderItemForDB := make([]data.OrderItem, 0, len(ordersItems)-1)
	for _, line := range ordersItems {
		if len(line) != 5 {
			continue
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}

		orderID, err := strconv.Atoi(line[1])
		if err != nil {
			return err
		}

		if line[2] == "" {
			line[2] = "0.00"
		}

		pricePerUnit, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return err
		}

		quantity, err := strconv.Atoi(line[3])
		if err != nil {
			return err
		}

		productName := line[4]

		orderItemForDB = append(orderItemForDB, data.OrderItem{
			ID:           int64(id),
			OrderID:      int64(orderID),
			PricePerUnit: pricePerUnit,
			Quantity:     int64(quantity),
			Product:      productName,
		})
	}

	deliveries, err := csvreader.ReadCSVFile("static/seed/Test task - Postgres - deliveries.csv")
	if err != nil {
		return err
	}

	deliveryForDB := make([]data.Delivery, 0, len(deliveries)-1)
	for _, line := range deliveries {
		if len(line) != 3 {
			continue
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}

		orderItemID, err := strconv.Atoi(line[1])
		if err != nil {
			return err
		}

		quantity, err := strconv.Atoi(line[2])
		if err != nil {
			return err
		}

		deliveryForDB = append(deliveryForDB, data.Delivery{
			ID:                int64(id),
			OrderItemID:       int64(orderItemID),
			DeliveredQuantity: quantity,
		})
	}

	tx, err := bunclient.GetConn().BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if _, err := tx.NewInsert().Model(&companyForDB).Exec(ctx); err != nil {
		return err
	}

	if _, err := tx.NewInsert().Model(&customerForDB).Exec(ctx); err != nil {
		return err
	}

	if _, err := tx.NewInsert().Model(&orderForDB).Exec(ctx); err != nil {
		return err
	}

	if _, err := tx.NewInsert().Model(&orderItemForDB).Exec(ctx); err != nil {
		return err
	}

	if _, err := tx.NewInsert().Model(&deliveryForDB).Exec(ctx); err != nil {
		return err
	}

	tx.Commit()

	return nil
}
